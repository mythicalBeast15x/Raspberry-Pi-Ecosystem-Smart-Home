package messaging

/*
This script contains the foundation of a data structure for creating messages to be sent out to
relevant device nodes.
*/

import (
	"CMPSC488SP24SecThursday/Helper"
	"CMPSC488SP24SecThursday/hashing"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

// globalPiID temp variable simulates the ID representing a message intended for all devices.
var globalPiID = "all"

// localPiID temp variable simulates the ID of the Pi device to show a case of the receiverID matching the global Pi ID.
var localPiID = "Pi-2"

// EncryptedMessage struct to hold data
type EncryptedMessage struct {
	EncryptedData string `json:"encryptedData"`
	Hash          string `json:"hash"`
}

// MessageQueue is a list of messages for managing incoming, outgoing, and deserialized versions of our system's messages.
/*
outgoing messages are in the form of the serialized JSON structure of type []byte.
incoming messages are encryption+hash versions of the serialized JSON structure of type []byte.
deserial messages are for the deserialized versions of the data of type Message so that its contents can be checked before
being serviced.
*/
type MessageQueue struct {
	IncomingMessages []string  // Queue for incoming encrypted messages as strings along with their hash.
	DeserialMessages []Message // Queue for post-processed messages that are waiting to be checked-in before servicing.
	OutgoingMessages [][]byte  // Queue for outgoing messages
}

// OpenMessages is a list to keep track of all the open messages
/*If a message is sent, it may receive multiple copies of the response. This struct keeps track of all the messages and can
send new messages if old ones haven't been responded to. It works in tandem with generateMessageID and resolveMessage
to keep messageIDs unique as well as ensuring that responses are handled only once*/
type OpenMessages struct {
	messages []string
}

// Message represents the structure of a message object.
type Message struct {
	MessageID   string                 `json:"messageID"`
	SenderID    string                 `json:"senderID"`
	ReceiverID  string                 `json:"receiverID"`
	Domain      string                 `json:"domain"`
	OperationID string                 `json:"operationID"`
	Data        map[string]interface{} `json:"Data"`
}

// generateMessageID creates a message ID generates a unique messageID for a message to be sent
/*This keeps allows the messages sent to be unique and have a signature
 */
func (oMessages *OpenMessages) generateMessageID() string {
	code := "NULL"
	for code == "NULL" {
		constants := Helper.SetConstants()
		randomBytes := make([]byte, constants.MESSAGE_ID_LENGTH/2) // Generate random bytes (half the length for hex encoding)
		_, err := rand.Read(randomBytes)
		if err != nil {
			return "ERROR"
		}
		x := 0
		code, err = hex.EncodeToString(randomBytes), nil
		if err != nil {
			return "ERROR"
		}
		for x < len(oMessages.messages) && oMessages.messages[x] != code {
			x++
		}
		if x == len(oMessages.messages) {
			oMessages.messages = append(oMessages.messages, code)
		} else {
			code = "NULL"
		}
	}
	return code
}

// resolveMessage removes a message from the OpenMessages struct
/* Removes a MessageID from OpenMessages when a response is received
responseID : string
Removes message from deserialized queue of to-be-serviced messages
*/
func (oMessages *OpenMessages) resolveMessage(responseID string, qMessages *MessageQueue) {
	x := 0
	for x < len(oMessages.messages) && oMessages.messages[x] != responseID {
		x++
	}
	if x < len(oMessages.messages) {
		oMessages.messages = append(oMessages.messages[:x], oMessages.messages[x+1:]...)
	}
	// Remove the message from the DeserialMessages queue
	for i, msg := range qMessages.DeserialMessages {
		if msg.MessageID == responseID {
			qMessages.DeserialMessages = append(qMessages.DeserialMessages[:i], qMessages.DeserialMessages[i+1:]...)
			break
		}
	}
}

// AddMessageID manually adds a message ID to the OpenMessages struct.
/*
Purpose: Manually adds a message ID to the OpenMessages struct for testing purposes.
Data:
- messageID (string): The message ID to be added to the OpenMessages struct.
- oMessages: Pointer to the OpenMessages struct where the message ID will be added.
Returns: None
*/
func (oMessages *OpenMessages) AddMessageID(messageID string) {
	oMessages.messages = append(oMessages.messages, messageID)
}

// NewMessage creates a new Message object with the given parameters.
/*
Purpose: Creates a new message object, serializes it, and adds it to a queue of outgoing messages
Data:
- MessageID: The unique ID of the message
- SenderID: The ID of the sending PI device.
- ReceiverID: The ID of the target PI device.
- Domain: The domain of the message (e.g., lighting, HVAC, security).
- OperationID: The ID of the specific operation that the message wants to invoke.
- Data: Any additional data/parameters that the message needs to service the operation.
- oMessages: A list to keep track of opened messages for the sake if giving the message a unique MessageID.
- qMessages: A queue to add the serialized message too.
Returns: Message: The newly created Message object.
*/
func NewMessage(senderID, receiverID, domain, operationID string, data map[string]interface{}, oMessages *OpenMessages, qMessages *MessageQueue) Message {
	messageID := oMessages.generateMessageID()
	message := Message{
		MessageID:   messageID,
		SenderID:    senderID,
		ReceiverID:  receiverID,
		Domain:      domain,
		OperationID: operationID,
		Data:        data,
	}
	// Serialize the message into a JSON structure
	jsonData, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return message
	}
	// Add serialized outgoing message to the outgoing queue
	qMessages.OutgoingMessages = append(qMessages.OutgoingMessages, jsonData)
	// Returns the un-serialized version of the data.
	return message
}

