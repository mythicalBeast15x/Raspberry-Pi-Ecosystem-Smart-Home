package ZigbeeTests

/*
import (
	"CMPSC488SP24SecThursday/messaging"
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


	//Simulating the message being received into the incoming queue
	qMessages.IncomingMessages = append(qMessages.IncomingMessages, outgoingMsg)
	fmt.Print("\n")

	// ------------------- MESSAGE RECEIVED ------------------- //

	// Takes message from incoming queue and runs MessageCheckIn function
	err = messaging.ValidateAndDecrypt(oInMessages, qMessages, key)
	if err != nil {
		return
	}

	nextMsg, err := qMessages.Dequeue("deserial")
	if err != nil {
		fmt.Println("Error dequeuing from outgoing queue:", err)
		return
	}

	// Turning deserialized data back into Message object
	Msg := nextMsg.(messaging.Message)

	// Accessing Message data individually
	fmt.Println()
	fmt.Println(Msg.MessageID)
	fmt.Println(Msg.SenderID)
	fmt.Println(Msg.ReceiverID)
	fmt.Println(Msg.Domain)
	fmt.Println(Msg.Data["ApplianceID"])
}
*/
