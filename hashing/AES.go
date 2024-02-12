package hashing

/*
import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
)

//go get github.com/google/gopacket
//go get github.com/google/gopacket/pcap


func main() {
	networkadapters()
	// Define the network interface you want to capture packets from
	device := "\\Device\\NPF_{285EEC32-9286-4AB7-83F9-0C2E8BDE60FC}"

	// Open the network device for packet capture
	handle, err := pcap.OpenLive(device, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Create a packet source to decode packets from the network interface
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	// Capture packets for a specified duration (e.g., 30 seconds)
	duration := 30 * time.Second
	endTime := time.Now().Add(duration)

	fmt.Printf("Capturing network traffic on interface %s for %s...\n", device, duration)

	// Loop to capture and analyze packets
	for packet := range packetSource.Packets() {
		// Print basic packet information
		fmt.Printf("Packet captured at %s\n", packet.Metadata().Timestamp)
		fmt.Printf("Packet length: %d bytes\n", packet.Metadata().Length)
		fmt.Printf("Packet data:\n%s\n", packet.Dump())

		// You can perform more advanced analysis on the packet data here

		// Check if the capture duration has elapsed
		if time.Now().After(endTime) {
			fmt.Println("Capture duration reached. Exiting...")
			break
		}
	}
}
*/
//AES:

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

// Encrypt encrypts plaintext using AES with the given key and a random IV.
func Encrypt(plaintext, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	plaintextBytes := []byte(plaintext)
	ciphertext := make([]byte, aes.BlockSize+len(plaintextBytes))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintextBytes)
	//fmt.Println(base64.URLEncoding.EncodeToString(ciphertext))
	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// Decrypt decrypts ciphertext using AES with the given key.
func Decrypt(ciphertext, key string) (string, error) {
	ciphertextBytes, err := base64.URLEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	if len(ciphertextBytes) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}
	iv := ciphertextBytes[:aes.BlockSize]
	ciphertextBytes = ciphertextBytes[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertextBytes, ciphertextBytes)
	return string(ciphertextBytes), nil
}

/*
func main() {
	key := "your-secret-aes-key-here" // Ensure this is 16/24/32 bytes for AES-128/192/256
	plaintext := "exampleplaintext"
	encrypted, err := Encrypt(plaintext, key)
	if err != nil {
		log.Fatalf("Encryption failed: %v", err)
	}
	fmt.Printf("Encrypted: %s\n", encrypted)
	decrypted, err := Decrypt(encrypted, key)
	if err != nil {
		log.Fatalf("Decryption failed: %v", err)
	}
	fmt.Printf("Decrypted: %s\n", decrypted)
}
*/
/*
AES Test case:
package main
import (
"testing"
)
func TestEncryptDecrypt(t *testing.T) {
key := "thisisasecretkey1234567890" // 32 bytes for AES-256
originalText := "Hello, World!"
encryptedText, err := Encrypt(originalText, key)
if err != nil {
t.Errorf("Failed to encrypt: %v", err)
}
decryptedText, err := Decrypt(encryptedText, key)
if err != nil {
t.Errorf("Failed to decrypt: %v", err)
}
if decryptedText != originalText {
t.Errorf("Decrypted text does not match original. Got %s, want %s", decryptedText,
originalText)
}
}
*/
