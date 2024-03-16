package ZigbeeTests

/*
import (
	"CMPSC488SP24SecThursday/messaging"
	"CMPSC488SP24SecThursday/networktraffic"
	"fmt"
	"time"
)

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
	messaging.PseudoTest(oOutMessages, qMessages) // Fake starter messages being made and going into the outgoing queue

	//constants := Helper.SetConstsants()
	// TODO-6: Add functionality to obtain constants from the helper package

	//-------------------------------------

	// Initialize go routine constantly looking for an incoming message
	go func() {
		for {
			fmt.Println("\nSearching for data...")
			networktraffic.Client(oInMessages, qMessages)
		}
	}()

	time.Sleep(8 * time.Second) // Wait for 8 seconds before beginning the next go routine

	// Initialize go routine constantly looking for an outgoing messages before sending them out
	go func() {
		for {
			// Continuously check if the outgoing queue is not empty
			for len(qMessages.OutgoingMessages) == 0 {
				time.Sleep(3 * time.Second) // Wait for 3 seconds before checking again
			}

			// Takes message from outgoing queue, so it can be sent over the Zigbee network
			nextOutgoingMsg, err := messaging.EncryptAndHash(qMessages, key)
			if err != nil {
				continue
			}

			fmt.Println("\nData to be sent:\n ", nextOutgoingMsg)

			fmt.Println("\nSending out data...")
			networktraffic.Controller(nextOutgoingMsg, oOutMessages, qMessages)
			fmt.Println("\nData Sent!\n ")
		}
	}()

	// Initialize go routine constantly looking for an incoming message before reverse processing and servicing
	go func() {
		for {
			// Continuously check if the incoming queue is not empty
			for len(qMessages.IncomingMessages) == 0 {
				time.Sleep(3 * time.Second) // Wait for 3 seconds before checking again
			}
			// Takes message from incoming queue and runs MessageCheckIn function
			err := messaging.ValidateAndDecrypt(oInMessages, qMessages, key)
			if err != nil {
				continue
			}

			// If previous function worked, then the new message should now be in the deserial queue
			nextMsg, err := qMessages.Dequeue("deserial")
			if err != nil {
				fmt.Println("\nError dequeuing from deserial queue:", err)
				continue
			}

			// Set the next message to type: Message
			msg := nextMsg.(messaging.Message)

			fmt.Println("\n(===========================================================")
			fmt.Println("\nData to be serviced:\n ", msg)
			fmt.Println("\n===========================================================)")

			// Send next message to the MessageDisperser
			//dal.MessageDisperser(msg) //, oOutMessages, qMessages)
			// TODO-7: Add additional Parameters/Functionalities to the MessageDisperser before uncommenting
		}
	}()

	// The go routines above will run for as long as the main function runs.
	// For testing purposes, this script runs for 2 minutes.
	// This is more than enough to attempt to transmit all given message requests.
	// The actual main.go may run forever as long as some flag is true unless an admin-level request changes said flag.
	time.Sleep(2 * time.Minute)
}
*/
