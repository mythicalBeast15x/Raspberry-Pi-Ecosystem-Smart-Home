package ZigbeeTests

/*
import (
	"CMPSC488SP24SecThursday/blockchain"
	"CMPSC488SP24SecThursday/messaging"
	"encoding/json"
	"time"
	//"CMPSC488SP24SecThursday/networktraffic"
	"fmt"
)

func main() {
	// Copied from main.go:
	//-----------ON START UP--------------
	// read in config file
	// set up device info
	// set encryption and hashing info
	// set send device id on network
	//	- request data for blockchain
	//constants := Helper.SetConstsants()

	//-------------------------------------
	//listen for messages
	fmt.Println("\n\n -------------------- Setting Up Message Demo -------------------- ")
	// Initialize message queues:
	oOutMessages := &messaging.OpenMessages{} // list of messageIDs for outgoing messages
	oInMessages := &messaging.OpenMessages{}  // list of messageIDs for incoming messages
	qMessages := &messaging.MessageQueue{}    // queues for outgoing, incoming, and incoming-serialized-to-be-serviced message

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

	//networktraffic.Controller(outgoingMsg)

	//	The Controller function should be called with the next in line outgoing message
	//	which should come from the messaging.EncryptAndHash function.

	//networktraffic.Client(qMessages)

	//	The Client function should be called with the same MessageQueue instance
	//	uses for all messages for any one Raspberry PI.

	//	Note: This Client function will probably need to run continuously for the
	//	entirety of the main scripts runtime, as such, using a separate thread may
	//	be needed to accomplish this.
	fmt.Println("\n\n -------------------- Message Being Received By Another PI (Simulation) -------------------- ")
	//Simulating the message being received into the incoming queue
	qMessages.IncomingMessages = append(qMessages.IncomingMessages, outgoingMsg)
	fmt.Print("\n")

	// ------------------- MESSAGE RECEIVED ------------------- //
	fmt.Println("\n\n -------------------- Message being Processed In Reverse For Servicing -------------------- ")
	// Takes message from incoming queue and runs MessageCheckIn function
	err = messaging.ValidateAndDecrypt(oInMessages, qMessages, key)
	if err != nil {
		return
	}

	// If previous function worked, then the new message should now be in the deserial queue
	nextMsg, err := qMessages.Dequeue("deserial")
	if err != nil {
		fmt.Println("Error dequeuing from outgoing queue:", err)
		return
	}

	// Turning deserialized data back into Message object
	Msg := nextMsg.(messaging.Message)
	fmt.Println("\n\n -------------------- Accessing Each Value From The Message -------------------- ")
	// Accessing Message data individually
	fmt.Println(Msg.MessageID)
	fmt.Println(Msg.SenderID)
	fmt.Println(Msg.ReceiverID)
	fmt.Println(Msg.Domain)
	fmt.Println(Msg.Data["Block"])

	// Create a new blockchain for testing
	receiverBlockchain := blockchain.NewBlockchain(4)

	// Add some blocks to the blockchain
	receiverBlockchain.CreateBlock("Block 1 Data")
	receiverBlockchain.CreateBlock("Block 2 Data")
	receiverBlockchain.CreateBlock("Block 3 Data")

	fmt.Println("\n\n -------------------- Accessing Internals Of A Block -------------------- ")

	// Convert Msg.Data["Block"] to JSON bytes
	jsonData, err := json.Marshal(Msg.Data["Block"])
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Unmarshal JSON bytes into blockchain.Block struct
	var receivedBlock blockchain.Block
	err = json.Unmarshal(jsonData, &receivedBlock)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	fmt.Println("Original Block: \n", newBlock)
	fmt.Println("Received Block:\n", receivedBlock)

	fmt.Println("Original Block Data: \n", string(newBlock.Data.Ciphertext))
	fmt.Println("Received Block Data:\n", string(receivedBlock.Data.Ciphertext))
	fmt.Println("\n\n -------------------- Adding Received Block To New Blockchain And Verifying The Updated Chain -------------------- ")
	receiverBlockchain.AddToBlockchain(receivedBlock)

	fmt.Println("Is receiver block chain verified?", receiverBlockchain.VerifyBlockChain())
}
*/
