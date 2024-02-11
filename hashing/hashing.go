package hashing

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	//"fmt"
	//"os"
	"time"
)

func CalculateMD5(input string) (string, time.Duration) {
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

	return hashString, duration
}

func CalculateSHA1(input string) (string, time.Duration) {
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

	return hashString, duration
}

func CalculateSHA256(input string) (string, time.Duration) {
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

	return hashString, duration
}

func CalculateSHA512(input string) (string, time.Duration) {
	// Record the start time
	startTime := time.Now()

	// Create a new SHA-512 hasher
	hasher := sha512.New()

	// Write the input string to the hasher
	hasher.Write([]byte(input))

	// Calculate the SHA-512 hash
	hash := hasher.Sum(nil)

	// Convert the hash to a hexadecimal string
	hashString := hex.EncodeToString(hash)

	// Record the end time
	endTime := time.Now()

	// Calculate the duration
	duration := endTime.Sub(startTime)

	return hashString, duration
}

/*
func main() {

	fmt.Printf("\n ----------------------------------------")

	input := "Hello, World!"

	hashString, duration := calculateMD5(input)
	fmt.Printf("\nInput: %s\n", input)
	fmt.Printf("MD5 Hash: %s\n", hashString)
	fmt.Printf("Hashing took %s\n", duration)

	hashString1, duration1 := calculateSHA1(input)
	fmt.Printf("\nInput: %s\n", input)
	fmt.Printf("SHA-1 Hash: %s\n", hashString1)
	fmt.Printf("Hashing took %s\n", duration1)

	hashString2, duration2 := calculateSHA256(input)
	fmt.Printf("\nInput: %s\n", input)
	fmt.Printf("SHA-256 Hash: %s\n", hashString2)
	fmt.Printf("Hashing took %s\n", duration2)

	hashString3, duration3 := calculateSHA512(input)
	fmt.Printf("\nInput: %s\n", input)
	fmt.Printf("SHA-512 Hash: %s\n", hashString3)
	fmt.Printf("Hashing took %s\n", duration3)

	fmt.Printf("\n ----------------------------------------")
	fmt.Printf("\n Testing Sample Document (20,000 byte file)\n")
	// Read the content of the sample document
	content, err := os.ReadFile("hashing/hashsample.txt")
	if err != nil {
		fmt.Println("Error reading sample document:", err)
		return
	}

	input1 := string(content)

	// Call the calculateMD5 function
	hashStr, dur := calculateMD5(input1)

	//fmt.Printf("\nInput: %s\n", input)
	fmt.Printf("\nMD5 Hash for doc: %s\n", hashStr)
	fmt.Printf("Hashing took %s\n", dur)

	// Call the calculateSHA1 function
	hashStr1, dur1 := calculateSHA1(input1)

	//fmt.Printf("\nInput: %s\n", input)
	fmt.Printf("\nSHA-1 Hash for doc: %s\n", hashStr1)
	fmt.Printf("Hashing took %s\n", dur1)

	// Call the calculateSHA256 function
	hashStr2, dur2 := calculateSHA256(input1)

	//fmt.Printf("\nInput: %s\n", input)
	fmt.Printf("\nSHA-256 Hash for doc: %s\n", hashStr2)
	fmt.Printf("Hashing took %s\n", dur2)

	// Call the calculateSHA512 function
	hashStr3, dur3 := calculateSHA512(input1)

	//fmt.Printf("\nInput: %s\n", input)
	fmt.Printf("\nSHA-512 Hash for doc: %s\n", hashStr3)
	fmt.Printf("Hashing took %s\n", dur3)

}
*/
