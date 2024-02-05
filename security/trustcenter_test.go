package main

import (
	"testing"
)

func TestUserJoiningProcess(t *testing.T) {
	tc := NewTrustCenter()

	// Authorize one user
	authorizedUser := "AuthorizedUser"
	tc.AuthorizeUser(authorizedUser) // Updated to AuthorizeUser

	// Simulate join attempt by an authorized user
	joinMessage := tc.AttemptJoin(authorizedUser)
	expectedMessage := authorizedUser + " has successfully joined the network."
	if joinMessage != expectedMessage {
		t.Errorf("Expected '%s', got '%s'", expectedMessage, joinMessage)
	}

	// Simulate join attempt by an unauthorized user
	unauthorizedUser := "UnauthorizedUser"
	joinMessage = tc.AttemptJoin(unauthorizedUser)
	expectedMessage = unauthorizedUser + " failed to join the network. Unauthorized user." // Updated message to reflect user context
	if joinMessage != expectedMessage {
		t.Errorf("Expected '%s', got '%s'", expectedMessage, joinMessage)
	}
}
