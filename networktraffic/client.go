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

	port, err := serial.Open(options)
	if err != nil {
		fmt.Printf("Error opening serial port: %v\n", err)
		return
	}
	defer port.Close()

	reader := bufio.NewReader(port)
	for {
		// Read encrypted data length from the serial port
		lengthStr, err := reader.ReadString(':')
		if err != nil {
			if err != io.EOF {
				fmt.Printf("Error reading from serial port: %v\n", err)
			}
			continue
		}

		// Parse the length
		length, err := strconv.Atoi(strings.TrimSuffix(lengthStr, ":"))
		if err != nil {
			fmt.Printf("Error parsing message length: %v\n", err)
			continue
		}

		// Read encrypted data of specified length
		encryptedData := make([]byte, length)
		_, err = io.ReadFull(reader, encryptedData)
		if err != nil {
			if err != io.EOF {
				fmt.Printf("Error reading from serial port: %v\n", err)
			}
			continue
		}

		// Find the index of the first '=' character
		charIndex := bytes.IndexByte(encryptedData, '=')
		if charIndex != -1 {
			// Remove characters after '=' character
			encryptedData = encryptedData[:charIndex+1]
		}

		fmt.Print("\nReceived Encrypted Message:\n", string(encryptedData))

		// Decrypt data
		key := []byte("1234567890123456") // TODO - replace with generated AES key
		decryptedMsg, err := hashing.Decrypt(string(encryptedData), key)
		if err != nil {
			fmt.Print("Error decrypting JSON:", err)
			return
		}
		fmt.Println("Decrypted Message:\n", decryptedMsg)

		// Add the received message to the incoming messages queue in MessageQueue instance
		qMessages.IncomingMessages = append(qMessages.IncomingMessages, string(decryptedMsg))
	}
}

/*
func main() {
	Client()
}
*/
