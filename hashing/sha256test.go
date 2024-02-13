package hashing

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

func TestSHA256Duration() string {
	input := "Hello, World!" // The string you want to hash

	// Record the start time
	startTime := time.Now()

	// Create a new SHA-256 hasher
	hasher := sha256.New()

	// Write the input string to the hasher
	hasher.Write([]byte(input))

	// Calculate the SHA-256 hash
	hash := hasher.Sum(nil)

	// Convert the hash to a hexadecimal string
	hashString := hex.EncodeToString(hash)

	// Record the end time
	endTime := time.Now()

	// Calculate the duration
	duration := endTime.Sub(startTime)

	return fmt.Sprintf("Input: %s\nSHA-256 Hash: %s\nHashing took %s\n", input, hashString, duration)
	/*
		fmt.Printf("SHA-256 Hash: %s\n", hashString)
		fmt.Printf("Hashing took %s\n", duration)
	*/
}
