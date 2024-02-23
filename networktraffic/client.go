package networktraffic

import (
	"bufio"
	"fmt"
	"github.com/jacobsa/go-serial/serial"
	"io"
)

func client() {
	options := serial.OpenOptions{
		PortName:        "/dev/ttyUSB1",
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
		message, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Printf("Error reading from serial port: %v\n", err)
			}
			continue
		}
		fmt.Printf("Message received: %s", message)
	}
}