// DisplayMessage prints out the Json data.
/*
Purpose: Shows the data in the output window for testing purposes.
Data:
- msg: The Message object to be displayed.
Prints:
- MessageID: The unique ID of the message
- SenderID: The sender PI's ID of the message.
- ReceiverID: The receiver PI's ID of the message.
- Domain: The domain of the message.
- OperationID: The operation ID of the message.
- Data: The additional data given to the message.
Returns: None
*/
func DisplayMessage(msg Message) {
	fmt.Println("MessageID:", msg.MessageID)
	fmt.Println("SenderID:", msg.SenderID)
	fmt.Println("ReceiverID:", msg.ReceiverID)
	fmt.Println("Domain:", msg.Domain)
	fmt.Println("OperationID:", msg.OperationID)
	fmt.Println("Data:", msg.Data)     // prints mapped version of data section
	for key, value := range msg.Data { // prints key value pairs of the data section
		// 'key' is the key of the current item in the map
		// 'value' is the value associated with the current key
		fmt.Printf("Key: %s, Value: %v\n", key, value)
	}
}

// MessageCheckIn checks the IDs of a given message.
/*
Purpose: Deserializes a JSON message and checks to see if the IDs are valid for the request to be serviced.
Data :
- msg: The Message object to be checked.
- oMessages: A list to keep track of opened messages for the sake if giving the message a unique MessageID.
- qMessages: A queue to add the serialized message too.
Prints:
- Case 1: MessageID already exists. This means it is a duplicate message.
- Case 2: ReceiverID matches the Global Pi ID meaning that the message is meant for all devices and should be serviced and echoed.
- Case 3: ReceiverID does not match the global Pi ID. This means that the message should be echoed.
- Case 4: Both previous checks pass and the message can be serviced.
Returns: bool
*/
func MessageCheckIn(serialMsg []byte, oMessages *OpenMessages, qMessages *MessageQueue) bool {
	var msg Message
	msg = deserial(serialMsg)
	// Check if the messageID already exists in the OpenMessages struct
	for _, existingMsgID := range oMessages.messages {
		if existingMsgID == msg.MessageID {
			fmt.Println("Message ID already exists in OpenMessages struct")
			return false
		}
	}
	if msg.ReceiverID == globalPiID {
		fmt.Println("Global ID matched! Message is meant for all devices")
		qMessages.DeserialMessages = append(qMessages.DeserialMessages, msg)
		oMessages.messages = append(oMessages.messages, msg.MessageID)
		EchoMessage(msg, qMessages) // call to echo the message out again
		return true
	}
	// Check if the receiverID matches the global Pi ID
	if msg.ReceiverID != localPiID {
		fmt.Println("Receiver ID does not match the local Pi ID")
		EchoMessage(msg, qMessages) // call to echo the message out again
		return false
	}
	// If MessageID and ReceiverID checks pass, then it can be serviced. Adds the message of type Message to the deserialized message queue.
	fmt.Println("Message is serviceable!")
	qMessages.DeserialMessages = append(qMessages.DeserialMessages, msg)
	oMessages.messages = append(oMessages.messages, msg.MessageID)
	return true
}

