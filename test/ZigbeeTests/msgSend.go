package ZigbeeTests

import (
	"CMPSC488SP24SecThursday/blockchain"
	"CMPSC488SP24SecThursday/messaging"
	"CMPSC488SP24SecThursday/networktraffic"
	"fmt"
	"time"
)

func main() {
	//listen for messages
	fmt.Println("\n\n -------------------- Setting Up Message Demo -------------------- ")
	// Initialize message queues:
	oOutMessages := &messaging.OpenMessages{} // list of messageIDs for outgoing messages
	//oInMessages := &messaging.OpenMessages{}  // list of messageIDs for incoming messages
	qMessages := &messaging.MessageQueue{} // queues for outgoing, incoming, and incoming-serialized-to-be-serviced message

	// Create a new blockchain
	mainBlockchain := blockchain.NewBlockchain(4)

	// Add some blocks to the blockchain
	mainBlockchain.CreateBlock("Block 1 Data")
	mainBlockchain.CreateBlock("Block 2 Data")
	mainBlockchain.CreateBlock("Block 3 Data")
	fmt.Println("\n\n -------------------- Creating Block -------------------- ")
	//Test Block as data
	newBlock := blockchain.Block{
		Index:     3,
		Timestamp: time.Now().String(),
		Data: blockchain.EncryptedData{
			Ciphertext: []byte("Block 4 Data"),
			Error:      nil},
		PrevHash: mainBlockchain.Chain[len(mainBlockchain.Chain)-1].Hash,
		Hash:     "",
	}
	fmt.Println("New block \n ", newBlock)
	fmt.Println("\n\n -------------------- Creating Message With Data As the Block -------------------- ")

	// Simulate message being made:
	data := map[string]interface{}{
		"Block": newBlock,
	}
	messaging.NewMessage("Pi-1", "Pi-2", "General", "21", data, oOutMessages, qMessages)

	fmt.Println("Message to be processed (String):\n ", string(qMessages.OutgoingMessages[0]))
	fmt.Println()
	key := []byte("1234567890123456") // TODO - replace with generated AES key
	fmt.Println("\n\n -------------------- Processing Message For Sending -------------------- ")
	// Takes message from outgoing queue, so it can be sent over the Zigbee network
	outgoingMsg, err := messaging.EncryptAndHash(qMessages, key)
	if err != nil {
		return
	}
	fmt.Println("Message to be sent (String):\n ", outgoingMsg)

	networktraffic.Controller(outgoingMsg, oOutMessages, qMessages)
}
