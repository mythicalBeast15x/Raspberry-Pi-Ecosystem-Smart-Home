package mongodb_dal

import (
	"CMPSC488SP24SecThursday/blockchain"
	"context"
	"encoding/base64"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

type MongoBlock struct {
	Index     int    `bson:"index"`
	Timestamp string `bson:"timestamp"`
	Data      []byte `bson:"data"`
	PrevHash  string `bson:"prev_hash"`
	Hash      string `bson:"hash"`
}

// AddBlock inserts a new user into the database
/*
Purpose: take record of the block in the database
block: the user to add to the database
Return: result and error
*/
func (d *MongoDBDAL) AddBlock(block MongoBlock) (interface{}, error) {
	// Convert the MongoBlock struct to a BSON document
	blockDocument, err := bson.Marshal(block)
	if err != nil {
		return nil, err
	}
	// Insert the block into the database
	insertResult, err := d.Collection.InsertOne(context.Background(), blockDocument)
	if err != nil {
		return nil, err
	}
	return insertResult.InsertedID, nil
}

// ExtractBlockInfo retrieves all the block database and returns them as a slice
/*
Purpose: get a copy of the blockchain from the database
Return: List of users and error
*/
func (d *MongoDBDAL) ExtractBlockInfo() ([]blockchain.Block, error) {
	// Initialize a slice to hold the data
	var blockChain []blockchain.Block

	// Query all documents from the collection
	cursor, err := d.Collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	// Iterate over the cursor and append each document to the slice
	for cursor.Next(context.Background()) {
		var block MongoBlock
		if err := cursor.Decode(&block); err != nil {
			return nil, err
		}
		blockchainBlock := ConvertToBlockchainBlock(block)
		blockChain = append(blockChain, blockchainBlock)
	}

	// Check if there was an error during cursor iteration
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// Return the list of blocks
	return blockChain, nil
}

// ConvertToBlockchainBlock from the database back into blockchain blocks.
func ConvertToBlockchainBlock(mongoBlock MongoBlock) blockchain.Block {
	blockchainBlock := blockchain.Block{
		Index:     mongoBlock.Index,
		Timestamp: mongoBlock.Timestamp,
		Data:      deserializeEncryptedData(mongoBlock.Data),
		PrevHash:  mongoBlock.PrevHash,
		Hash:      mongoBlock.Hash,
	}
	return blockchainBlock
}

func ConvertToMongoBlock(block blockchain.Block) MongoBlock {
	mongoBlock := MongoBlock{
		Index:     block.Index,
		Timestamp: block.Timestamp,
		Data:      serializeEncryptedData(block.Data),
		PrevHash:  block.PrevHash,
		Hash:      block.Hash,
	}
	return mongoBlock
}

// BlockchainToList converts a blockchain to a list of blocks.
func BlockchainToList(bc blockchain.Blockchain) []blockchain.Block {
	// Create an empty list to store blocks
	blockList := make([]blockchain.Block, 0)

	// Iterate through each block in the blockchain
	for _, block := range bc.Chain {
		// Append the block to the list
		blockList = append(blockList, block)
	}

	return blockList
}

// ListToBlockchain converts a list of blocks back to a blockchain.
func ListToBlockchain(blocks []blockchain.Block, difficulty int) blockchain.Blockchain {
	// Ensure the blocklist is not empty
	if len(blocks) == 0 {
		bc := blockchain.NewBlockchain(difficulty)
		return *bc // Return an empty blockchain
	}

	// Use the first block in the list as the genesis block
	genesisBlock := blocks[0]

	// Create a new blockchain with the given difficulty and the genesis block
	bc := blockchain.Blockchain{
		Difficulty: difficulty,
		Chain:      []blockchain.Block{genesisBlock},
	}

	// Add each subsequent block from the list to the blockchain
	for _, block := range blocks[1:] {
		bc.AddToBlockchain(block)
	}
	return bc
}

func serializeEncryptedData(data blockchain.EncryptedData) []byte {
	return []byte(base64.StdEncoding.EncodeToString(data.Ciphertext))
}

func deserializeEncryptedData(data []byte) blockchain.EncryptedData {
	// Deserialize BSON-compatible data back into EncryptedData struct
	// For example, decode base64 string into []byte
	ciphertext, _ := base64.StdEncoding.DecodeString(string(data))
	return blockchain.EncryptedData{
		Ciphertext: ciphertext,
		Error:      nil, // Error field may not be needed in this context
	}
}

// DeleteAllBlocks to remove all blocks from the database
/*
Purpose: to manage a database
return: error
*/
func (d *MongoDBDAL) DeleteAllBlocks() error {
	// Define an empty filter to match all documents
	filter := bson.D{}

	// Perform the delete operation
	deleteResult, err := d.Collection.DeleteMany(context.Background(), filter)
	if err != nil {
		return err
	}

	// Print delete result
	fmt.Println("Deleted", deleteResult.DeletedCount, "documents")

	return nil
}
