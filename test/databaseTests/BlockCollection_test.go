package databaseTests

import (
	mongodb_dal "CMPSC488SP24SecThursday/bam"
	"CMPSC488SP24SecThursday/blockchain"
	"fmt"
	"log"
	"reflect"
	"testing"
)

func TestBlockCollectionOperation(t *testing.T) {
	// Initialize MongoDBDAL

	// MongoURI MongoDB connection URI
	const MongoURI = "mongodb://localhost:27017"

	const difficulty = 4
	dbName := "test"
	collectionName := "blockchain"
	dal, err := mongodb_dal.NewMongoDBDAL(MongoURI, dbName, collectionName)

	if err != nil {
		log.Fatal("Failed to initialize MongoDBDAL:", err)
	}
	defer dal.Close() // Defer closing the connection

	testBlockChian := blockchain.NewBlockchain(difficulty)

	// Add some blocks to the blockchain
	testBlockChian.CreateBlock("Block 1 Data")
	testBlockChian.CreateBlock("Block 2 Data")

	blockList := mongodb_dal.BlockchainToList(*testBlockChian)

	// Add each block to the database
	for _, block := range blockList {
		mongoBlock := mongodb_dal.ConvertToMongoBlock(block)
		_, err := dal.AddBlock(mongoBlock)
		if err != nil {
			log.Fatal("Failed to add block to database:", err)
		}
	}

	fmt.Println()

	extractedBlockList, err := dal.ExtractBlockInfo()
	if err != nil {
		log.Fatal("Failed to retrieve blocks: ", err)
	}

	CompareBlockLists(t, extractedBlockList, blockList)

	orginalBlockChain := mongodb_dal.ListToBlockchain(blockList, difficulty)
	fromDBBlockChain := mongodb_dal.ListToBlockchain(extractedBlockList, difficulty)

	CompareBlockchains(t, orginalBlockChain, fromDBBlockChain)

	// Call the DeleteAllUsers method to delete all documents in the user collection
	err = dal.DeleteAllUsers()
	if err != nil {
		log.Fatal(err)
	}

}

func CompareBlockLists(t *testing.T, expectedBlocks, actualBlocks []blockchain.Block) {
	// Compare the length of both block lists
	if len(expectedBlocks) != len(actualBlocks) {
		t.Errorf("Block lists have different lengths: expected %d, got %d", len(expectedBlocks), len(actualBlocks))
		return
	}

	// Iterate over each block and compare their attributes
	for i, expectedBlock := range expectedBlocks {
		actualBlock := actualBlocks[i]
		if !reflect.DeepEqual(expectedBlock, actualBlock) {
			t.Errorf("Block %d is different: %+v != %+v", i, expectedBlock, actualBlock)
			return
		}
	}
	t.Logf("Blocks are equal")
}

// CompareBlockchains compares two blockchain instances
func CompareBlockchains(t *testing.T, expectedBlockchain, actualBlockchain blockchain.Blockchain) {
	// Check if the lengths of the blockchains are different
	if len(expectedBlockchain.Chain) != len(actualBlockchain.Chain) {
		t.Errorf("Blockchain lengths are different: expected %d blocks, got %d blocks", len(expectedBlockchain.Chain), len(actualBlockchain.Chain))
		return
	}

	// Iterate over each block in the blockchains and compare them
	for i := 0; i < len(expectedBlockchain.Chain); i++ {
		if !reflect.DeepEqual(expectedBlockchain.Chain[i], actualBlockchain.Chain[i]) {
			t.Errorf("Block %d is different: expected %+v, got %+v", i, expectedBlockchain.Chain[i], actualBlockchain.Chain[i])
			return
		}
	}

	// If all blocks are equal, print success message
	t.Logf("Blockchains are equal")
}