// deserial deserializes the Json data.
/*
Purpose: Takes in the serialized data and outputs the deserialized version of it.
Data:
- serialMsg: The JSON structure of type []byte to be deserialized.
Returns: msg of type Message.
*/
func deserial(serialMsg []byte) Message {
	var msg Message
	err := json.Unmarshal(serialMsg, &msg)
	if err != nil {
		fmt.Println("Error:", err)
		return Message{}
	}
	return msg
}

// EchoMessage repeats the message by sending it to the outgoing queue.
/*
Purpose: Serializes the message into JSON and adds it to the outgoing message queue.
Parameters:
- msg: The message of type Message to be echoed.
- qMessages: Pointer to the MessageQueue where the outgoing messages are stored.
*/
func EchoMessage(msg Message, qMessages *MessageQueue) {
	// Serialize the message into JSON
	jsonData, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	// Add the serialized message to the outgoing queue
	qMessages.OutgoingMessages = append(qMessages.OutgoingMessages, jsonData)
}

// Dequeue pops the first element from the specified queue from the MessageQueue structure.
/*
Purpose: To allow for the next in-line message to be taken from any desired messaging queue and to then save that item
to a variable for further processing or servicing.
Data:
- queueType (string): Specifies the type of queue to dequeue from. Valid values are:
  - "incoming": Dequeues from the incoming messages queue.
  - "deserial": Dequeues from the deserialized messages queue.
  - "outgoing": Dequeues from the outgoing messages queue.
Returns:
- interface{}: The dequeued element from the top of the specified queue.
- error: An error if the specified queue is empty or if an unknown queue type is provided.
*/
func (msgQueue *MessageQueue) Dequeue(queueType string) (interface{}, error) {
	switch queueType {
	case "incoming":
		if len(msgQueue.IncomingMessages) == 0 {
			return nil, fmt.Errorf("incoming queue is empty")
		}
		msg := msgQueue.IncomingMessages[0]
		msgQueue.IncomingMessages = msgQueue.IncomingMessages[1:]
		return msg, nil
	case "deserial":
		if len(msgQueue.DeserialMessages) == 0 {
			return nil, fmt.Errorf("deserial queue is empty")
		}
		msg := msgQueue.DeserialMessages[0]
		msgQueue.DeserialMessages = msgQueue.DeserialMessages[1:]
		return msg, nil
	case "outgoing":
		if len(msgQueue.OutgoingMessages) == 0 {
			return nil, fmt.Errorf("outgoing queue is empty")
		}
		msg := msgQueue.OutgoingMessages[0]
		msgQueue.OutgoingMessages = msgQueue.OutgoingMessages[1:]
		return msg, nil
	default:
		return nil, fmt.Errorf("unknown queue type: %s", queueType)
	}
}

// EncryptAndHash dequeues a message from the outgoing queue and encrypts it.
/*
Purpose: To dequeue a serialized message from the outgoing queue, encrypt it, generate a hash for it, combine those
two pieces of data into one object, and call the Zigbee sender function to send out the processed message out to the
Zigbee network.
Data:
- qMessages: Pointer to the MessageQueue where the outgoing messages are stored.
- key: The AES encryption key of type []byte
Returns:
- A JSON string to be sent out over the Zigbee network.
- error: An error if the specified queue is empty.
*/
func EncryptAndHash(qMessages *MessageQueue, key []byte) (string, error) {
	// Dequeue a message from the outgoing queue
	deqMsg, err := qMessages.Dequeue("outgoing")
	if err != nil {
		fmt.Println("Error dequeuing from outgoing queue:", err)
		return "", nil
	}
	fmt.Println("Dequeued from outgoing queue: Some byte message")
	// Convert the dequeued message to []byte
	jsonData, ok := deqMsg.([]byte)
	if !ok {
		fmt.Println("Expected dequeued message to be of type []byte")
		return "", nil
	}
	// Encrypt the JSON data
	encrypt, err := hashing.Encrypt(jsonData, key)
	if err != nil {
		fmt.Println("Error encrypting JSON:", err)
		return "", nil
	}
	fmt.Println("Encrypted message:", encrypt)

	// Generate the HMAC hash and append it to the encrypted data:
	hash, time := hashing.CalculateSHA256(encrypt)
	fmt.Println("Time taken: ", time)

	// Create the EncryptedMessage struct
	encryptedMsg := EncryptedMessage{
		EncryptedData: encrypt,
		Hash:          hash,
	}

	// Marshal the EncryptedMessage struct into JSON
	encryptPlusHash, err := json.Marshal(encryptedMsg)
	if err != nil {
		return "", fmt.Errorf("error marshalling encrypted message to JSON: %v", err)
	}

	return string(encryptPlusHash), err
}

