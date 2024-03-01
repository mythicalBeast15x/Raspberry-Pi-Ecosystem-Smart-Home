package ZigbeeTests

/*
import (
	"CMPSC488SP24SecThursday/messaging"
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

	// Initialize message queues:
	oOutMessages := &messaging.OpenMessages{} // list of messageIDs for outgoing messages
	oInMessages := &messaging.OpenMessages{}  // list of messageIDs for incoming messages
	qMessages := &messaging.MessageQueue{}    // queues for outgoing, incoming, and incoming-serialized-to-be-serviced message

	// Simulate message being made:
	data := map[string]interface{}{
		"ApplianceID": "lamp-1",
	}
	messaging.NewMessage("Pi-1", "Pi-2", "Lighting", "2", data, oOutMessages, qMessages)

	key := []byte("1234567890123456") // TODO - replace with generated AES key

	// Takes message from outgoing queue and sends out message to Zigbee function
	incomingMsg, err := messaging.EncryptAndHash(qMessages, key)
	if err != nil {
		return
	}
	fmt.Println("Message to be sent (String):\n ", incomingMsg)

	qMessages.IncomingMessages = append(qMessages.IncomingMessages, incomingMsg)
	fmt.Print("\n")
	// Takes message from incoming queue and runs MessageCheckIn function
	err = messaging.ValidateAndDecrypt(oInMessages, qMessages, key)
	if err != nil {
		return
	}
}
*/
