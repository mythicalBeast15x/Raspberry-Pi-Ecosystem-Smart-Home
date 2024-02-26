package messagingTests

import (
	"CMPSC488SP24SecThursday/messaging"
	"testing"
)

func TestMessageCheckIn(t *testing.T) {
	// Define test cases
	tests := []struct {
		name       string
		serialMsg  []byte
		oMessages  *messaging.OpenMessages
		qMessages  *messaging.MessageQueue
		wantResult bool
	}{
		{
			name:       "Duplicate Message",
			serialMsg:  []byte(`{"messageID":"123","senderID":"Pi-1","receiverID":"Pi-2","domain":"Lighting","operationID":"2","Data":{"ApplianceID":"lamp-1","Brightness":70}}`),
			oMessages:  &messaging.OpenMessages{},
			qMessages:  &messaging.MessageQueue{},
			wantResult: false,
		},
		{
			name:       "Valid Message: Intended for All",
			serialMsg:  []byte(`{"messageID":"123","senderID":"Pi-1","receiverID":"all","domain":"Lighting","operationID":"2","Data":{"ApplianceID":"lamp-1","Brightness":70}}`),
			oMessages:  &messaging.OpenMessages{},
			qMessages:  &messaging.MessageQueue{},
			wantResult: true,
		},
		{
			name:       "Invalid Message: Not Intended for self",
			serialMsg:  []byte(`{"messageID":"123","senderID":"Pi-1","receiverID":"Pi-1","domain":"Lighting","operationID":"2","Data":{"ApplianceID":"lamp-1","Brightness":70}}`),
			oMessages:  &messaging.OpenMessages{},
			qMessages:  &messaging.MessageQueue{},
			wantResult: false,
		},
		{
			name:       "Valid Message: Intended for self",
			serialMsg:  []byte(`{"messageID":"123","senderID":"Pi-1","receiverID":"Pi-2","domain":"Lighting","operationID":"2","Data":{"ApplianceID":"lamp-1","Brightness":70}}`),
			oMessages:  &messaging.OpenMessages{},
			qMessages:  &messaging.MessageQueue{},
			wantResult: true,
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "Duplicate Message" {
				// Manually add message ID for duplicate message test case
				tt.oMessages.AddMessageID("123")
			}
			gotResult := messaging.MessageCheckIn(tt.serialMsg, tt.oMessages, tt.qMessages)
			if gotResult != tt.wantResult {
				t.Errorf("MessageCheckIn() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