// ValidateAndDecrypt dequeues a message from the incoming queue, validates it and decrypts it.
/*
Purpose: To dequeue an encrypted message from the incoming queue, regenerate a hash for it, validate the new hash
against the accompanying hash, decrypt it, and send the decrypted data to the message check in function.
Data:
- oMessages: A list to keep track of opened messages for the sake if giving the message a unique MessageID.
- qMessages: Pointer to the MessageQueue where the incoming messages are stored.
- key: The AES encryption key of type []byte
Returns:
- error: An error if the specified queue is empty.
*/
func ValidateAndDecrypt(oMessages *OpenMessages, qMessages *MessageQueue, key []byte) error {
	// Dequeue a JSON string message from the incoming queue
	deqMsg, err := qMessages.Dequeue("incoming")
	if err != nil {
		fmt.Println("Error dequeuing from incoming queue:", err)
		return err
	}
	fmt.Println("Dequeued from incoming queue: Some string message")

	// Convert the dequeued message to string
	jsonStr, ok := deqMsg.(string)
	if !ok {
		return fmt.Errorf("expected dequeued message to be of type string")
	}

	// Unmarshal the JSON string into EncryptedMessage struct
	var encryptedPlusHash EncryptedMessage
	err = json.Unmarshal([]byte(jsonStr), &encryptedPlusHash)
	if err != nil {
		return fmt.Errorf("error unmarshalling JSON message: %v", err)
	}

	// Split the encrypted message from the hash and validate the HMAC hash to see if the message was tampered with:
	hash, time := hashing.CalculateSHA256(encryptedPlusHash.EncryptedData)
	fmt.Println("Time taken: ", time)
	if hash == encryptedPlusHash.Hash { // if the generated hash matches the appended hash:
		// Decrypt the JSON data
		decrypt, err := hashing.Decrypt(encryptedPlusHash.EncryptedData, key)
		if err != nil {
			fmt.Print("Error decrypting JSON:", err)
			return err
		}
		// Use the encrypted message as needed
		fmt.Println("Decrypted message:", decrypt)

		// Add decrypted data to message check in function
		decryptedByte := decrypt
		MessageCheckIn(decryptedByte, oMessages, qMessages)
		return err
	}
	return nil
}

