package ZigbeeTests

/*
import (
	"CMPSC488SP24SecThursday/blockchain"
	"CMPSC488SP24SecThursday/messaging"
	"CMPSC488SP24SecThursday/networktraffic"
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	fmt.Println("\n\n -------------------- Setting Up Message Demo -------------------- ")
	// Initialize message queues:
	oOutMessages := &messaging.OpenMessages{} // list of messageIDs for outgoing messages
	oInMessages := &messaging.OpenMessages{}  // list of messageIDs for incoming messages
	qMessages := &messaging.MessageQueue{}    // queues for outgoing, incoming, and incoming-serialized-to-be-serviced message

	key := []byte("1234567890123456") // TODO - replace with generated AES key

	// Go routine constantly looking for an incomming message
	go func() {
		for {
			fmt.Println("Searching for data...")
			networktraffic.Client(oOutMessages, qMessages)
			fmt.Println("Found data!")
		}
	}()

	// Continuously check if the incoming queue is not empty
	for len(qMessages.IncomingMessages) == 0 {
		time.Sleep(1 * time.Second) // Wait for 1 second before checking again
	}

	// ------------------- MESSAGE RECEIVED ------------------- //
	fmt.Println("\n\n -------------------- Message being Processed In Reverse For Servicing -------------------- ")
	// Takes message from incoming queue and runs MessageCheckIn function
	err := messaging.ValidateAndDecrypt(oInMessages, qMessages, key)
	if err != nil {
		return
	}

	// If previous function worked, then the new message should now be in the deserial queue
	nextMsg, err := qMessages.Dequeue("deserial")
	if err != nil {
		fmt.Println("Error dequeuing from deserial queue:", err)
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

	fmt.Println("Received Block:\n", receivedBlock)
	fmt.Println("Received Block Data:\n", string(receivedBlock.Data.Ciphertext))

	fmt.Println("\n\n -------------------- Adding Received Block To New Blockchain And Verifying The Updated Chain -------------------- ")
	receiverBlockchain.AddToBlockchain(receivedBlock)

	fmt.Println("Is receiver block chain verified?", receiverBlockchain.VerifyBlockChain())
}
*/
