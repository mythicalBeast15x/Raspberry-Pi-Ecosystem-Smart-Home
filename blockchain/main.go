package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"
)

// Block represents a single block in the blockchain.
type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
}

// Blockchain is a simple blockchain structure.
type Blockchain struct {
	Difficulty int // the number of leading zeroes required in the hash
	Chain      []Block
}

// EncryptData encrypts the data in a block. TODO: Implement function
func EncryptData(data string) string {
	return data
}

// CalculateHash calculates the SHA-256 hash of a block. TODO: Refactor?
func CalculateHash(block Block, difficulty int) string {
	hasher := sha256.New()
	nonce := 0
	var data string       // The data to be hashed
	var result string     // The resulting hash
	var leadingStr string // The leading string of the hash (used to check for difficulty)

	// Loop until the hash has the required number of leading zeroes
	for {
		hasher.Reset()
		data = fmt.Sprintf("%d%s%s%s%d", block.Index, block.Timestamp, block.Data, block.PrevHash, nonce)
		hasher.Write([]byte(data))
		result = hex.EncodeToString(hasher.Sum(nil))
		leadingStr = result[:difficulty]

		// If the hash has the required number of leading zeroes, break the loop, otherwise increment the nonce
		if leadingStr == "0000" {
			break
		} else {
			nonce += 1
		}
	}

	return result
}

// VerifyBlock verifies the hash of a block.
func VerifyBlock(block Block, difficulty int) bool {
	result := CalculateHash(block, difficulty)
	return result == block.Hash
}

// CreateBlock creates a new block in the blockchain.
func (bc *Blockchain) CreateBlock(data string) {
	prevBlock := bc.Chain[len(bc.Chain)-1]
	newBlock := Block{
		Index:     prevBlock.Index + 1,
		Timestamp: time.Now().String(),
		Data:      EncryptData(data), // encrypt the data before storing it
		PrevHash:  prevBlock.Hash,
		Hash:      "",
	}
	newBlock.Hash = CalculateHash(newBlock, bc.Difficulty)
	bc.Chain = append(bc.Chain, newBlock)
}

// NewBlockchain creates a new blockchain with a genesis block.
func NewBlockchain(difficulty int) *Blockchain {
	genesisBlock := Block{
		Index:     0,
		Timestamp: time.Now().String(),
		Data:      "Genesis Block",
		PrevHash:  "",
		Hash:      "",
	}
	genesisBlock.Hash = CalculateHash(genesisBlock, difficulty)
	return &Blockchain{Chain: []Block{genesisBlock}, Difficulty: difficulty}
}

func main() {
	// Create a new blockchain
	blockchain := NewBlockchain(4)

	// Add some blocks to the blockchain
	blockchain.CreateBlock("Block 1 Data")
	blockchain.CreateBlock("Block 2 Data")
	blockchain.CreateBlock("Block 3 Data")

	// Print the blockchain
	blockchainJSON, _ := json.MarshalIndent(blockchain, "", "  ")
	fmt.Println(string(blockchainJSON))

	// Verify the blockchain  TODO: Move to test module
	blockchain.Chain[3].Data = "Block 3 Data has been tampered with" // Tamper with the blockchain to test verification
	fmt.Println("\nVerifying blockchain...")
	for i := 1; i < len(blockchain.Chain); i++ {
		if !VerifyBlock(blockchain.Chain[i], blockchain.Difficulty) {
			fmt.Println("Block", i, "is invalid")
		} else {
			fmt.Println("Block", i, "is valid")
		}
	}
}
