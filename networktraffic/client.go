package networktraffic

import (
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

		// Unmarshal JSON data into response struct
		var response Response
		err = json.Unmarshal(jsonData, &response)
		if err != nil {
			fmt.Printf("Error unmarshalling JSON: %v\n", err)
			continue
		}

		fmt.Printf("Message received: %s\n", response.Content)
	}
}

/*
func main() {
	client()
}

*/
