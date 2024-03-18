package ZigbeeTests

/*
import (
	"CMPSC488SP24SecThursday/blockchain"
	"CMPSC488SP24SecThursday/messaging"
	"CMPSC488SP24SecThursday/networktraffic"
	"fmt"
	"sync"
	"time"
)

func PseudoTest(oMessages *messaging.OpenMessages, qMessages *messaging.MessageQueue) {
	// Create a new blockchain
	mainBlockchain := blockchain.NewBlockchain(4)

	//Test Block as data
	newBlock := blockchain.Block{
		Index:     3,
		Timestamp: time.Now().String(),
		Data: blockchain.EncryptedData{
			Ciphertext: []byte("Block 1 Data"),
			Error:      nil},
		PrevHash: mainBlockchain.Chain[len(mainBlockchain.Chain)-1].Hash,
		Hash:     "",
	}
	// Simulate message being made:
	data := map[string]interface{}{
		"Block": newBlock,
	}
	messaging.NewMessage("Pi-1", "Pi-2", "General", "21", data, oMessages, qMessages)
}

// Function to process the element
func processMessage(msg messaging.Message) {
	// Process the received message
	fmt.Println("\n(===========================================================")
	fmt.Println("\nData to be serviced:\n ", msg)
	fmt.Println("\n===========================================================)")

	// Send next message to the MessageDisperser
	//dal.MessageDisperser(msg, oOutMessages, qMessages)
	// TODO-7: Add additional Parameters/Functionalities to the MessageDisperser before uncommenting
}

func main() {
	//-----------ON START UP--------------

	// Initialize message structures:
	oOutMessages := &messaging.OpenMessages{} // list of messageIDs for outgoing messages
	oInMessages := &messaging.OpenMessages{}  // list of messageIDs for incoming messages
	qMessages := &messaging.MessageQueue{}    // queues for outgoing, incoming, and deserial Messages

	// read in config file
	// TODO-1: Add functionality to read config file

	// set up device info
	// TODO-2: Add functionality to setup device information and update the config file accordingly

	// set encryption and hashing info
	// TODO-3: Add functionality to setup various keys for encryption
	key := []byte("1234567890123456") // TODO-4: Replace with generated AES key

	// set send device id on network
	//	- request data for blockchain
	// TODO-5: Replace with actual initial startup message requests

	PseudoTest(oOutMessages, qMessages) // Fake starter messages being made and going into the outgoing queue

	//constants := Helper.SetConstsants()
	// TODO-6: Add functionality to obtain constants from the helper package

	//-------------------------------------

	// Create a channel to communicate between the goroutine and main
	//incoming := make(chan string)
	outgoing := make(chan string)
	deserial := make(chan messaging.Message)

	// Create a wait group to ensure all goroutines finish before exiting
	var wg sync.WaitGroup

	// Initialize goroutine constantly looking for an incoming message
	go func() {
		for {
			fmt.Println("\nSearching for data...")
			networktraffic.Client(oInMessages, qMessages)

			time.Sleep(3 * time.Second) // Wait for 3 seconds before checking again
		}
	}()

	time.Sleep(8 * time.Second) // Wait for 8 seconds before beginning the next go routine

	// Initialize goroutine constantly looking for an outgoing messages before sending them out
	go func() {
		for {
			if len(qMessages.OutgoingMessages) > 0 {
				msg, err := messaging.EncryptAndHash(qMessages, key)
				if err != nil {
					fmt.Println("Error encrypting and hashing message:", err)
					continue
				}
				outgoing <- msg
			}
			time.Sleep(3 * time.Second) // Wait for 3 seconds before checking again
		}
	}()

	// Initialize goroutine to check for outgoing messages
	go func() {
		for msg := range outgoing {
			fmt.Println("\nSending out data...")
			networktraffic.Controller(msg, oOutMessages, qMessages)
			fmt.Println("\nData Sent!\n ")
		}
	}()

	// Initialize goroutine to check for incoming messages
	go func() {
		for {
			if len(qMessages.IncomingMessages) > 0 {
				err := messaging.ValidateAndDecrypt(oInMessages, qMessages, key)
				if err != nil {
					fmt.Println("Error encrypting and hashing message:", err)
					continue
				}
			}

			if len(qMessages.DeserialMessages) > 0 {
				// If previous function worked, then the new message should now be in the deserial queue
				nextMsg, err := qMessages.Dequeue("deserial")
				if err != nil {
					fmt.Println("\nError dequeuing from deserial queue:", err)
					continue
				}
				// Set the next message to type: Message
				msg := nextMsg.(messaging.Message)

				deserial <- msg
			}
		}
	}()

	// Function to process the elements received from the goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for message := range deserial {
			// Process the element in a separate goroutine
			go processMessage(message)
		}
	}()

	// Wait for all goroutines to finish
	wg.Wait()

	// The go routines above will run for as long as the main function runs.
	// For testing purposes, this script runs for 2 minutes.
	// This is more than enough to attempt to transmit all given message requests.
	// The actual main.go may run forever as long as some flag is true unless an admin-level request changes said flag.
	time.Sleep(2 * time.Minute)
}
*/
