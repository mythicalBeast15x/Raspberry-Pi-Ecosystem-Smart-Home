package lighting

import (
	"fmt"
	"time"
)

// Lighting represents a lighting fixture with on/off and brightness control.
/*
Purpose: Defines a struct representing a lighting fixture with various attributes.
The Lighting struct encapsulates properties of a lighting fixture, including its name, state (on/off),
brightness level, color, and power consumption. Instances of this struct can be created to model individual
lighting fixtures within a system.
*/
type Lighting struct {
	Name             string  // For example: "Living Room Light"
	State            bool    // On or off
	Brightness       float32 // 0.0-100.0 (0.0% to 100.0% brightness)
	Color            string  // White, additional colors to be specified
	PowerConsumption int     // in watts
}

// NewLighting creates a new Lighting instance with the given name and initial state.
/*
Purpose: This function initializes a new instance of the Lighting struct with the specified name and
initial state. It sets default values for brightness, color, and power consumption. The newly created
Lighting instance is returned as a pointer.
Params:
    name : string - the name of the lighting fixture
    initialState : bool - the initial state of the lighting (true for on, false for off)
Returns:
    *Lighting - a pointer to the newly created Lighting instance
*/
func NewLighting(name string, initialState bool) *Lighting {
	return &Lighting{
		Name:             name,         // Defaults as "Light" for now
		State:            initialState, // Defaults as off
		Brightness:       0.0,          // Defaults with 0.0 brightness
		Color:            "white",      // Defaults with the color white
		PowerConsumption: 0,            // Defaults with 0 watts in usage
	}
}

// TurnOn turns the lighting on.
/*
Purpose: This method changes the state of the lighting instance to "on" and prints a message indicating
that the lighting is now turned on.
*/
func (l *Lighting) TurnOn() bool {
	l.State = true
	fmt.Printf("%s is now turned ON\n", l.Name)
	return true
}

// TurnOff turns the lighting off.
/*
Purpose: This method changes the state of the lighting instance to "off" and prints a message indicating
that the lighting is now turned off.
*/
func (l *Lighting) TurnOff() bool {
	l.State = false
	fmt.Printf("%s is now turned OFF\n", l.Name)
	return true
}

// SetTimer sets a timer to turn on or off the lighting after a specified duration.
/*
Purpose: This function checks the current state of the lighting instance and schedules it to
turn on or off after the specified number of seconds. It prints a message indicating when the action
will occur.
Params:
    seconds : int - the duration in seconds after which the action will be taken
*/
func (l *Lighting) SetTimer(seconds int) {
	if l.State == true {
		fmt.Printf("%s will turn off in %d seconds\n", l.Name, seconds)
		time.Sleep(time.Duration(seconds) * time.Second)
		l.TurnOff()
	} else {
		fmt.Printf("%s will turn on in %d seconds\n", l.Name, seconds)
		time.Sleep(time.Duration(seconds) * time.Second)
		l.TurnOn()
	}
}

// SetBrightness sets the brightness of the lighting.
/*
Purpose: This method sets the brightness of the lighting instance to the provided value, ensuring that it
falls within the valid range (0.0 to 100.0). It optionally prints a message indicating the new
brightness value.
Params:
    brightness : float32 - the brightness value to set (0.0-100.0)
    print : bool - whether to print a message indicating the brightness change
*/
func (l *Lighting) SetBrightness(brightness float32, print bool) bool {
	if brightness < 0.0 {
		brightness = 0.0
	} else if brightness > 100.0 {
		brightness = 100.0
	}
	l.Brightness = brightness
	if print {
		fmt.Printf("%s brightness is set to %.0f\n", l.Name, l.Brightness)
	}
	return true
}

