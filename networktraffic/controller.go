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
	"CMPSC488SP24SecThursday/messaging" // Import the messaging package
	"encoding/json"
	"fmt"
	"github.com/jacobsa/go-serial/serial"
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
	defer port.Close()

	for {
		// Unmarshal the JSON string to create a message
		var message messaging.Message
		err := json.Unmarshal([]byte(msg), &message)
		if err != nil {
			fmt.Printf("Error unmarshalling JSON: %v\n", err)
			continue
		}

		// Marshal the message struct to JSON
		jsonData, err := json.Marshal(message)
		if err != nil {
			fmt.Printf("Error marshalling JSON: %v\n", err)
			continue
		}

		// Write the JSON data to the serial port
		_, err = port.Write(jsonData)
		if err != nil {
			fmt.Printf("Error writing to serial port: %v\n", err)
			continue
		}

		fmt.Printf("Message sent: %s\n", message.MessageID)
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
