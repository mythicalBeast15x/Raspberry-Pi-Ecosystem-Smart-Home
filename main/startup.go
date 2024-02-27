// TODO: Add test functions

package main

import (
	"CMPSC488SP24SecThursday/blockchain"
	"encoding/json"
	"fmt"
	"os"
)

// Config struct holds the configuration data for the device
type Config struct {
	ID         int                   `json:"id"`         // the id of the device
	SecretKey  string                `json:"secret-key"` // the secret key for the network
	AESKey     string                // the AES key for the network  TODO: Implement properly
	Blockchain blockchain.Blockchain // the blockchain for the device
	Devices    []string              // the devices on the network  TODO: Implement properly
}

// startUp function runs on Pi start up
func startUp() {
	// parse startup config JSON file into config struct
	var config Config
	parseConfigJSON("main/startup.json", &config)
	fmt.Println(config) // DEBUG

	// confirm secret key
	confirmKey(config.SecretKey)

	// retrieve info form another device on the network
	retrieveInfo(&config)

	// announce successful startup to the network
	announceStartup()

	fmt.Println(config) // DEBUG
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

// confirmKey confirms the device's secret key with the other devices on the network
func confirmKey(secretKey string) {} // TODO: Implement this function

// retrieveInfo retrieves the device's information from another device on the network
func retrieveInfo(config *Config) {
	// TODO: Retrieve AES key from another device on the network
	config.AESKey = "Placeholder AES Key" // Placeholder AES key

	// TODO: Retrieve blockchain from another device on the network
	localBlockchain := blockchain.NewBlockchain(4) // Placeholder blockchain
	config.Blockchain = *localBlockchain           // Placeholder blockchain

	// TODO: Retrieve device list from another device on the network
	config.Devices = []string{"Placeholder Device 1", "Placeholder Device 2"} // Placeholder device list
}

// announceStartup announces the device's startup to the other devices on the network
func announceStartup() {} // TODO: Implement this function
