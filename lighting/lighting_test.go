package main

import (
	"testing"
	"time"
)

func TestLighting(t *testing.T) {
	light := NewLighting("TestLight", false)

	// Test TurnOn and TurnOff methods
	light.TurnOn()
	if !light.State {
		t.Errorf("Expected light to be turned on, but it's turned off")
	}
	light.TurnOff()
	if light.State {
		t.Errorf("Expected light to be turned off, but it's turned on")
	}

	// Test SetTimer method
	light.TurnOn()
	go light.SetTimer(2)
	time.Sleep(3 * time.Second)
	if light.State {
		t.Errorf("Expected light to be turned off after timer, but it's still on")
	}

	// Test SetBrightness method
	light.TurnOn()
	light.SetBrightness(50, true)
	if light.Brightness != 50 {
		t.Errorf("Expected brightness to be 50, got %f", light.Brightness)
	}

	// Test AdjustBrightnessOverTime method
	light.SetBrightness(0, false) // Reset brightness
	light.TurnOn()
	go light.AdjustBrightnessOverTime(100, 5)
	time.Sleep(6 * time.Second)
	if light.Brightness != 100 {
		t.Errorf("Expected brightness to be 100 after adjustment, got %f", light.Brightness)
	}

	// Test SetColor method
	light.SetColor("red", true)
	if light.Color != "red" {
		t.Errorf("Expected color to be red, got %s", light.Color)
	}

	// Test ColorCycle method
	go light.ColorCycle(7)
	time.Sleep(8 * time.Second)
	if light.Color != "purple" {
		t.Errorf("Expected color to be purple after cycle, got %s", light.Color)
	}
}
