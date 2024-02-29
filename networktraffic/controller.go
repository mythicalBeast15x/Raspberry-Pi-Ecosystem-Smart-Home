package main

import (
	"encoding/json"
	"fmt"
	"github.com/jacobsa/go-serial/serial"
	"time"
)

// Message struct represents the JSON message format
type Message struct {
	Content string `json:"content"`
	// You can add more fields as needed
}

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

	// Create a message struct
	message := Message{
		Content: "Hello from Server Zigbee",
	}

	for {
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

		fmt.Printf("Message sent: %s\n", message.Content)
		time.Sleep(1 * time.Second) // Send a message every second
	}
}

func main() {
	controller()
}
