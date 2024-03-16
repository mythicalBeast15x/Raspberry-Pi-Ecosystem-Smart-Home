package blockchainTests

import (
	"CMPSC488SP24SecThursday/dal"
	"CMPSC488SP24SecThursday/messaging"
	"fmt"
	"testing"
)

//for reference
//// Message represents the structure of a message object.
//type Message struct {
//	MessageID   string                 `json:"messageID"`
//	SenderID    string                 `json:"senderID"`
//	ReceiverID  string                 `json:"receiverID"`
//	Domain      string                 `json:"domain"`
//	OperationID string                 `json:"operationID"`
//	Data        map[string]interface{} `json:"Data"`
//}

// testing the block creations
func TestLights(t *testing.T) {
	//test data
	Data := map[string]interface{}{
		"Structure": map[string]interface{}{
			"DeviceName": "Light",
			"Status":     "false",
		},
		"Status": "false",
	}

	TestMessage := messaging.NewMessage("1", "2", "Lighting", "0", Data, &messaging.OpenMessages{}, &messaging.MessageQueue{})
	fmt.Print(dal.MessageDisperser(TestMessage))
	//Test for Validating a correct block
}

//func TestHVAC(t *testing.T) {
//	testHVAC := hvac.NewHVAC("fan")
//	testResult := dal.HVAC(9, testHVAC, 100, 100, "cool")
//	if testResult != true {
//		t.Errorf("Expected: %t, Got: %t", true, testResult)
//	}
//
////}
//
//func TestSecurity(t *testing.T) {
//	testSecurity := security.NewAlarm("Fire")
//	testResult := dal.SecuritySystem(6, testSecurity)
//	if testResult != true {
//		t.Errorf("Expected: %t, Got: %t", true, testResult)
//	}
//}
