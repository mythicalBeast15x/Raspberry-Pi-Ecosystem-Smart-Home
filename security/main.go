package main

import (
	"fmt"
)

// MotionSensor represents a motion sensor component.
type MotionSensor struct {
	Name             string
	MotionDetected   bool
	OnMotionDetected func(string) // Callback function to execute when motion is detected
}

// NewMotionSensor creates a new MotionSensor instance with the given name.
func NewMotionSensor(name string, onMotionDetected func(string)) *MotionSensor {
	return &MotionSensor{
		Name:             name,
		OnMotionDetected: onMotionDetected,
	}
}

// DetectMotion simulates detecting motion.
func (m *MotionSensor) DetectMotion() {
	m.MotionDetected = true
	fmt.Printf("%s detected motion!\n", m.Name)
	if m.OnMotionDetected != nil {
		m.OnMotionDetected(m.Name)
	}
}

// Callback function that could be used to trigger an alarm when motion is detected.
func alarmTriggerCallback(sensorName string) {
	fmt.Printf("Alarm triggered by %s.\n", sensorName)
	// going to call the alarm's Trigger method here
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

	// Authorizing users and attempting to join the network.
	tc.AuthorizeUser("AuthorizedUser")
	fmt.Println(tc.AttemptJoin("AuthorizedUser"))
	fmt.Println(tc.AttemptJoin("UnauthorizedUser"))

	// Create an alarm instance and arm it.
	securityAlarm := NewAlarm("Security Alarm")
	securityAlarm.Arm()

	// Creating a motion sensor with a callback to automatically trigger the alarm on motion detection.
	motionSensor := NewMotionSensor("Motion Sensor", func(sensorName string) {
		if securityAlarm.Armed {
			securityAlarm.Trigger()
		}
	})

	// Simulate motion detection,should automatically trigger the alarm if it's armed.
	motionSensor.DetectMotion()

	// Trigger the alarm to demonstrate manual activation, regardless of motion detection.
	// This call might represent a direct interaction, like a panic button being pressed.
	fmt.Println("Manually triggering the alarm to test behavior when armed.")
	securityAlarm.Trigger()

	// Disarm the security system to prevent the alarm from sounding on further motion detection.
	securityAlarm.Disarm()

	// Attempt to manually trigger the alarm again to illustrate it doesn't sound when disarmed.
	// This might be useful for testing the disarm functionality.
	fmt.Println("Attempting to manually trigger the alarm after disarming to test behavior.")
	securityAlarm.Trigger()
}
