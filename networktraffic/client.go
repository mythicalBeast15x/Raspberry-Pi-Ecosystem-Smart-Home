package networktraffic

import (
	"CMPSC488SP24SecThursday/messaging" // Importing messaging package
	"bufio"
	"fmt"
	"github.com/jacobsa/go-serial/serial"
	"io"
)

func client() {
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
		// Read JSON data from the serial port
		jsonData, err := reader.ReadBytes('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Printf("Error reading from serial port: %v\n", err)
			}
			continue
		}

		// Enqueue the incoming message
		qMessages := &messaging.MessageQueue{}
		qMessages.IncomingMessages = append(qMessages.IncomingMessages, string(jsonData))

		// Validate and decrypt the message
		err = messaging.ValidateAndDecrypt(nil, qMessages, nil)
		if err != nil {
			fmt.Printf("Error validating and decrypting message: %v\n", err)
			continue
		}

		// Dequeue a message from the deserialized queue
		deserialMsg, err := qMessages.Dequeue("deserial")
		if err != nil {
			fmt.Println("Error dequeuing from deserialized queue:", err)
			continue
		}

		// Convert the message to a Message struct
		message, ok := deserialMsg.(messaging.Message)
		if !ok {
			fmt.Println("Expected dequeued message to be of type messaging.Message")
			continue
		}

		// Display the received message
		fmt.Println("Message received:")
		messaging.DisplayMessage(message)
	}
}

/*
func main() {
	client()
}
*/
