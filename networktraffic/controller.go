package networktraffic

import (
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

	message := "Hello from Server Zigbee\n"
	for {
		_, err := port.Write([]byte(message))
		if err != nil {
			fmt.Printf("Error writing to serial port: %v\n", err)
			continue
		}
		fmt.Printf("Message sent: %s", message)
		time.Sleep(1 * time.Second) // Send a message every second
	}
}
