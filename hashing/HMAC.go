package hashing

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/joho/godotenv"
)

// CalculateHMAC calculates HMAC for a given message and key
func CalculateHMAC(message, key []byte) string {
	hash := hmac.New(sha256.New, key)
	hash.Write(message)
	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}

// VerifyHMAC verifies the HMAC of a received message
func VerifyHMAC(receivedMessage, receivedHMAC, key []byte) bool {
	calculatedHMAC := CalculateHMAC(receivedMessage, key)
	return hmac.Equal([]byte(calculatedHMAC), receivedHMAC)
}

func loadingEnvVariables() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file: %v", err)
	}
	return nil
}

/*
func main() {
	// Load environment variables from .env file
	err := loadingEnvVariables()
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve the HMAC key from the environment variable
	hmacKey := os.Getenv("HMAC_KEY")
	if hmacKey == "" {
		fmt.Println("HMAC_KEY not found in environment variables")
		return
	}

	// Simulate sending a message from one XBee device to another
	senderMessage := []byte("Hello, receiver!")
	fmt.Println("Sender message:", string(senderMessage))

	// Calculate HMAC for the message
	hmacValue := CalculateHMAC(senderMessage, []byte(hmacKey))
	fmt.Println("Sender HMAC:", hmacValue)

	// Simulate receiving the message and HMAC on the receiver end
	receivedMessage := senderMessage
	receivedHMAC := []byte(hmacValue)

	// Verify the received HMAC
	isValid := VerifyHMAC(receivedMessage, receivedHMAC, []byte(hmacKey))
	if isValid {
		fmt.Println("HMAC verified. Message integrity confirmed.")
	} else {
		fmt.Println("HMAC verification failed. Message integrity compromised.")
	}
}

*/

/* 
attempted to fix and run 
import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

func main() {
	// Load environment variables from .env file
	err := loadingEnvVariables()
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve the HMAC key from the environment variable
	hmacKey := os.Getenv("HMAC_KEY")
	if hmacKey == "" {
		fmt.Println("HMAC_KEY not found in environment variables")
		return
	}

	// Simulate sending a message from one XBee device to another
	senderMessage := []byte("Hello, receiver!")
	fmt.Println("Sender message:", string(senderMessage))

	// Calculate HMAC for the message
	hmacValue := CalculateHMAC(senderMessage, []byte(hmacKey))
	fmt.Println("Sender HMAC:", hmacValue)

	// Simulate receiving the message and HMAC on the receiver end
	receivedMessage := senderMessage
	receivedHMAC := []byte(hmacValue)

	// Verify the received HMAC
	isValid := VerifyHMAC(receivedMessage, receivedHMAC, []byte(hmacKey))
	if isValid {
		fmt.Println("HMAC verified. Message integrity confirmed.")
	} else {
		fmt.Println("HMAC verification failed. Message integrity compromised.")
	}


// CalculateHMAC calculates HMAC for the given message and key
func CalculateHMAC(message []byte, key []byte) string {
	h := hmac.New(sha256.New, key)
	h.Write(message)
	return hex.EncodeToString(h.Sum(nil))
}

// VerifyHMAC verifies the HMAC for the received message
func VerifyHMAC(message []byte, receivedHMAC []byte, key []byte) bool {
	expectedHMAC := CalculateHMAC(message, key)
	return hmac.Equal(receivedHMAC, []byte(expectedHMAC))
} 
*/