// AdjustBrightnessOverTime gradually adjusts the brightness of the lighting over a specified duration.
/*
Purpose: This function calculates the change in brightness per second required to reach the target
brightness over the specified duration. It then incrementally adjusts the brightness over time,
ensuring smooth transition. It prints messages indicating the ongoing adjustment.
Params:
    targetBrightness : int - the target brightness value to reach (0-100)
    seconds : int - the duration in seconds over which to adjust the brightness
*/
func (l *Lighting) AdjustBrightnessOverTime(targetBrightness int, seconds int) {
	fmt.Printf("%s brightness will gradually change to %d over %d seconds\n", l.Name, targetBrightness, seconds)
	// Calculate the change in brightness per second.
	deltaBrightness := (float32(targetBrightness) - l.Brightness) / float32(seconds)
	// Define a ticker with a tick interval of 1 second.
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	// Loop until the specified duration is reached.
	for i := 0; i < seconds; i++ {
		// Incrementally adjust the brightness.
		l.SetBrightness(l.Brightness+deltaBrightness, true)
		// Wait for the next tick.
		<-ticker.C
	}
	// Ensure that the target brightness is reached exactly at the end.
	l.SetBrightness(float32(targetBrightness), false)
}

// SetColor sets the color of the lighting.
/*
Purpose: This method updates the color attribute of the lighting instance to the provided color value.
It optionally prints a message indicating the color change.
Params:
    color : string - the color value to set
    print : bool - whether to print a message indicating the color change
*/
func (l *Lighting) SetColor(color string, print bool) bool {
	l.Color = color
	if print {
		fmt.Printf("%s color is set to %s\n", l.Name, l.Color)
	}
	return true
}

// ColorCycle initiates cycling through colors for a specified duration.
/*
Purpose: This function cycles through a list of colors, updating the color attribute of the lighting
instance at regular intervals. It prints messages indicating the ongoing color change process.
Params:
    seconds : int - the duration in seconds for which to cycle through colors
*/
func (l *Lighting) ColorCycle(seconds int) {
	fmt.Printf("%s will cycle colors for %d seconds\n", l.Name, seconds)
	// Define a list of colors to cycle through.
	colors := []string{"red", "orange", "white", "yellow", "green", "blue", "purple"}
	// Define a ticker with a tick interval of 1 second.
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	// Counter to keep track of color changes.
	colorChangeCount := 0
	// Loop until the total number of color changes reaches the required duration.
	for colorChangeCount < seconds {
		// Iterate through colors.
		for _, color := range colors {
			// Break the loop if the required duration is reached.
			if colorChangeCount >= seconds {
				break
			}
			// Set the color.
			l.SetColor(color, true)
			// Increment the color change count.
			colorChangeCount++
			// Wait for the next tick to change color.
			<-ticker.C
		}
	}
}

// SaveScene saves the current state, brightness, color, etc. as a preset scene.
/*
Purpose: This function is responsible for saving the current state of the lighting instance, including its
brightness, color, and other attributes, as a named scene. It prints a message confirming the
successful saving of the scene.
Params:
    name : string - the name of the scene to save
*/
func (l *Lighting) SaveScene(name string) {
	// Implement save scene logic
	fmt.Printf("Scene '%s' saved for %s\n", name, l.Name)
}

// RecallScene sets the light device to the saved scene.
/*
Purpose: This function retrieves a previously saved scene by name and applies its settings to the lighting
instance, restoring its state, brightness, color, etc., to match those saved in the scene.
It prints a message confirming the successful recall of the scene.
Params:
    name : string - the name of the scene to recall
*/
func (l *Lighting) RecallScene(name string) {
	// Implement recall scene logic
	fmt.Printf("Recalling scene '%s' for %s\n", name, l.Name)
}

/*
func main() {
	light := NewLighting("Light", false)
	light.TurnOn()
	time.Sleep(2 * time.Second)
	light.TurnOff()
	time.Sleep(2 * time.Second)
	light.SetTimer(5)
	light.SetBrightness(50, true)
	time.Sleep(2 * time.Second)
	light.AdjustBrightnessOverTime(100, 8)
	time.Sleep(2 * time.Second)
	light.SetColor("blue", true)
	time.Sleep(2 * time.Second)
	light.ColorCycle(10)
	time.Sleep(2 * time.Second)
	light.SaveScene("Evening")
	time.Sleep(2 * time.Second)
	light.RecallScene("Evening")
	time.Sleep(2 * time.Second)
	light.SetTimer(5)
}
*/
