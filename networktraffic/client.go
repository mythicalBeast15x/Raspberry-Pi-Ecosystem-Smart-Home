package networktraffic

import (
	"CMPSC488SP24SecThursday/hashing"
	"CMPSC488SP24SecThursday/messaging"
	"bufio"
	"fmt"
	"github.com/jacobsa/go-serial/serial"
	"io"
	"strings"
	"time"
)

// Client function to handle receiving and processing messages
func Client(oMessages *messaging.OpenMessages, qMessages *messaging.MessageQueue) {
	options := serial.OpenOptions{
		PortName:        "/dev/ttyUSB0",
		BaudRate:        9600,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}

clientLoop:
	for {
		// Open serial port
		port, err := serial.Open(options)
		if err != nil {
			fmt.Printf("\nError opening serial port: %v\n", err)
			time.Sleep(1000 * time.Millisecond)
			continue // Restart the loop to try opening the serial port again
		}

		reader := bufio.NewReader(port)
		for {
			// Clear the reader buffer
			reader.Reset(port)

			// Read until the start of a message ('{')
			for {
				startByte, err := reader.ReadByte()
				if err != nil {
					if err != io.EOF {
						fmt.Printf("\nError reading from serial port: %v\n", err)
					}
					err := port.Close()
					if err != nil {
						return
					}
					continue clientLoop // Restart the client loop if serial port error occurs
				}
				if startByte == '{' {
					break // Break out of the loop if the open bracket is found: Beginning of new message
				}
			}

			// Read bytes until the end of the message ('}')
			var encryptedData strings.Builder
			for {
				nextByte, err := reader.ReadByte()
				if err != nil {
					if err != io.EOF {
						fmt.Printf("\nError reading from serial port: %v\n", err)
					}
					err := port.Close()
					if err != nil {
						return
					}
					continue clientLoop // Restart the client loop  if serial port error occurs
				}
				if nextByte == '}' {
					break // Break out of the loop if the closed bracket is found: End of new message
				}
				encryptedData.WriteByte(nextByte)
			}

			// Process the received message
			fmt.Print("\nReceived Encrypted Message:\n", encryptedData.String())

			// Remove '{' and '}' character markers
			message := strings.Trim(encryptedData.String(), "{}")

			// Decrypt data
			key := []byte("1234567890123456") // TODO: replace with generated AES key
			decryptedMsg, err := hashing.Decrypt(message, key)
			if err != nil {
				fmt.Printf("\nError decrypting JSON: %v\n", err)
				err := port.Close()
				if err != nil {
					return
				}
				continue clientLoop // Restart the client loop if decryption error occurs
			}

			fmt.Println("\nDecrypted Message:\n", string(decryptedMsg))

			// Add the received message to the incoming messages queue in MessageQueue instance
			if string(decryptedMsg) != "" {
				qMessages.IncomingMessages = append(qMessages.IncomingMessages, string(decryptedMsg))
			} else {
				fmt.Println("\nDecrypted message is empty.")
			}
			time.Sleep(1000 * time.Millisecond)
		}
	}
}

/*
func main() {
	Client()
}
*/
