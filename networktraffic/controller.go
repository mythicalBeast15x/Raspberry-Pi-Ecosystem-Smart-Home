/*
Controller function may be changed to take in a value like a JSON string
to act as the message to be sent.

Example:
func Controller(msg string){
	message := Message{
	Content: msg,
	}
}
*/

package networktraffic

import (

	"CMPSC488SP24SecThursday/messaging" // Importing messaging package

	"fmt"
	"github.com/jacobsa/go-serial/serial"
	"time"
)


func controller() {

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


	// Create a message queue
	qMessages := &messaging.MessageQueue{}

	for {
		// Dequeue a message from the outgoing queue
		outgoingMsg, err := qMessages.Dequeue("outgoing")

		if err != nil {
			fmt.Println("Error dequeuing from outgoing queue:", err)
			continue
		}

		// Convert the message to []byte
		jsonData, ok := outgoingMsg.([]byte)
		if !ok {
			fmt.Println("Expected dequeued message to be of type []byte")
			continue
		}

		// Write the JSON data to the serial port
		_, err = port.Write(jsonData)
		if err != nil {
			fmt.Printf("Error writing to serial port: %v\n", err)
			continue
		}


		fmt.Println("Message sent:", string(jsonData))

		time.Sleep(1 * time.Second) // Send a message every second
	}
}

/* Testing
// Example usage:
func main() {
	// Initialize necessary structures
	oMessages := &messaging.OpenMessages{}
	qMessages := &messaging.MessageQueue{}

	// Dummy message in JSON format
	dummyMessage := `{"messageID": "12345", "senderID": "Pi-1", "receiverID": "Pi-2", "domain": "Testing", "operationID": "TestOperation", "Data": {"key1": "value1", "key2": 42, "key3": true}}`

	// Call the Controller function with the dummy message
	Controller(dummyMessage, oMessages, qMessages)
}



func main() {
	Controller()
}
*/
