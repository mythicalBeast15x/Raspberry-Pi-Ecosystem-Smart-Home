package securityTest

import (
	"CMPSC488SP24SecThursday/security"
	"testing"
)

func TestUserJoiningProcess(t *testing.T) {
	tc := security.NewTrustCenter()

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

func TestMotionDetectionAndAlarmTrigger(t *testing.T) {
	// Initialize the TrustCenter and authorize a user.
	tc := security.NewTrustCenter()
	tc.AuthorizeUser("AuthorizedUser")

	// Create an alarm and a motion sensor with a callback to trigger the alarm.
	alarmed := false // Use a flag to capture if the alarm was triggered.
	securityAlarm := security.NewAlarm("Security Alarm")
	motionSensor := security.NewMotionSensor("Motion Sensor", func(sensorName string) {
		if securityAlarm.Armed {
			securityAlarm.Trigger()
			alarmed = true // Set the flag if the alarm is triggered.
		}
	})

	// Arm the security system.
	securityAlarm.Arm()

	// Simulate motion detection.
	motionSensor.DetectMotion()

	// Check if the alarm was triggered.
	if !alarmed {
		t.Error("Expected the alarm to be triggered upon motion detection when armed.")
	}

	// Reset for disarmed scenario.
	alarmed = false
	securityAlarm.Disarm()

	// Simulate motion detection again with the system disarmed.
	motionSensor.DetectMotion()

	// The alarm should not be triggered this time.
	if alarmed {
		t.Error("Expected the alarm not to be triggered upon motion detection when disarmed.")
	}
}
