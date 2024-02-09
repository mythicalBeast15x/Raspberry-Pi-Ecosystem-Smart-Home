package blockchain

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
	nonce : int - takes input nonce to include nonce in encoded string
	returns: string
*/

func CalculateHash(block Block, nonce int) string {
	hasher := sha256.New()
	var data string   // The data to be hashed
	var result string // The resulting hash

	data = fmt.Sprintf("%d%s%s%s%d", block.Index, block.Timestamp, block.Data, block.PrevHash, nonce)
	hasher.Write([]byte(data))
	result = hex.EncodeToString(hasher.Sum(nil))

	return result
}

// Proof of Work Algorithm Base
/*
Purpose: Helper function//consensus and validation algorithm to ensure a valid hash is generated.
	block : Block - input block
	difficulty : int - difficulty requirement of hash
	returns: string
*/
func ProofOfWork(block Block, difficulty int) string {
	var leadingStr string
	nonce := 0
	hash := CalculateHash(block, nonce)
	leadingStr = hash[:difficulty]
	for {
		// If the hash has the required number of leading zeroes, break the loop, otherwise increment the nonce
		if leadingStr == strings.Repeat("0", difficulty) {
			break
		} else {
			nonce += 1
			hash = CalculateHash(block, nonce)
			leadingStr = hash[:difficulty]
		}
	}
	return hash

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
	result := ProofOfWork(block, difficulty)
	previousHash := bc.Chain[len(bc.Chain)-1].Hash
	return result == block.Hash && previousHash == block.PrevHash
}

// VerifyBlockChain verifies an entire chain using VerifyBlock
/*
Purpose: Uses VerifyBlock (see above) and Cycles through every block on the blockchain. This function
is a prototype for verifying an entire blockchain. This could be run at certain intervals to verify
the integrity and continuity of an entire chain.
***
	--NOTE: This is a Prototype for functionality that might be needed to verify a chain
that already exists--
***
	returns: bool
*/
// verifies the entire chain
func (bc *Blockchain) VerifyBlockChain() bool {

	badBlocks := 0
	for i := 1; i < len(bc.Chain); i++ {
		if !bc.VerifyBlock(bc.Chain[i], bc.Difficulty) {
			badBlocks += 1
		}
	}
	if badBlocks > 0 {
		return true
	} else {
		return false
	}
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
	newBlock.Hash = ProofOfWork(newBlock, bc.Difficulty)
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
	genesisBlock.Hash = ProofOfWork(genesisBlock, difficulty)
	return &Blockchain{Chain: []Block{genesisBlock}, Difficulty: difficulty}
}

//
//func main() {
//	// Create a new blockchain
//	blockchain := NewBlockchain(4)
//
//	// Add some blocks to the blockchain
//	blockchain.CreateBlock("Block 1 Data")
//	blockchain.CreateBlock("Block 2 Data")
//	blockchain.CreateBlock("Block 3 Data")
//
//	fmt.Println(blockchain.Chain[0].Data)
//	fmt.Println(blockchain.Chain[0].Hash)
//	fmt.Println(blockchain.Chain[1].Data)
//	fmt.Println(blockchain.Chain[1].PrevHash)
//	// Print the blockchain
//	blockchainJSON, _ := json.MarshalIndent(blockchain, "", "  ")
//	fmt.Println(string(blockchainJSON))
//
//	// Verify the blockchain  TODO: Move to test module
//	blockchain.Chain[3].Data = "Block 3 Data has been tampered with" // Tamper with the blockchain to test verification
//	VerifyBlockChain(blockchain)
//}

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
