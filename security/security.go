package security

import (
	"fmt"
)

// DoorLock represents a door lock system.
type DoorLock struct {
	Name   string
	Locked bool
}

// NewDoorLock creates a new DoorLock instance with the given name.
func NewDoorLock(name string) *DoorLock {
	return &DoorLock{
		Name:   name,
		Locked: true, // Default to locked for security purposes.
	}
}

// Lock will lock the door.
func (dl *DoorLock) Lock() {
	dl.Locked = true
	fmt.Printf("%s is now locked.\n", dl.Name)
}

// Unlock will unlock the door.
func (dl *DoorLock) Unlock() {
	dl.Locked = false
	fmt.Printf("%s is now unlocked.\n", dl.Name)
}

// Camera represents a surveillance camera.
type Camera struct {
	Name   string
	Active bool
}

// NewCamera creates a new Camera instance with the given name.
func NewCamera(name string) *Camera {
	return &Camera{
		Name:   name,
		Active: false, // Default to not active.
	}
}

// Activate starts the camera.
func (c *Camera) Activate() {
	c.Active = true
	fmt.Printf("%s is now active.\n", c.Name)
}

// Deactivate stops the camera.
func (c *Camera) Deactivate() {
	c.Active = false
	fmt.Printf("%s is now inactive.\n", c.Name)
}

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
func (a *Alarm) Arm() bool {
	a.Armed = true
	fmt.Printf("%s is armed.\n", a.Name)
	return a.Armed == true

}

// Disarm disarms the alarm.
func (a *Alarm) Disarm() bool {
	a.Armed = false
	fmt.Printf("%s is disarmed.\n", a.Name)
	return a.Armed == false
}

// Trigger activates the alarm if it's armed.
func (a *Alarm) Trigger() {
	if a.Armed {
		a.Sounded = true
		fmt.Printf("%s alarm is triggered!\n", a.Name)
	}
}

// TrustCenter manages user authorization to interact with the security system.
type TrustCenter struct {
	AuthorizedUsers map[string]bool
}

// NewTrustCenter creates a new TrustCenter instance.
func NewTrustCenter() *TrustCenter {
	return &TrustCenter{
		AuthorizedUsers: make(map[string]bool),
	}
}

// AuthorizeUser adds a user to the list of authorized users.
func (tc *TrustCenter) AuthorizeUser(userName string) {
	tc.AuthorizedUsers[userName] = true
	fmt.Printf("%s is authorized.\n", userName)
}

// IsUserAuthorized checks if a user is authorized.
func (tc *TrustCenter) IsUserAuthorized(userName string) bool {
	authorized, exists := tc.AuthorizedUsers[userName]
	return exists && authorized
}

// AttemptJoin simulates a user attempting to join the network.
func (tc *TrustCenter) AttemptJoin(userName string) string {
	if tc.IsUserAuthorized(userName) {
		return fmt.Sprintf("%s has successfully joined the network.\n", userName)
	}
	return fmt.Sprintf("%s is not authorized to join the network.\n", userName)
}
