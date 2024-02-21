package appliances

import (
	"fmt"
)

// Appliance represents a simple appliance with an on/off state.
type Appliance struct {
	Name  string
	State bool
}

// NewAppliance creates a new Appliance instance with the given name and initial state.
func NewAppliance(name string, initialState bool) *Appliance {
	return &Appliance{
		Name:  name,
		State: initialState,
	}
}

// TurnOn turns the appliance on.
func (a *Appliance) TurnOn() {
	a.State = true
	fmt.Printf("%s is now turned ON\n", a.Name)
}

// TurnOff turns the appliance off.
func (a *Appliance) TurnOff() {
	a.State = false
	fmt.Printf("%s is now turned OFF\n", a.Name)
}

// ToggleState used to turn "on" or "off"
func (a *Appliance) ToggleState() {
	// Inverts the State
	a.State = !a.State
	state := "OFF"
	if a.State {
		state = "ON"
	}
	fmt.Printf("%s is now toggled to %s\n", a.Name, state)
}

// CheckState used to get the latest state of the appliance
func (a *Appliance) CheckState() string {
	if a.State {
		return "ON"
	}
	return "OFF"
}

/*
func main() {
	// Create a new light switch appliance
	lightSwitch := NewAppliance("Light Switch", false)
	fan := NewAppliance("Fan", false)

	// Use the light appliance
	fmt.Printf("%s is currently %s\n", lightSwitch.Name, lightSwitch.CheckState())
	//lightSwitch.TurnOn()
	lightSwitch.ToggleState()
	//Testing to see if retreiving current state of the appliance works
	fmt.Printf("%s is currently %s\n", lightSwitch.Name, lightSwitch.CheckState())
	lightSwitch.ToggleState()
	fmt.Printf("%s is currently %s\n", lightSwitch.Name, lightSwitch.CheckState())

	// Use the fan appliance
	//fan.TurnOn()
	// Use the light appliance
	fmt.Printf("%s is currently %s\n", fan.Name, fan.CheckState())
	//fan.TurnOn()
	fan.ToggleState()
	//Testing to see if the current state of the appliance
	fmt.Printf("%s is currently %s\n", fan.Name, fan.CheckState())
	fan.ToggleState()
	fmt.Printf("%s is currently %s\n", fan.Name, fan.CheckState())

	//Is used to check multiple states at once
	appliances := []*Appliance{lightSwitch, fan}
	for _, appliance := range appliances {
		appliance.CheckState()
	}
}
*/
