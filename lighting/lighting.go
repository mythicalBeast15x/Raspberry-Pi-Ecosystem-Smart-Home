package lighting

import (
	"fmt"
	"sync"
	"time"
)

// Lighting represents a lighting fixture with on/off and brightness control.
type Lighting struct {
	Name             string     // For example: "Living Room Light"
	State            bool       // On or off
	Brightness       float64    // 0.0-100.0 (0.0% to 100.0% brightness)
	Color            string     // White, additional colors to be specified
	PowerConsumption float64    // in watts
	brightnessMutex  sync.Mutex // This ensures that only one method can modify the brightness at a time
}

// NewLighting creates a new Lighting instance with the given name and initial state.
func NewLighting(name string, initialState bool) *Lighting {
	return &Lighting{
		Name:             name,         // Defaults as "Light" for now
		State:            initialState, // Defaults as off
		Brightness:       0,            // Defaults with 0 brightness
		Color:            "white",      // Defaults with the color white
		PowerConsumption: 0,            // Defaults with 0 watts in usage
	}
}

// TurnOn turns the lighting on.
func (l *Lighting) TurnOn() {
	l.State = true
	fmt.Printf("%s is now turned ON\n", l.Name)
}

// TurnOff turns the lighting off.
func (l *Lighting) TurnOff() {
	l.State = false
	fmt.Printf("%s is now turned OFF\n", l.Name)
}

// GetState returns the current state of the lighting (on/off).
func (l *Lighting) GetState() bool {
	return l.State
}

// SetBrightness sets the brightness of the lighting.
func (l *Lighting) SetBrightness(brightness float64, print bool) {
	l.brightnessMutex.Lock()
	defer l.brightnessMutex.Unlock()
	if brightness < 0.0 {
		brightness = 0.0
	} else if brightness > 100.0 {
		brightness = 100.0
	}
	l.Brightness = brightness
	if print {
		fmt.Printf("%s brightness is set to %.0f\n", l.Name, l.Brightness)
	}
}

// AdjustBrightnessOverTime gradually adjusts the brightness of the lighting over a specified duration.
func (l *Lighting) AdjustBrightnessOverTime(targetBrightness int, duration int) {
	l.brightnessMutex.Lock()
	defer l.brightnessMutex.Unlock()
	// Calculate the change in brightness per second.
	deltaBrightness := (float64(targetBrightness) - l.Brightness) / float64(duration)
	go func() {
		// Define a ticker with a tick interval of 1 second.
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		// Loop until the specified duration is reached.
		for i := 0; i < duration; i++ {
			// Incrementally adjust the brightness.
			l.SetBrightness(l.Brightness+deltaBrightness, false)
			// Wait for the next tick.
			<-ticker.C
		}
		// Ensure that the target brightness is reached exactly at the end.
		l.SetBrightness(float64(targetBrightness), false)
	}() // Return immediately without waiting for the adjustment to complete.
	fmt.Printf("%s brightness will gradually change to %d%% over %d seconds\n", l.Name, targetBrightness, duration)
}

// SetColor sets the color of the lighting.
func (l *Lighting) SetColor(color string) {
	l.Color = color
	fmt.Printf("%s color is set to %s\n", l.Name, l.Color)
}

// ColorCycle initiates cycling through colors for a specified duration.
func (l *Lighting) ColorCycle(duration int) {
	// Implement color cycling logic
	fmt.Printf("%s is now cycling colors for %d seconds\n", l.Name, duration)
}

// GetPowerConsumption returns the power consumption of the lighting.
func (l *Lighting) GetPowerConsumption() float64 {
	return l.PowerConsumption
}

// SetTimer sets a timer to turn on or off the lighting after a specified duration.
func (l *Lighting) SetTimer(minutes int) {
	fmt.Printf("%s will turn off in %d minutes\n", l.Name, minutes)
	// Implement timer logic
	go func() {
		time.Sleep(time.Duration(minutes) * time.Minute)
		l.TurnOff()
	}()
}

// SaveScene saves the current state, brightness, color, etc. as a preset scene.
func (l *Lighting) SaveScene(name string) {
	// Implement save scene logic
	fmt.Printf("Scene '%s' saved for %s\n", name, l.Name)
}

// RecallScene sets the light device to the saved scene.
func (l *Lighting) RecallScene(name string) {
	// Implement recall scene logic
	fmt.Printf("Recalling scene '%s' for %s\n", name, l.Name)
}

/*
func main() {
	light := NewLighting("Light", false)
	light.TurnOn()
	light.SetBrightness(50, true)
	light.AdjustBrightnessOverTime(100, 30)
	light.SetColor("blue")
	light.ColorCycle(60)
	light.SetTimer(30)
	light.SaveScene("Evening")
	light.RecallScene("Evening")
	light.TurnOff()
}
*/
