package hashing

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"time"
)

func TestSHA1Duration() string {
	input := "Hello, World!" // The string you want to hash

	// Record the start time
	startTime := time.Now()

	// Create a new SHA-1 hasher
	hasher := sha1.New()

	// Write the input string to the hasher
	hasher.Write([]byte(input))

	// Calculate the SHA-1 hash
	hash := hasher.Sum(nil)

	// Convert the hash to a hexadecimal string
	hashString := hex.EncodeToString(hash)

	// Record the end time
	endTime := time.Now()

	// Calculate the duration
	duration := endTime.Sub(startTime)

	return fmt.Sprintf("Input: %s\nSHA-1 Hash: %s\nHashing took %s\n", input, hashString, duration)
	/*
		fmt.Printf("SHA-1 Hash: %s\n", hashString)
		fmt.Printf("Hashing took %s\n", duration)
	*/
}
