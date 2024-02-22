package main

import (
	"fmt"
	"go.bug.st/serial"
	"log"
)

// Function to get the list of available serial ports
func getPortsList() ([]string, error) {
	return serial.GetPortsList()
}

// Function to open a serial port
func openPort(portName string) (serial.Port, error) {
	mode := &serial.Mode{
		BaudRate: 115200, // Adjust baud rate as needed
	}
	return serial.Open(portName, mode)
}

// Function to continuously read from a serial port
func readFromPort(port serial.Port) {
	buff := make([]byte, 100)
	for {
		n, err := port.Read(buff)
		if err != nil {
			log.Fatal(err)
			break
		}
		if n == 0 {
			fmt.Println("\nEOF")
			break
		}
		fmt.Printf("%v", string(buff[:n]))
	}
}

// Function to continuously write to a serial port
func writeToPort(port serial.Port) {
	for {
		_, err := port.Write([]byte("Hello, Zigbee!\n")) // Adjust message as needed
		if err != nil {
			log.Fatal(err)
			break
		}
	}
}

/*
func main() {
	// Get the list of available serial ports
	ports, err := getPortsList()
	if err != nil {
		log.Fatal(err)
	}

	// Print the list of available serial ports
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}
	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port)
	}

	// Open the first serial port (you may need to adjust this based on your setup)
	port, err := openPort(ports[0])
	if err != nil {
		log.Fatal(err)
	}
	defer port.Close()

	// Concurrently start reading and writing to the serial port
	go readFromPort(port)
	go writeToPort(port)

	// Keep the main goroutine alive
	select {}
}
*/
