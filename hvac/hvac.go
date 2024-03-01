package main

import (
	"fmt"
	"time"
)

// HVAC represents an HVAC system with temperature control, fan speed, and mode.
type HVAC struct {
	Name        string
	Temperature int    // Desired temperature in Fahrenheit
	FanSpeed    int    // Fan speed (0-100%)
	Mode        string // HVAC mode (e.g., "Cool", "Heat", "Fan", "Off")
}

// NewHVAC creates a new HVAC instance with the given name and initial settings.
func NewHVAC(name string) *HVAC {
	return &HVAC{
		Name:        name,
		Temperature: 77,    // Initial temperature setting in Fahrenheit
		FanSpeed:    50,    // Initial fan speed setting (50%)
		Mode:        "Off", // Initial mode is Off
	}
}

<<<<<<< HEAD
// SetTemperature sets the desired temperature for the HVAC system.
func (h *HVAC) SetTemperature(temperature int) bool {
	h.Temperature = temperature
	fmt.Printf("%s temperature is set to %d°C\n", h.Name, h.Temperature)
	return true
=======
// SetTemperature sets the desired temperature for the HVAC system in Fahrenheit.
func (h *HVAC) SetTemperature(temperature int) {
	if temperature < 60 {
		temperature = 60
	} else if temperature > 90 {
		temperature = 90
	}
	h.Temperature = temperature
	fmt.Printf("%s temperature is set to %d°F\n", h.Name, h.Temperature)
>>>>>>> ZigBee_Network_Group
}

// SetFanSpeed sets the fan speed for the HVAC system.
func (h *HVAC) SetFanSpeed(speed int) bool {
	if speed < 0 {
		speed = 0
	} else if speed > 100 {
		speed = 100
	}
	h.FanSpeed = speed
	fmt.Printf("%s fan speed is set to %d%%\n", h.Name, h.FanSpeed)
	return true
}

// SetMode sets the mode (e.g., "Cool", "Heat", "Fan", "Off") for the HVAC system.
func (h *HVAC) SetMode(mode string) bool {
	h.Mode = mode
	fmt.Printf("%s mode is set to %s\n", h.Name, h.Mode)
	return true
}

// ScheduledEvent represents a scheduled event for the HVAC system.
type ScheduledEvent struct {
	Name        string
	Time        time.Time
	Temperature int // Temperature in Fahrenheit
	FanSpeed    int
	Mode        string
}

// Scheduler manages scheduled events for an HVAC system.
type Scheduler struct {
	Events []ScheduledEvent
}

// NewScheduler creates a new Scheduler instance.
func NewScheduler() *Scheduler {
	return &Scheduler{}
}

// AddEvent adds a new scheduled event to the Scheduler.
func (s *Scheduler) AddEvent(event ScheduledEvent) {
	s.Events = append(s.Events, event)
}

// CheckAndApplyEvents checks and applies any due scheduled events.
func (s *Scheduler) CheckAndApplyEvents(hvac *HVAC) {
	now := time.Now()
	var remainingEvents []ScheduledEvent
	for _, event := range s.Events {
		if now.After(event.Time) && now.Sub(event.Time) < 1*time.Minute {
			hvac.SetMode(event.Mode)
			hvac.SetTemperature(event.Temperature)
			hvac.SetFanSpeed(event.FanSpeed)
			fmt.Printf("Applied scheduled event: %s - %+v\n", event.Name, event)
		} else {
			remainingEvents = append(remainingEvents, event)
		}
	}
	s.Events = remainingEvents
}

func main() {
	// Create a new HVAC system
	livingRoomHVAC := NewHVAC("Living Room HVAC")

	// Create a new scheduler
	scheduler := NewScheduler()

	// Schedule some named events
	morningWarmUp := ScheduledEvent{
		Name:        "Morning Warm-Up",
		Time:        time.Now().Add(30 * time.Second), // Example: 30 seconds from now
		Temperature: 68,                               // Temperature in Fahrenheit
		FanSpeed:    75,
		Mode:        "Heat",
	}
	scheduler.AddEvent(morningWarmUp)

	eveningCoolDown := ScheduledEvent{
		Name:        "Evening Cool-Down",
		Time:        time.Now().Add(60 * time.Second), // Example: 60 seconds from now
		Temperature: 72,                               // Temperature in Fahrenheit
		FanSpeed:    50,
		Mode:        "Cool",
	}
	scheduler.AddEvent(eveningCoolDown)

	// Periodically check and apply scheduled events
	for range time.Tick(30 * time.Second) {
		scheduler.CheckAndApplyEvents(livingRoomHVAC)
	}
}
