// TODO: Add test functions

package main

import (
	"CMPSC488SP24SecThursday/blockchain"
	"CMPSC488SP24SecThursday/hashing"
	"CMPSC488SP24SecThursday/messaging"
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Config struct holds the configuration data for the device
type Config struct {
	ID         string                `json:"id"`         // the id of the device
	SecretKey  string                `json:"secret-key"` // the secret key for the network
	AESKey     string                `json:"aes-key"`    // the AES key for the network
	HashKey    string                `json:"hash-key"`   // the hash key for the network
	Appliances []string              `json:"appliances"` // the appliances on the network
	Devices    []string              `json:"devices"`    // the devices on the network
	Blockchain blockchain.Blockchain // the blockchain for the device  // TODO: To be implemented
}

// parseConfigJSON reads a JSON file and parses the data into a Config struct
func parseConfigJSON(filepath string, config *Config) {
	// open JSON file
	jsonFile, err := os.ReadFile(filepath)

	// report errors reading file
	if err != nil {
		fmt.Println("Error reading", filepath, ":", err)
	} else {
		fmt.Println(filepath, "read successfully")
	}

	// parse JSON into Config struct
	err = json.Unmarshal(jsonFile, &config)

	// report errors parsing JSON
	if err != nil {
		fmt.Println("Error parsing", filepath, ":", err)
	} else {
		fmt.Println(filepath, "parsed successfully")
	}
}

// TODO: Remove - Temporary for testing purposes, keys will be stored elsewhere
func generateKeyPair() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}
	publicKey := &privateKey.PublicKey
	return privateKey, publicKey, nil
}

// response is used for confirming the existence of other devices in the network
func response(config *Config, privateKey *rsa.PrivateKey) {
	attempts := 0
	maxAttempts := 3
	sendInterval := time.Second * 10

	// TODO: Remove and update - Temporary for testing purposes, will be in main file
	oOutMessages := &messaging.OpenMessages{}
	qMessages := &messaging.MessageQueue{}

	// Store key within a format for message
	data := map[string]interface{}{
		"Data": privateKey,
	}

	// Make the new message
	// OID 21 = waiting for response?
	messaging.NewMessage(config.ID, "all", "General", "21", data, oOutMessages, qMessages)

	// Dequeue message for setting encryption and hash
	deqMsg, err := qMessages.Dequeue("outgoing")
	if err != nil {
		fmt.Println("Error dequeuing from outgoing queue:", err)
		return
	}
	jsonData, ok := deqMsg.([]byte)
	if !ok {
		fmt.Println("Expected dequeued message to be of type []byte")
		return
	}
	encrypt, err := hashing.Encrypt(jsonData, []byte(config.AESKey))
	if err != nil {
		fmt.Println("Error encrypting JSON:", err)
		return
	}
	hash, _ := hashing.CalculateSHA256(encrypt)
	config.HashKey = hash
	encryptedMsg := messaging.EncryptedMessage{
		EncryptedData: encrypt,
		Hash:          hash,
	}
	encryptPlusHash, err := json.Marshal(encryptedMsg)
	if err != nil {
		return
	}
	fmt.Println("Message to be sent (String):\n ", encryptPlusHash)

	// Start process to check for devices in the network
	for attempts < maxAttempts {
		// networktraffic.Controller()
		time.Sleep(1000 * time.Millisecond)
		if len(qMessages.IncomingMessages) != 0 {
			fmt.Printf("There is another device on the network\n")
			announceStartup(2)
			return
		}
		attempts++
		// Wait for sendInterval before sending the message again
		time.Sleep(sendInterval)
	}
	// Check if it's the only device in the network
	if attempts == maxAttempts {
		fmt.Printf("This device is the only one in the network\n")
		announceStartup(1)
		return
	}
}

// confirmKey confirms the device's secret key with the other devices on the network
func confirmKey(secretKey string) {
	// confirm device ID is unique
} // TODO: Implement this function

// retrieveInfo retrieves the device's information from another device on the network
func retrieveInfo(config *Config) {
	// TODO: Stop other Pis from validating once one begins

	// TODO: Retrieve AES key from another device on the network

	// TODO: Retrieve blockchain from another device on the network

	// TODO: Retrieve device list from another device on the network

	// TODO: Retrieve hash key

	// TODO: Retrieve HMAC key
}

// announceStartup announces the device's startup to the other devices on the network
func announceStartup(caseValue int) {
	switch caseValue {
	case 1:
		// To be implemented
	case 2:
		// To be implemented
	default:
		fmt.Println("Unknown startup case")
	}
} // TODO: Implement this function

// startUp function runs on Pi start up
func startUp() {
	// parse startup config JSON file into config struct
	var config Config
	parseConfigJSON("main/startup.json", &config)

	// Generate temporary key for AES
	aesKey, err := hashing.GenerateRandomKey(16)
	if err != nil {
		fmt.Println("Error generating key for device:", err)
		return
	}
	config.AESKey = string(aesKey)

	// TODO: Remove and update, keys will be added from elsewhere
	privateKey, _, err := generateKeyPair()
	if err != nil {
		fmt.Println("Error generating key pair for device:", err)
		return
	}

	response(&config, privateKey)

	// Account for first Pi on the network
	// 	- generate AES and HMAC key
	// 	- confirm secret key
	// 	- mongoDB
	// 		- blockchain?
	// 		- user IDs?

	// confirmKey(config.SecretKey)

	// retrieve info form another device on the network
	// retrieveInfo(&config)

	// announce successful startup to the network
	// announceStartup()

	// fmt.Println(config) // DEBUG
}
