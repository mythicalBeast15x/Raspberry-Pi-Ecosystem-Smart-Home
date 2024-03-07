package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"fmt"
	"github.com/jacobsa/go-serial/serial"
	"io"
)

// Response struct represents the JSON response format from controller
type Response struct {
	Content string `json:"content"`
}

func Decrypt(ciphertext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// CBC mode requires an initialization vector (IV)
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// Decrypt the ciphertext
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	// Unpad plaintext
	plaintext := PKCS7Unpadding(ciphertext)

	return plaintext, nil
}

// PKCS7Unpadding removes padding from data that was padded using the PKCS#7 scheme
func PKCS7Unpadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}

func client(key []byte) {
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

		// Decrypt the received data
		decryptedData, err := Decrypt(jsonData, key)
		if err != nil {
			fmt.Printf("Error decrypting data: %v\n", err)
			continue
		}

		// Unmarshal JSON data into response struct
		var response Response
		err = json.Unmarshal(decryptedData, &response)
		if err != nil {
			fmt.Printf("Error unmarshalling JSON: %v\n", err)
			continue
		}

		fmt.Printf("Message received: %s\n", response.Content)
	}
}

func main() {
	// Use the same encryption key as the controller (for testing only)
	// Network team will modify this with their AES function
	key := []byte{49, 212, 173, 219, 62, 163, 49, 72, 148, 209, 213, 50, 153, 218, 240, 68}
	client(key)
}
