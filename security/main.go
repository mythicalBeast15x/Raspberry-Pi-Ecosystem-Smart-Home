package main

import (
	"fmt"
)

// MotionSensor represents a motion sensor component.
type MotionSensor struct {
	Name           string
	MotionDetected bool
}

// NewMotionSensor creates a new MotionSensor instance with the given name.
func NewMotionSensor(name string) *MotionSensor {
	return &MotionSensor{
		Name: name,
	}
}

// DetectMotion simulates detecting motion.
func (m *MotionSensor) DetectMotion() {
	m.MotionDetected = true
	fmt.Printf("%s detected motion!\n", m.Name)
}

// Alarm represents a simple alarm component.
type Alarm struct {
	Name    string
	Armed   bool
	Sounded bool
}

// NewAlarm creates a new Alarm instance with the given name.
func NewAlarm(name string) *Alarm {
	return &Alarm{
		Name:  name,
		Armed: false,
	}
}

// Arm sets the alarm to the armed state.
func (a *Alarm) Arm() {
	a.Armed = true
	fmt.Printf("%s is armed and ready.\n", a.Name)
}

// Disarm disarms the alarm.
func (a *Alarm) Disarm() {
	a.Armed = false
	fmt.Printf("%s is disarmed.\n", a.Name)
}

// Trigger activates the alarm if it's armed.
func (a *Alarm) Trigger() {
	if a.Armed {
		a.Sounded = true
		fmt.Printf("%s is triggered! Alarm is sounding.\n", a.Name)
	}
}

// TrustCenter now keeps track of users allowed to join and communicate within the network.
type TrustCenter struct {
	// authorizedUsers maps user names to their authorization status.
	authorizedUsers map[string]bool
}

// NewTrustCenter initializes and returns a new instance of TrustCenter.
func NewTrustCenter() *TrustCenter {
	return &TrustCenter{
		authorizedUsers: make(map[string]bool), // Initialize the map to track authorized users.
	}
}

// AuthorizeUser marks a user as authorized within the network.
func (tc *TrustCenter) AuthorizeUser(userName string) {
	tc.authorizedUsers[userName] = true
}

// IsUserAuthorized checks if a given user is authorized to operate within the network.
func (tc *TrustCenter) IsUserAuthorized(userName string) bool {
	authorized, exists := tc.authorizedUsers[userName]
	return exists && authorized
}

// AttemptJoin simulates a user attempting to join the network.
func (tc *TrustCenter) AttemptJoin(userName string) string {
	if tc.IsUserAuthorized(userName) {
		return fmt.Sprintf("%s has successfully joined the network.", userName)
	}
	return fmt.Sprintf("%s failed to join the network. Unauthorized user.", userName)
}

func main() {
	tc := NewTrustCenter()

	// Example: Authorizing users and attempting to join the network.
	tc.AuthorizeUser("AuthorizedUser")

	// Simulate an authorized user attempting to join.
	fmt.Println(tc.AttemptJoin("AuthorizedUser"))

	// Simulate an unauthorized user attempting to join.
	fmt.Println(tc.AttemptJoin("UnauthorizedUser"))

	// Create a motion sensor and an alarm
	motionSensor := NewMotionSensor("Motion Sensor")
	securityAlarm := NewAlarm("Security Alarm")

	// Arm the security system
	securityAlarm.Arm()

	// Simulate motion detection
	motionSensor.DetectMotion()

	// Check if the alarm is triggered
	securityAlarm.Trigger()

	// Disarm the security system
	securityAlarm.Disarm()

	// Simulate motion detection again
	motionSensor.DetectMotion()

	// Check if the alarm is triggered (it should not be triggered this time)
	securityAlarm.Trigger()
}
