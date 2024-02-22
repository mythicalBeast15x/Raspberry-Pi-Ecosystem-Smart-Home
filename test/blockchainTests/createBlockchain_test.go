package blockchainTests

import (
	"CMPSC488SP24SecThursday/blockchain"
	"CMPSC488SP24SecThursday/crypto"
	"bytes"
	"testing"
	"time"
)

// testing the block creations
func TestBlockchainCreation(t *testing.T) {

	// Create a new blockchain
	mainBlockchain := blockchain.NewBlockchain(4)

	// Add some blocks to the blockchain
	mainBlockchain.CreateBlock("Block 1 Data")
	mainBlockchain.CreateBlock("Block 2 Data")
	//expected data
	expectedGenData := []byte("Genesis Block")
	expectedBlock1Data := []byte("Block 1 Data")

	//actual data
	genBlockData, genDataError := crypto.Decrypt(mainBlockchain.Chain[0].Data.Ciphertext, []byte("1234567890123456"))
	genBlockHash := mainBlockchain.Chain[0].Hash

	if genDataError != nil {
		t.Errorf("Got error %s", genDataError)
	}
	block1Data, block1DataError := crypto.Decrypt(mainBlockchain.Chain[1].Data.Ciphertext, []byte("1234567890123456"))
	block1PrevHash := mainBlockchain.Chain[1].PrevHash

	if block1DataError != nil {
		t.Errorf("Got error %s", genDataError)
	}

	if !bytes.Equal(expectedGenData, genBlockData) {
		t.Errorf("Expected: %s, Got: %s", expectedGenData, genBlockData)
	}
	if block1PrevHash != genBlockHash {
		t.Errorf("Expected: %s, Got: %s", block1PrevHash, genBlockHash)
	}
	if !bytes.Equal(expectedBlock1Data, block1Data) {
		t.Errorf("Expected: %s, Got: %s", expectedBlock1Data, block1Data)
	}

}

// testing the blockchain validation functions
func TestVerifyBlockChain(t *testing.T) {

	// Create a new blockchain
	mainBlockchain := blockchain.NewBlockchain(4)

	// Add some blocks to the blockchain
	mainBlockchain.CreateBlock("Block 1 Data")
	mainBlockchain.CreateBlock("Block 2 Data")
	mainBlockchain.CreateBlock("Block 3 Data")

	//Test Blocks
	newBlock := blockchain.Block{
		Index:     3,
		Timestamp: time.Now().String(),
		Data: blockchain.EncryptedData{
			Ciphertext: []byte("Block 4 Data"),
			Error:      nil},
		PrevHash: mainBlockchain.Chain[len(mainBlockchain.Chain)-1].Hash,
		Hash:     "",
	}
	newBlock.Hash = blockchain.ProofOfWork(newBlock, 4)
	tamperedBlock := blockchain.Block{
		Index:     1,
		Timestamp: time.Now().String(),
		Data: blockchain.EncryptedData{
			Ciphertext: []byte("Block -5 Data"),
			Error:      nil},
		PrevHash: "0000000000000000000",
		Hash:     "0000a7e8c316d005ef48354a0f3f9b8b78d52103380840b1219b049873b79f3e",
	}

	//actual data
	expectedBlockValidity := mainBlockchain.VerifyBlock(newBlock, 4)
	expectedTamperedValidity := mainBlockchain.VerifyBlock(tamperedBlock, 4)
	expectedBlockchainLength := len(mainBlockchain.Chain)

	//Test for Validating a correct block
	if expectedBlockValidity != true {
		t.Errorf("Expected: %t, Got: %t", true, expectedBlockValidity)
	} else {
		//adding the block to the blockchain
		mainBlockchain.AddToBlockchain(newBlock)
	}
	//test to see if the tampered block is valid
	if expectedTamperedValidity != false {
		t.Errorf("Expected: %t, Got: %t", false, expectedTamperedValidity)
	}
	//Test to see if the blockchain is being updated
	if expectedBlockchainLength != 4 {
		t.Errorf("Expected: %d, Got: %d", 4, expectedBlockchainLength)
	}

}

// testing the VerifyBlockChainWhole functions
func TestVerifyBlockChainWhole(t *testing.T) {
	// Create a new blockchain
	mainBlockchain := blockchain.NewBlockchain(4)

	// Add some blocks to the blockchain
	mainBlockchain.CreateBlock("Block 1 Data")
	mainBlockchain.CreateBlock("Block 2 Data")
	mainBlockchain.CreateBlock("Block 3 Data")

	//test Validity should be true
	blockchainValidity := mainBlockchain.VerifyBlockChain()
	if blockchainValidity != true {
		t.Errorf("Expected: %t, Got: %t", true, blockchainValidity)
	}

	tamperedBlock := blockchain.Block{
		Index:     1,
		Timestamp: time.Now().String(),
		Data: blockchain.EncryptedData{
			Ciphertext: []byte("Block -5 Data"),
			Error:      nil},
		PrevHash: "0000000000000000000",
		Hash:     "0000a7e8c316d005ef48354a0f3f9b8b78d52103380840b1219b049873b79f3e",
	}

	mainBlockchain.AddToBlockchain(tamperedBlock)

	//test Validity should be false
	blockchainValidity = mainBlockchain.VerifyBlockChain()
	if blockchainValidity != true {
		t.Errorf("Expected: %t, Got: %t", false, blockchainValidity)
	}

}

// testing the SerializeBlock and DeserializeBlock functions
func TestBlockSerializationDeserialization(t *testing.T) {
	// Create a new blockchain
	sampleBlockchain := blockchain.NewBlockchain(4)

	//Create a block
	sampleBlockchain.CreateBlock("Block 1 Data")
	block := sampleBlockchain.Chain[1] // Getting the block that's created

	// Serialize the block
	serializedBlock, err := blockchain.SerializeBlock(block)
	if err != nil {
		t.Error("Error:", err)
	}

	// Deserialize the block
	deserializedBlock, err := blockchain.DeserializeBlock(serializedBlock)
	if err != nil {
		t.Error("Error:", err)
	}

	// Check if deserialized block matches the original block
	if !blockchain.EqualBlocks(block, deserializedBlock) {
		t.Error("Block does not match when deserialized")
	}
}
