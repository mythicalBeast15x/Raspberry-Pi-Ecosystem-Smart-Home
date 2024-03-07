package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/jacobsa/go-serial/serial"
	"io"
	"time"
)

// Message struct represents the JSON message format
type Message struct {
	Content string `json:"content"`
}

func generateRandomKey(keySize int) ([]byte, error) {
	key := make([]byte, keySize)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func Encrypt(plaintext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// CBC mode requires an initialization vector (IV)
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	// Pad plaintext if needed
	plaintext = PKCS7Padding(plaintext, aes.BlockSize)

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	copy(ciphertext[:aes.BlockSize], iv)

	return ciphertext, nil
}

// PKCS7Padding pads data using the PKCS#7 scheme
func PKCS7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

func controller(key []byte) {
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

		// Encrypt the JSON data
		encryptedData, err := Encrypt(jsonData, key)
		if err != nil {
			fmt.Printf("Error encrypting data: %v\n", err)
			continue
		}

		// Write the encrypted data to the serial port
		_, err = port.Write(encryptedData)
		if err != nil {
			fmt.Printf("Error writing to serial port: %v\n", err)
			continue
		}

		fmt.Printf("Message sent: %s\n", message.Content)
		time.Sleep(1 * time.Second) // Send a message every second
	}
}

func main() {
	// Generate a random encryption key
	key, err := generateRandomKey(16)
	if err != nil {
		fmt.Println("Error generating random key:", err)
		return
	}

	controller(key)
}
