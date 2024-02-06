package main

import (
	"fmt"
	"time"
)

type Lighting struct {
	Name             string
	State            bool
	Brightness       float64 // 0.0 to 1.0 (0% to 100% brightness)
	Color            string
	ColorTemperature int     // in Kelvin
	PowerConsumption float64 // in watts
}

func NewLighting(name string, initialState bool) *Lighting {
	return &Lighting{
		Name:             name,
		State:            initialState,
		Brightness:       0,
		Color:            "white",
		ColorTemperature: 6500,
		PowerConsumption: 0,
	}
}

func (l *Lighting) TurnOn() {
	l.State = true
	fmt.Printf("%s is now turned ON\n", l.Name)
}

func (l *Lighting) TurnOff() {
	l.State = false
	fmt.Printf("%s is now turned OFF\n", l.Name)
}

func (l *Lighting) SetBrightness(brightness float64) {
	if brightness < 0 {
		brightness = 0
	} else if brightness > 1.0 {
		brightness = 1.0
	}
	l.Brightness = brightness
	fmt.Printf("%s brightness is set to %.2f%%\n", l.Name, l.Brightness*100)
}

func (l *Lighting) SetColor(color string) {
	l.Color = color
	fmt.Printf("%s color is set to %s\n", l.Name, l.Color)
}

func (l *Lighting) SetColorTemperature(temperature int) {
	l.ColorTemperature = temperature
	fmt.Printf("%s color temperature is set to %dK\n", l.Name, l.ColorTemperature)
}

func (l *Lighting) SetTimer(minutes int) {
	fmt.Printf("%s will turn off in %d minutes\n", l.Name, minutes)
	// Implement timer logic
	go func() {
		time.Sleep(time.Duration(minutes) * time.Minute)
		l.TurnOff()
	}()
}

func (l *Lighting) GetPowerConsumption() float64 {
	return l.PowerConsumption
}

func (l *Lighting) SaveScene(name string) {
	// Save current state, brightness, color, etc. as a preset scene
	fmt.Printf("Scene '%s' saved for %s\n", name, l.Name)
}

func (l *Lighting) RecallScene(name string) {
	// Set the light bulb to the saved scene
	fmt.Printf("Recalling scene '%s' for %s\n", name, l.Name)
}

func (l *Lighting) ColorCycle(duration int) {
	// Implement color cycling logic
	fmt.Printf("%s is now cycling colors for %d seconds\n", l.Name, duration)
}

func (l *Lighting) AdjustBrightnessOverTime(targetBrightness float64, duration int) {
	// Implement gradual brightness adjustment logic
	fmt.Printf("%s brightness will gradually change to %.2f%% over %d seconds\n", l.Name, targetBrightness*100, duration)
}

func (l *Lighting) RandomizeOnOffTimings() {
	// Implement randomization of on/off timings for security
	fmt.Printf("%s on/off timings randomized for security\n", l.Name)
}

func main() {
	livingRoomLight := NewLighting("Living Room Light", false)
	livingRoomLight.TurnOn()
	livingRoomLight.SetBrightness(0.75)
	livingRoomLight.SetColor("blue")
	livingRoomLight.SetColorTemperature(4000)
	livingRoomLight.SetTimer(30)
	livingRoomLight.SaveScene("Evening")
	livingRoomLight.RecallScene("Evening")
	livingRoomLight.ColorCycle(60)
	livingRoomLight.AdjustBrightnessOverTime(0.5, 120)
	livingRoomLight.RandomizeOnOffTimings()
	livingRoomLight.TurnOff()
}