/*
func main() {
	// Initialize message queues
	oOutMessages := &OpenMessages{} // list of messageIDs for outgoing messages
	oInMessages := &OpenMessages{}  // list of messageIDs for incoming messages
	qMessages := &MessageQueue{}    // queues for outgoing, incoming, and incoming-serialized-to-be-serviced messages

	// Create and display first message
	data := map[string]interface{}{
		"ApplianceID": "lamp-1",
		"Brightness":  70,
	}
	msg := NewMessage("Pi-1", "Pi-3", "Lighting", "2", data, oOutMessages, qMessages)
	fmt.Println("------------------------------------------------------------------------------------------------------------------------\nDisplayed Message One:\n ")
	DisplayMessage(msg)

	// Display message ID list
	fmt.Printf("\nOutgoing MessageID list: %s", oOutMessages)
	fmt.Println()

	// Create and display second message
	msg2 := NewMessage("Pi-1", "Pi-2", "Lighting", "2", map[string]interface{}{
		"ApplianceID": "lamp-1",
		"Brightness":  70,
	}, oOutMessages, qMessages)
	fmt.Println("------------------------------------------------------------------------------------------------------------------------\nDisplayed Message Two:\n ")
	DisplayMessage(msg2)

	// Display message ID list
	fmt.Printf("\nOutgoing MessageID list: %s", oOutMessages)
	fmt.Println()

	// Create and display third message
	msg3 := NewMessage("Pi-1", "all", "HVAC", "2", map[string]interface{}{
		"ApplianceID": "fan-1",
		"speed":       100,
	}, oOutMessages, qMessages)
	fmt.Println("------------------------------------------------------------------------------------------------------------------------\nDisplayed Message Three:\n ")
	DisplayMessage(msg3)

	// Display message ID list
	fmt.Printf("\nOutgoing MessageID list: %s", oOutMessages)
	fmt.Println()

	// Create and display third message
	msg4 := NewMessage("Pi-1", "Pi-3", "Lighting", "2", map[string]interface{}{
		"ApplianceID": "lamp-1",
		"Brightness":  70,
	}, oOutMessages, qMessages)
	fmt.Println("------------------------------------------------------------------------------------------------------------------------\nDisplayed Message Four:\n ")
	DisplayMessage(msg4)

	// Display message ID list
	fmt.Printf("\nOutgoing MessageID list: %s", oOutMessages)
	fmt.Println()

	// Display the first element of the outgoing messages queue
	fmt.Println("------------------------------------------------------------------------------------------------------------------------\nDisplayed Outgoing queue:\n ")
	fmt.Printf("First in line in queue: %s", qMessages.OutgoingMessages[0])
	fmt.Printf("Second in line in queue: %s", qMessages.OutgoingMessages[1])
	fmt.Printf("Third in line in queue: %s", qMessages.OutgoingMessages[2])

	// Setting each message in the message list to a variable
	test := qMessages.OutgoingMessages[0]
	test1 := qMessages.OutgoingMessages[1]
	test2 := qMessages.OutgoingMessages[2]
	test3 := qMessages.OutgoingMessages[3]

	// Adding a couple of the outgoing messages to the incoming message queue for demo purposes
	qMessages.IncomingMessages = append(qMessages.IncomingMessages, test)
	qMessages.IncomingMessages = append(qMessages.IncomingMessages, test1)

	fmt.Println()
	fmt.Println("------------------------------------------------------------------------------------------------------------------------\nTesting message check-in:\n ")

	// Testing if echo function works correctly
	fmt.Println("Length of Outgoing Queue before test:", len(qMessages.OutgoingMessages))
	fmt.Println("Removing element from outgoing queue for testing purposes...")
	// Dequeue from the outgoing messages queue
	_, err := qMessages.Dequeue("outgoing")
	if err != nil {
		fmt.Println("Error dequeuing from outgoing queue:", err)
	} else {
		fmt.Println("Dequeued from outgoing queue: Some byte message")
	}
	fmt.Println("Length of Outgoing Queue:", len(qMessages.OutgoingMessages))
	// Test message check-in function with first message
	MessageCheckIn(test, oInMessages, qMessages)
	// Since the receiverID is not matched, the outgoing queue should have the message added to it.
	fmt.Println("Length of Outgoing Queue:", len(qMessages.OutgoingMessages))
	fmt.Println("NOTE: Since the ReceiverID did not match, the message was placed into the \noutgoing queue via echo function as seen with the length of the outgoing queue \nbeing upped by one.")
	fmt.Println()

	// Test message check-in function with second message
	MessageCheckIn(test1, oInMessages, qMessages)
	fmt.Println("OutgoingMessages queue size should be the same. Actual size:", len(qMessages.OutgoingMessages))

	// Test message check-in function with third message
	fmt.Println()
	MessageCheckIn(test2, oInMessages, qMessages)
	fmt.Println("OutgoingMessages queue size should be incremented. Reason: Echo - Actual size:", len(qMessages.OutgoingMessages))

	// Test message check-in function with third message again: should already exist
	fmt.Println()
	MessageCheckIn(test2, oInMessages, qMessages)
	fmt.Println("OutgoingMessages queue size should be the same. Actual size:", len(qMessages.OutgoingMessages))

	// Test message check-in function with forth message: receiver ID should not match
	fmt.Println()
	MessageCheckIn(test3, oInMessages, qMessages)
	fmt.Println("OutgoingMessages queue size should be incremented. Reason: Echo - Actual size:", len(qMessages.OutgoingMessages))
	fmt.Println()

	fmt.Println("------------------------------------------------------------------------------------------------------------------------\nDisplay deserialized message queue - Before resolveMessage:\n ")
	fmt.Println("Deserialized Queue:", qMessages.DeserialMessages)

	// Display message ID list
	fmt.Println()
	fmt.Println("Displayed incoming MessageID list - Before resolveMessage:\n ")
	fmt.Printf("MessageID list: %s\n", oInMessages)

	// Remove serviced message from queue
	responseID := msg.MessageID
	oInMessages.resolveMessage(responseID, qMessages)

	// Display message ID list
	fmt.Println()
	fmt.Println("Displayed incoming MessageID list - After One resolveMessage:\n ")
	fmt.Printf("MessageID list: %s\n", oInMessages)

	// Remove serviced messages from queue
	responseID2 := msg2.MessageID
	oInMessages.resolveMessage(responseID2, qMessages)
	fmt.Println()
	fmt.Println("Display deserialized message queue - After resolveMessage:\n ")
	fmt.Println("Deserialized Queue:", qMessages.DeserialMessages)

	// Display message ID list
	fmt.Println()
	fmt.Println("Displayed incoming MessageID list:\n ")
	fmt.Printf("MessageID list: %s\n", oInMessages)

	fmt.Println()
	fmt.Println("Display effects of dequeue function on all three types of message queues:\n ")
	fmt.Println("Length of Deserialized Queue:", len(qMessages.DeserialMessages))
	fmt.Println("Length of Incoming Queue:", len(qMessages.IncomingMessages))
	fmt.Println("Length of Outgoing Queue:", len(qMessages.OutgoingMessages))
	fmt.Println()

	// Dequeue from the deserialized messages queue
	msgFromDeserial, err := qMessages.Dequeue("deserial")
	if err != nil {
		fmt.Println("Error dequeuing from deserial queue:", err)
	} else {
		fmt.Println("Dequeued from deserial queue:", msgFromDeserial)
	}

	// Dequeue from the incoming messages queue
	_, err = qMessages.Dequeue("incoming")
	if err != nil {
		fmt.Println("Error dequeuing from incoming queue:", err)
	} else {
		fmt.Println("Dequeued from incoming queue: Some byte message")
	}

	// Dequeue from the outgoing messages queue
	_, err = qMessages.Dequeue("outgoing")
	if err != nil {
		fmt.Println("Error dequeuing from outgoing queue:", err)
	} else {
		fmt.Println("Dequeued from outgoing queue: Some byte message")
	}

	fmt.Println()
	fmt.Println("Length of Deserialized Queue:", len(qMessages.DeserialMessages))
	fmt.Println("Length of Incoming Queue:", len(qMessages.IncomingMessages))
	fmt.Println("Length of Outgoing Queue:", len(qMessages.OutgoingMessages))

	fmt.Print("\n\nTesting Encryption and decryption of serialized data:\n")
	messageID := oOutMessages.generateMessageID()
	msgtest := Message{
		MessageID:   messageID,
		SenderID:    "Pi-1",
		ReceiverID:  "Pi-2",
		Domain:      "lighting",
		OperationID: "2",
		Data:        data,
	}

	jsonData, err := json.Marshal(msgtest)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	key := []byte("1234567890123456")

	encrypt, err := hashing.EncryptGCM(string(jsonData), string(key))
	if err != nil {
		fmt.Println("Error encrypting JSON:", err)
		return
	}
	fmt.Println()

	fmt.Print("\nBefore encrypted DATA TEST:\n", string(jsonData))
	fmt.Println()
	fmt.Print("\nAfter encrypted DATA TEST:\n", string(encrypt))
	fmt.Println()
	fmt.Print("\nConvert encrypted to []byte:\n", []byte(encrypt))

	decrypt, err := hashing.DecryptGCM(encrypt, string(key))
	if err != nil {
		fmt.Print("Error decrypting JSON:", err)
		return
	}
	fmt.Println()
	fmt.Print("\nAfter decrypted DATA TEST:\n", string(decrypt))
	fmt.Println("\n ")
	idecrypt := []byte(decrypt)
	MessageCheckIn(idecrypt, oInMessages, qMessages)
}
*/
