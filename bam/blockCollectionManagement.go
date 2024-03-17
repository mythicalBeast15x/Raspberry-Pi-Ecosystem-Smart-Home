package mongodb_dal

import (
	"CMPSC488SP24SecThursday/blockchain"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

// AddBlock inserts a new user into the database
/*
Purpose: take record of the block in the database
block: the user to add to the database
Return: result and error
*/
func (d *MongoDBDAL) AddBlock(block blockchain.Block) (interface{}, error) {
	// Insert the block into the database
	insertResult, err := d.InsertOne(block)
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
		var block blockchain.Block
		if err := cursor.Decode(&block); err != nil {
			return nil, err
		}
		blockChain = append(blockChain, block)
	}

	// Check if there was an error during cursor iteration
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// Return the list of users
	return blockChain, nil
}
