package networktraffic

import (
	"CMPSC488SP24SecThursday/hashing"
	"CMPSC488SP24SecThursday/messaging"
	"bufio"
	"bytes"
	"fmt"
	"github.com/jacobsa/go-serial/serial"
	"io"
	"strconv"
	"strings"
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
	for {
		// Open serial port
		port, err := serial.Open(options)
		if err != nil {
			fmt.Printf("Error opening serial port: %v\n", err)
			continue // Restart the loop to try opening the serial port again
		}

		reader := bufio.NewReader(port)
		for {
			// Clear the reader buffer
			reader.Reset(port)

			// Read encrypted data length from the serial port
			lengthStr, err := reader.ReadString(':')
			if err != nil {
				if err != io.EOF {
					fmt.Printf("Error reading from serial port: %v\n", err)
				}
				break // Continue to the next iteration of the outer loop to reopen the serial port
			}

			// Parse the length
			length, err := strconv.Atoi(strings.TrimSuffix(lengthStr, ":"))
			if err != nil {
				fmt.Printf("Error parsing message length: %v\n", err)
				continue // Continue listening for the next message
			}

			// Read encrypted data of specified length
			encryptedData := make([]byte, length)
			_, err = io.ReadFull(reader, encryptedData)
			if err != nil {
				if err != io.EOF {
					fmt.Printf("Error reading from serial port: %v\n", err)
				}
				break // Continue to the next iteration of the outer loop to reopen the serial port
			}

			// Find the index of the first ' ' character
			charIndex := bytes.IndexByte(encryptedData, ' ')
			if charIndex != -1 {
				// Remove characters after ' ' character
				encryptedData = encryptedData[:charIndex]
			}

			fmt.Print("\nReceived Encrypted Message:\n", string(encryptedData))

			// Decrypt data
			key := []byte("1234567890123456") // TODO - replace with generated AES key
			decryptedMsg, err := hashing.Decrypt(string(encryptedData), key)
			if err != nil {
				fmt.Printf("Error decrypting JSON: %v\n", err)

				// Check if the error is due to base64 decoding issues
				if strings.Contains(err.Error(), "illegal base64 data") {
					// Handle base64 decoding error
					fmt.Println("Base64 decoding error.")
					continue // Continue listening for the next message
				} else {
					// Handle other decryption errors
					fmt.Println("Other decryption error.")
					decryptedMsg = []byte("") // Set decryptedMsg to empty string to prevent corrupted data from being added to the queue
					continue                  // Continue listening for the next message
				}
			}

			fmt.Println("Decrypted Message:\n", string(decryptedMsg))

			// Add the received message to the incoming messages queue in MessageQueue instance
			if string(decryptedMsg) != "" {
				qMessages.IncomingMessages = append(qMessages.IncomingMessages, string(decryptedMsg))
			} else {
				fmt.Println("Decrypted message is empty.")
			}

		}
		// Close serial port
		if err := port.Close(); err != nil {
			fmt.Printf("Error closing serial port: %v\n", err)
		}
	}
}

/*
func main() {
	Client()
}
*/
