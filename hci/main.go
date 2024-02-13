package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Simulated smart home devices
var lights bool
var thermostat int
var appliances map[string]bool
var security bool

func main() {
	lights = false
	thermostat = 70
	appliances = make(map[string]bool)
	security = true

	reader := bufio.NewReader(os.Stdin)

menuLoop:
	for {
		fmt.Println("\nSmart Home Control System")
		fmt.Println("1. Toggle Lights")
		fmt.Println("2. Adjust Thermostat")
		fmt.Println("3. Add Appliance")
		fmt.Println("4. Check Appliance Status")
		fmt.Println("5. Turn On/Off Appliance")
		fmt.Println("6. Trigger Heating/Cooling")
		fmt.Println("7. Check Energy Usage")
		fmt.Println("8. Check Security Status")
		fmt.Println("9. Exit")
		fmt.Print("Select an option: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			ToggleLights()

		case "2":
			AdjustThermostat(reader)

		case "3":
			AddAppliance(reader)

		case "4":
			CheckApplianceStatus(reader)

		case "5":
			TurnOnOffAppliance(reader)

		case "6":
			TriggerHeatingCooling(reader)

		case "7":
			CheckEnergyUsage()

		case "8":
			CheckSecurityStatus()

		case "9":
			fmt.Println("Goodbye!")
			break menuLoop

		default:
			fmt.Println("Invalid option. Please select a valid option.")
		}
	}
}

func ToggleLights() {
	lights = !lights
	if lights {
		fmt.Println("Lights turned on.")
	} else {
		fmt.Println("Lights turned off.")
	}
}

func AdjustThermostat(reader *bufio.Reader) {
	fmt.Print("Enter new temperature: ")
	tempStr, _ := reader.ReadString('\n')
	tempStr = strings.TrimSpace(tempStr)
	temp, err := strconv.Atoi(tempStr)
	if err != nil {
		fmt.Println("Invalid input for thermostat.")
	} else {
		thermostat = temp
		fmt.Printf("Thermostat set to %d°F\n", thermostat)
	}
}

func AddAppliance(reader *bufio.Reader) {
	fmt.Print("Enter name of appliance: ")
	applianceName, _ := reader.ReadString('\n')
	applianceName = strings.TrimSpace(applianceName)
	if _, exists := appliances[applianceName]; exists {
		fmt.Println("Appliance already exists.")
	} else {
		appliances[applianceName] = false
		fmt.Printf("Appliance '%s' added.\n", applianceName)
	}
}

func CheckApplianceStatus(reader *bufio.Reader) {
	fmt.Print("Enter name of appliance: ")
	applianceName, _ := reader.ReadString('\n')
	applianceName = strings.TrimSpace(applianceName)
	status, exists := appliances[applianceName]
	if exists {
		if status {
			fmt.Printf("Appliance '%s' is On.\n", applianceName)
		} else {
			fmt.Printf("Appliance '%s' is Off.\n", applianceName)
		}
	} else {
		fmt.Println("Appliance not found.")
	}
}

func TurnOnOffAppliance(reader *bufio.Reader) {
	fmt.Print("Enter name of appliance: ")
	applianceName, _ := reader.ReadString('\n')
	applianceName = strings.TrimSpace(applianceName)
	status, exists := appliances[applianceName]
	if exists {
		appliances[applianceName] = !status
		if !status {
			fmt.Printf("Appliance '%s' turned on.\n", applianceName)
		} else {
			fmt.Printf("Appliance '%s' turned off.\n", applianceName)
		}
	} else {
		fmt.Println("Appliance not found.")
	}
}

func TriggerHeatingCooling(reader *bufio.Reader) {
	fmt.Print("Enter desired temperature: ")
	tempStr, _ := reader.ReadString('\n')
	tempStr = strings.TrimSpace(tempStr)
	temp, err := strconv.Atoi(tempStr)
	if err != nil {
		fmt.Println("Invalid temperature.")
	} else {
		if temp < thermostat {
			fmt.Println("Heating triggered.")
		} else if temp > thermostat {
			fmt.Println("Cooling triggered.")
		} else {
			fmt.Println("Temperature is already at desired level.")
		}
	}
}

func CheckEnergyUsage() {
	// Simulated energy usage calculation
	// API replace this with actual calculations based on device usage
	fmt.Println("Energy usage: 100 kWh")
}

func CheckSecurityStatus() {
	if security {
		fmt.Println("House is secure.")
	} else {
		fmt.Println("House is not secure.")
	}
}
