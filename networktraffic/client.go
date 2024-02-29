package main

import (
	"bufio"
	"encoding/json"
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

		// Unmarshal JSON data into message struct
		var message Message
		err = json.Unmarshal(jsonData, &message)
		if err != nil {
			fmt.Printf("Error unmarshalling JSON: %v\n", err)
			continue
		}

		fmt.Printf("Message received: %s\n", message.Content)
	}
}

func main() {
	client()
}
