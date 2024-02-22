
package main

package energy


import (
	"fmt"
	"sync" // Sync Package: https://pkg.go.dev/sync
)

// SolarPanel represents a solar panel as a power source.
type SolarPanel struct {
	Name        string
	PowerOutput int // Power output in watts
}

// NewSolarPanel creates a new SolarPanel instance with the given name and power output.
func NewSolarPanel(name string, powerOutput int) *SolarPanel {
	return &SolarPanel{
		Name:        name,
		PowerOutput: powerOutput,
	}
}

// Battery represents an energy storage battery.
type Battery struct {
	Name     string
	Capacity int // Battery capacity in watt-hours (Wh)
	Charge   int // Current charge level in watt-hours (Wh)
}

// NewBattery creates a new Battery instance with the given name and capacity.
func NewBattery(name string, capacity int) *Battery {
	return &Battery{
		Name:     name,
		Capacity: capacity,
		Charge:   0, // Initialize with 0 charge
	}
}

// Appliance represents an appliance with energy consumption.
type Appliance struct {
	Name        string
	PowerRating int  // Power rating of the appliance in watts
	IsOn        bool // Whether the appliance is turned on
}

// NewAppliance creates a new Appliance instance with the given name and power rating.
func NewAppliance(name string, powerRating int) *Appliance {
	return &Appliance{
		Name:        name,
		PowerRating: powerRating,
		IsOn:        false,
	}
}

// TurnOn turns the appliance on.
func (a *Appliance) TurnOn() {
	a.IsOn = true
	fmt.Printf("%s is turned ON\n", a.Name)
}

// TurnOff turns the appliance off.
func (a *Appliance) TurnOff() {
	a.IsOn = false
	fmt.Printf("%s is turned OFF\n", a.Name)
}

/**********************************************************************************
	My Intention with the following code additions are as follows:

	The given code supplied me with solar panels, batteries, and appliance-type objects.

	The solar panel feeds electricity to the battery and the appliances contain a name,
	power rating of how much power it uses and a boolean value of its on/off state.

	NOTE: The energy security.go file is not the only file to contain functions to turn on/off
		  appliances. This leaves me to believe that the energy file is mainly a source of
		  documenting the usage of power by each appliance if they are turned on/off instead
		  of physically turning them on/off.

	The following code adds two more types:
	- ZigbeeRequest: Represents an incoming request.
		- Currently not implementing the Go-Zigbee package to represent a real request.

	- ApplianceController: Holds a list of all registered appliances.
		- Also holds a mutex so that multiple actions cannot be serviced at once.

	This object is intended to help check to see if a request holds a valid appliance.
	If the appliance exists on record, then the state change request can be compared
	to the current state of the appliance.

	This is done through the HandelZigbeeRequest function which takes in the request
	and utilizes the ApplianceController to compare the device IDs and State change
	requests. Depending on the request, the handler will either change the state
	of the appliance if the request differs, relay that the appliance's state
	already matches the requested state, or relay that the requested device is not
	currently registered to the controller.

	For now, since it would be problematic for many requests to be submitted
	simultaneously, I decided to use the go-sync package to add in a mutex lock.
	This is to protect the list of appliances from race conditions when modifying
	the data. This is seen both during the adding of an appliance to the controller,
	and during the request handling which may alter the state of the appliance.

	I also used the 'defer' keyword to make sure that the mutex would be scheduled to
	unlock at the end of the function regardless of the function exiting in success
	or in an error.

	sync.Mutex Example: https://go.dev/tour/concurrency/9
*********************************************************************************/

// ZigbeeRequest represents an incoming request for an appliance's state to be changed.
// NOTE: This currently does not implement the go-zigbee package!
type ZigbeeRequest struct {
	ApplianceName  string
	RequestedState bool
}

// ApplianceController manages the state of appliances and handles Zigbee requests.
type ApplianceController struct {
	appliances []*Appliance
	mutex      sync.Mutex
}

// NewApplianceController create a new ApplianceController instance.
/*
Purpose: Creates a new ApplianceController.
Creates a list that will be filled with all appliances that the user wishes to
add to the network.
*/
func NewApplianceController() *ApplianceController {
	return &ApplianceController{
		appliances: make([]*Appliance, 0),
	}
}

// AddAppliance adds a new appliance to the controller's list.
/*
Purpose: Adds an appliance to the ApplianceController.
Since the ApplianceController is a shared resource, it is locked to ensure it
handles one request at a time. An unlock command is deferred to ensure the ApplianceController
is free upon exiting the function. The function then attempts to add the new appliance to the
ApplianceController list.

Data : ApplianceController - list holding registered elements of type=appliances.
*/
func (ac *ApplianceController) AddAppliance(appliance *Appliance) {
	// Lock the mutex since modification is possible.
	ac.mutex.Lock()
	defer ac.mutex.Unlock()
	ac.appliances = append(ac.appliances, appliance)
}

