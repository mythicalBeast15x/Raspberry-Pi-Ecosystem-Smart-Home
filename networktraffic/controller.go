package networktraffic

import (
	"CMPSC488SP24SecThursday/hashing"
	"CMPSC488SP24SecThursday/messaging"
	"fmt"
	"github.com/jacobsa/go-serial/serial"
	"io"
	"time"
)

// Controller function to handle message creation and sending
func Controller(msg string, oMessages *messaging.OpenMessages, qMessages *messaging.MessageQueue) {
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

	defer func(port io.ReadWriteCloser) {
		err := port.Close()
		if err != nil {

		}
	}(port)

	// Encrypt message into a single byte string of characters
	key := []byte("1234567890123456") // TODO - replace with generated AES key
	encryptedMsg, err := hashing.Encrypt([]byte(msg), key)
	if err != nil {
		fmt.Println("Error encrypting JSON:", err)
		return
	}

	// Encase the encrypted message with open and closed brackets to mark the start and end of the message
	completeMsg := fmt.Sprintf("{%s}", encryptedMsg)

	for i := 0; i < 1; i++ {
		for {
			time.Sleep(4500 * time.Millisecond)

			// Write the JSON data to the serial port
			_, err = port.Write([]byte(completeMsg))
			if err != nil {
				fmt.Printf("Error writing to serial port: %v\n", err)
				continue
			}
			//fmt.Printf("\nMessage sent: %s\n", msgBytes)
			fmt.Printf("\nMessage sent: %s\n", completeMsg)
			time.Sleep(2500 * time.Millisecond)
			break
		}
	}
}

/*
func main() {
	Controller()
}
*/
