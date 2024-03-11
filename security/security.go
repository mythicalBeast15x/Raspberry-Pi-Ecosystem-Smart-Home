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
