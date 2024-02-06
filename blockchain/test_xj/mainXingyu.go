package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

// Block represents a single block in the blockchain.
type Block struct {
	Index      int
	Timestamp  string
	Activities []string
	Hash       string
	PrevHash   string
}

// Blockchain represents the blockchain for the ledger.
type Blockchain struct {
	Chain []Block
	mu    sync.Mutex
}

// CalculateHash calculates the SHA-256 hash of a block.
func CalculateHash(block Block) string {
	hasher := sha256.New()
	data := fmt.Sprintf("%d%s%v%s%s", block.Index, block.Timestamp, block.Activities, block.PrevHash)
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}

// CreateBlock creates a new block in the blockchain.
func (bc *Blockchain) CreateBlock(activities []string) {
	bc.mu.Lock()
	defer bc.mu.Unlock()

	prevBlock := bc.Chain[len(bc.Chain)-1]
	newBlock := Block{
		Index:      prevBlock.Index + 1,
		Timestamp:  time.Now().String(),
		Activities: activities,
		PrevHash:   prevBlock.Hash,
		Hash:       "",
	}

	newBlock.Hash = CalculateHash(newBlock)
	bc.Chain = append(bc.Chain, newBlock)
}

// NewBlockchain creates a new blockchain with a genesis block.
func NewBlockchain() *Blockchain {
	genesisBlock := Block{
		Index:      0,
		Timestamp:  time.Now().String(),
		Activities: []string{"Genesis Activity"},
		PrevHash:   "",
		Hash:       "",
	}
	genesisBlock.Hash = CalculateHash(genesisBlock)
	return &Blockchain{Chain: []Block{genesisBlock}}
}

// DisplayBlockchainAsJSON returns the JSON representation of the blockchain.
func (bc *Blockchain) DisplayBlockchainAsJSON() string {
	bc.mu.Lock()
	defer bc.mu.Unlock()

	jsonData, err := json.MarshalIndent(bc.Chain, "", "  ")
	if err != nil {
		return fmt.Sprintf("Error marshaling JSON: %v", err)
	}

	return string(jsonData)
}

func main() {
	// Create a new blockchain
	blockchain := NewBlockchain()

	// Add some blocks to the blockchain
	blockchain.CreateBlock([]string{"Turn on light 1", "Turn on light 2"})
	blockchain.CreateBlock([]string{"Turn on light 3"})

	// Print the blockchain as JSON
	fmt.Println(blockchain.DisplayBlockchainAsJSON())
}
