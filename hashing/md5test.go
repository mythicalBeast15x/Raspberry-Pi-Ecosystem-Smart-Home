package hashing

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

func TestMD5Duration() string {
	input := "Hello, World!" // The string you want to hash

	// Record the start time
	startTime := time.Now()

	// Create an MD5 hasher
	hasher := md5.New()

	// Write the input string to the hasher
	hasher.Write([]byte(input))

	// Calculate the MD5 hash
	hash := hasher.Sum(nil)

	// Convert the hash to a hexadecimal string
	hashString := hex.EncodeToString(hash)

	// Record the end time
	endTime := time.Now()

	// Calculate the duration
	duration := endTime.Sub(startTime)

	return fmt.Sprintf("Input: %s\nMD5 Hash: %s\nHashing took %s\n", input, hashString, duration)

	/*
		fmt.Printf("MD5 Hash: %s\n", hashString)
		fmt.Printf("Hashing took %s\n", duration)
	*/
}
