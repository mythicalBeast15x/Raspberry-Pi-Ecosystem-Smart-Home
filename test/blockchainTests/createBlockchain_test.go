package blockchainTests

import (
	"CMPSC488SP24SecThursday/blockchain"
	"testing"
	"time"
)

// testing the blocks creation
func TestBlockchainCreation(t *testing.T) {

	// Create a new blockchain
	mainBlockchain := blockchain.NewBlockchain(4)

	// Add some blocks to the blockchain
	mainBlockchain.CreateBlock("Block 1 Data")
	mainBlockchain.CreateBlock("Block 2 Data")
	//expected data
	expectedGenData := "Genesis Block"

	expectedBlock1Data := "Block 1 Data"

	//actual data
	genBlockData := mainBlockchain.Chain[0].Data
	genBlockHash := mainBlockchain.Chain[0].Hash
	black1Data := mainBlockchain.Chain[1].Data
	block1PrevHash := mainBlockchain.Chain[1].PrevHash

	if expectedGenData != genBlockData {
		t.Errorf("Expected: %s, Got: %s", expectedGenData, genBlockData)
	}
	if block1PrevHash != genBlockHash {
		t.Errorf("Expected: %s, Got: %s", block1PrevHash, genBlockHash)
	}
	if expectedBlock1Data != black1Data {
		t.Errorf("Expected: %s, Got: %s", expectedBlock1Data, black1Data)
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
	newBlock := blockchain.Block{3,
		time.Now().String(),
		"Block 4 Data", mainBlockchain.Chain[len(mainBlockchain.Chain)-1].Hash,
		"",
	}
	newBlock.Hash = blockchain.ProofOfWork(newBlock, 4)
	tamperedBlock := blockchain.Block{1,
		time.Now().String(),
		"Block -5 Data", "0000000000000000000",
		"0000a7e8c316d005ef48354a0f3f9b8b78d52103380840b1219b049873b79f3e",
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

	tamperedBlock := blockchain.Block{1,
		time.Now().String(),
		"Block -5 Data", "0000000000000000000",
		"0000a7e8c316d005ef48354a0f3f9b8b78d52103380840b1219b049873b79f3e",
	}

	mainBlockchain.AddToBlockchain(tamperedBlock)

	//test Validity should be false
	blockchainValidity = mainBlockchain.VerifyBlockChain()
	if blockchainValidity != true {
		t.Errorf("Expected: %t, Got: %t", false, blockchainValidity)
	}

}
