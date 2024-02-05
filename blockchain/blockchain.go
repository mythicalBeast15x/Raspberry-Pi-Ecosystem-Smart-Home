package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
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

// EncryptData encrypts the data in a block.
//TODO: Implement function
/*
Purpose: To Encrypt Data for Block Creation
	data: string - data to be encrypted
	returns: string
*/
func EncryptData(data string) string {
	return data
}

// CalculateHash calculates the SHA-256 hash of a block. TODO: Refactor?
/*
Purpose: Calculates the hash of a block with sha256.
uses the fields of Block to generate hash of a defined difficulty
	block : Block - input block
	difficulty : int - difficulty requirement of hash
	returns: string
*/
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
		if leadingStr == strings.Repeat("0", difficulty) {
			break
		} else {
			nonce += 1
		}
	}

	return result
}

// VerifyBlock verifies the hash of a block.
/*
Purpose: Verifies given Block data has not been tampered.
uses the fields of Block to generate hash and verifies that both hashes match.
compares both previous block's hash with this block to ensure that they match
verifies block by mining given block and checking previous block in blockchain
***
	--NOTE: The block must not be in the blockchain
	This Function should be called before the block is added to the Blockchain--
***
	block : Block - input block
	difficulty : int - difficulty requirement of hash
	returns: string
*/
func (bc *Blockchain) VerifyBlock(block Block, difficulty int) bool {
	result := CalculateHash(block, difficulty)
	previousHash := bc.Chain[len(bc.Chain)-1].Hash
	return result == block.Hash && previousHash == block.PrevHash
}

// CreateBlock creates a new block in the blockchain.
/*
Purpose: Creates a new block and adds it into the blockchain
creates a new block based on the given data, time, and previous hash and calculated hash
encrypts hash with EncryptData()
calculates hash with CalculateHash()
inserts the block into the blockchain with AddToBlockchain()
	data : string - input data for hash
*/
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
	bc.AddToBlockchain(newBlock)
}

// AddToBlockchain adds a new block into the blockchain
/*
Purpose: Adds a given block into the blockchain
	block : Block - input block to be inserted
*/
func (bc *Blockchain) AddToBlockchain(block Block) {
	bc.Chain = append(bc.Chain, block)
}

// NewBlockchain creates a new blockchain with a genesis block.
/*
Purpose: Creates a new Blockchain and initializes it with a genesis block.
	difficulty : determines the difficulty requirement of the hash
	returns: Blockchain
*/
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

/*
func main() {
	// Create a new blockchain
	blockchain := NewBlockchain(4)

	// Add some blocks to the blockchain
	blockchain.CreateBlock("Block 1 Data")
	blockchain.CreateBlock("Block 2 Data")
	blockchain.CreateBlock("Block 3 Data")

	fmt.Println(blockchain.Chain[0].Data)
	fmt.Println(blockchain.Chain[0].Hash)
	fmt.Println(blockchain.Chain[1].Data)
	fmt.Println(blockchain.Chain[1].PrevHash)
	// Print the blockchain
	blockchainJSON, _ := json.MarshalIndent(blockchain, "", "  ")
	fmt.Println(string(blockchainJSON))

	// Verify the blockchain  TODO: Move to test module
	blockchain.Chain[3].Data = "Block 3 Data has been tampered with" // Tamper with the blockchain to test verification
	fmt.Println("\nVerifying blockchain...")
	for i := 1; i < len(blockchain.Chain); i++ {
		if !blockchain.VerifyBlock(blockchain.Chain[i], blockchain.Difficulty) {
			fmt.Println("Block", i, "is invalid")
		} else {
			fmt.Println("Block", i, "is valid")
		}
	}
}
*/