// HandelZigbeeRequest handles zigbee requests to change the state of an appliance.
/*
Purpose: To handle incoming appliance related requests.
The ApplianceController's mutex is locked since the data may be altered with a deferred
unlock command upon exiting the function.

The function checks first if the requested appliance exists within the ApplianceController's
list of registered appliances. If it is not, then it notifies via a printf statement and exits.

If the appliance exists, then the function checks if the desired state matches the current
state of the appliance. If so, then it notifies via a printf statement and exits.

If the desired and current state of the appliance differ, then the state is changed and the
user is notified of the change.

Data : ApplianceController - list holding registered elements of type=appliances.
Data : ZigbeeRequest - Holds the input data of the incoming request with...
					   Name of type=string
					   State of type=bool
*/
func (ac *ApplianceController) HandelZigbeeRequest(request ZigbeeRequest) {
	// Lock the mutex since modification is possible.
	ac.mutex.Lock()
	defer ac.mutex.Unlock()

	// Find the requested appliance in the list.
	for _, appliance := range ac.appliances {
		if appliance.Name == request.ApplianceName {
			// If the requested Appliance is found, check and update the
			// state if it differs from the appliance's current state.
			// Otherwise, print that the request is already met.
			if appliance.IsOn != request.RequestedState {
				appliance.IsOn = request.RequestedState
				fmt.Printf("%s is turned %s\n", appliance.Name, map[bool]string{true: "ON", false: "OFF"}[appliance.IsOn])
			} else {
				fmt.Printf("%s is already %s\n", appliance.Name, map[bool]string{true: "ON", false: "OFF"}[appliance.IsOn])
			}
			return
		}
	}
	// If the requested appliance is not found, print the following:
	fmt.Printf("Appliance '%s' not found\n", request.ApplianceName)
}

func main() {
	// Create a solar panel, battery, and appliances
	solarPanel := NewSolarPanel("Solar Panel", 500)         // 500 watts of power output
	houseBattery := NewBattery("House Battery", 2000)       // 2000 watt-hours capacity
	fridge := NewAppliance("Fridge", 200)                   // 200 watts
	airConditioner := NewAppliance("Air Conditioner", 1500) // 1500 watts

	// Simulate powering the appliances with solar energy
	solarEnergy := solarPanel.PowerOutput
	houseBattery.Charge += solarEnergy

	// Turn on appliances
	fridge.TurnOn()
	airConditioner.TurnOn()

	// Simulate appliance power consumption
	if fridge.IsOn {
		houseBattery.Charge -= fridge.PowerRating
	}
	if airConditioner.IsOn {
		houseBattery.Charge -= airConditioner.PowerRating
	}

	// Check battery charge level
	fmt.Printf("House Battery Charge Level: %d Wh\n", houseBattery.Charge)

	/*********************************************************************************
		Beginning of a demonstration of the added request handling functions:

		The program creates a new instance of the ApplianceController. It then creates
		and adds both a light and fan appliances. Both begin in the off state.

		Three test cases are then showcased:
		- Appliance initially off and asked to turn on.
		- Appliance already turned on and asked to be turned on.
		- An unregistered appliance asked to be turned on.

		NOTE: Further additions to accommodate the power consumption of each appliance will be added.
	*********************************************************************************/

	fmt.Printf("\nDemo Start:\n")
	// New ApplianceController.
	controller := NewApplianceController()

	// Light and Fan appliances created.
	light := NewAppliance("Light", 60)
	fan := NewAppliance("Fan", 100)

	// Light and Fan appliances registered to the controller
	controller.AddAppliance(light)
	controller.AddAppliance(fan)

	// Requesting to turn on light - initially off
	zigRequest1 := ZigbeeRequest{ApplianceName: "Light", RequestedState: true}
	controller.HandelZigbeeRequest(zigRequest1)

	// Requesting to turn on fan - initially off
	zigRequest2 := ZigbeeRequest{ApplianceName: "Fan", RequestedState: true}
	controller.HandelZigbeeRequest(zigRequest2)

	// Requesting to turn on light - currently on
	zigRequest3 := ZigbeeRequest{ApplianceName: "Light", RequestedState: true}
	controller.HandelZigbeeRequest(zigRequest3)

	// Requesting to turn on coffee maker - unregistered
	zigRequest4 := ZigbeeRequest{ApplianceName: "Coffee_Maker", RequestedState: true}
	controller.HandelZigbeeRequest(zigRequest4)

	// Requesting to turn off fan - currently on
	zigRequest5 := ZigbeeRequest{ApplianceName: "Fan", RequestedState: false}
	controller.HandelZigbeeRequest(zigRequest5)

	// Requesting to turn off light - currently on
	zigRequest6 := ZigbeeRequest{ApplianceName: "Light", RequestedState: false}
	controller.HandelZigbeeRequest(zigRequest6)

}
