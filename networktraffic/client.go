package networktraffic

import (
	//"CMPSC488SP24SecThursday/messaging"
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/jacobsa/go-serial/serial"
	"io"
)

// Response struct represents the JSON response format from controller
type Response struct {
	Content string `json:"content"`
}

/*
Client function may take in the global instance of a Message Queue like the following:

Client(qMessages messaging.MessageQueue) {
}
*/

func Client() {
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

		// Unmarshal JSON data into response struct
		var response Response
		err = json.Unmarshal(jsonData, &response)
		if err != nil {
			fmt.Printf("Error unmarshalling JSON: %v\n", err)
			continue
		}

		fmt.Printf("Message received: %s\n", response.Content)
		/* Once a message is received, the content of the response may be placed into the
		MessageQueue instance like shown below.*/
		//qMessages.IncomingMessages = append(qMessages.IncomingMessages, string(response.Content))
	}
}

/*
func main() {
	Client()
}

*/
