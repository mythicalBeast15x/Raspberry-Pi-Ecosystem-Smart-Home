package blockchain

import (
	"CMPSC488SP24SecThursday/crypto"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// EncryptedData represents the encrypted data in a block.
type EncryptedData struct {
	Ciphertext []byte
	Error      error
}

// Block represents a single block in the blockchain.
type Block struct {
	Index     int
	Timestamp string
	Data      EncryptedData
	PrevHash  string
	Hash      string
}

// Blockchain is a simple blockchain structure.
type Blockchain struct {
	Difficulty int // the number of leading zeroes required in the hash
	Chain      []Block
}

// CalculateHash calculates the SHA-256 hash of a block. TODO: Refactor?
/*
Purpose: Calculates the hash of a block with sha256.
uses the fields of Block to generate hash of a defined difficulty
	block: Block - input block
	nonce: int - takes input nonce to include nonce in encoded string
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

// ProofOfWork Proof of Work Algorithm Base
/*
Purpose: Helper function//consensus and validation algorithm to ensure a valid hash is generated.
	block: Block - input block
	difficulty: int - difficulty requirement of hash
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
uses the fields of Block to generate hash and verify that both hashes match.
compare both previous block's hash with this block to ensure that they match
verifies block by mining given block and checking the previous block in blockchain
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
	ciphertext, err := crypto.Encrypt([]byte(data), []byte("1234567890123456"))
	newBlock := Block{
		Index:     prevBlock.Index + 1,
		Timestamp: time.Now().String(),
		Data: EncryptedData{
			Ciphertext: ciphertext,
			Error:      err,
		}, // encrypt the data before storing it
		PrevHash: prevBlock.Hash,
		Hash:     "",
	}
	newBlock.Hash = ProofOfWork(newBlock, bc.Difficulty)
	bc.AddToBlockchain(newBlock)
}

// AddToBlockchain adds a new block into the blockchain
/*
Purpose: Adds a given block into the blockchain
	block: Block - input block to be inserted
*/
func (bc *Blockchain) AddToBlockchain(block Block) {
	bc.Chain = append(bc.Chain, block)
}

// NewBlockchain creates a new blockchain with a genesis block.
/*
Purpose: Creates a new Blockchain and initializes it with a genesis block.
	difficulty: determines the difficulty requirement of the hash
	returns: Blockchain
*/
func NewBlockchain(difficulty int) *Blockchain {
	ciphertext, err := crypto.Encrypt([]byte("Genesis Block"), []byte("1234567890123456"))
	genesisBlock := Block{
		Index:     0,
		Timestamp: time.Now().String(),
		Data: EncryptedData{
			Ciphertext: ciphertext,
			Error:      err,
		},
		PrevHash: "",
		Hash:     "",
	}
	genesisBlock.Hash = ProofOfWork(genesisBlock, difficulty)
	return &Blockchain{Chain: []Block{genesisBlock}, Difficulty: difficulty}
}

// SerializeBlock Converting the block struct to JSON
/*
Purpose: Serialize blocks that are created.
	block: block to serialize
	Returns: serialized data of the block
*/
func SerializeBlock(block Block) ([]byte, error) {
	return json.Marshal(block)
}

// DeserializeBlock Converting the JSON string back to a block struct.
/*
Purpose: Deserialize blocks from a JSON file.
	data: data to Deserialize
	Returns: Block and Error value
*/
func DeserializeBlock(data []byte) (Block, error) {
	var block Block
	err := json.Unmarshal(data, &block)
	return block, err
}

// EqualBlocks Check if two blocks are equal
/*
Purpose: To check if all the fields of the two given blocks are the same
	block1, block2: The two block to compare
	Return: bool
*/
func EqualBlocks(block1, block2 Block) bool {
	if block1.Index != block2.Index {
		return false
	}
	if block1.Timestamp != block2.Timestamp {
		return false
	}
	// Compare the Ciphertext field of EncryptedData using bytes.Equal()
	if !bytes.Equal(block1.Data.Ciphertext, block2.Data.Ciphertext) {
		return false
	}
	if block1.PrevHash != block2.PrevHash {
		return false
	}
	if block1.Hash != block2.Hash {
		return false
	}
	return true
}
