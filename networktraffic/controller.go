package networktraffic

import (
	"CMPSC488SP24SecThursday/hashing"
	"CMPSC488SP24SecThursday/messaging"
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

	// Encrypt message into a single byte string of characters
	key := []byte("1234567890123456") // TODO - replace with generated AES key
	encryptedMsg, err := hashing.Encrypt([]byte(msg), key)
	if err != nil {
		fmt.Println("Error encrypting JSON:", err)
		return
	}

	// Prefix the encrypted message with its length
	completeMsg := fmt.Sprintf("%d:%s", len(encryptedMsg), encryptedMsg)

	for i := 0; i < 2; i++ {
		for {
			//time.Sleep(1 * time.Second)
			// Write the JSON data to the serial port
			_, err = port.Write([]byte(completeMsg))
			if err != nil {
				fmt.Printf("Error writing to serial port: %v\n", err)
				continue
			}
			//fmt.Printf("\nMessage sent: %s\n", msgBytes)
			fmt.Printf("\nMessage sent: %s\n", completeMsg)
			time.Sleep(2000 * time.Millisecond)
			break
		}
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
