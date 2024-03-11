package securityTest

import (
	"CMPSC488SP24SecThursday/security"
	"testing"
)

// TestMotionDetection checks if the motion sensor properly detects motion.
func TestMotionDetection(t *testing.T) {
	detected := false
	motionSensor := security.NewMotionSensor("TestMotionSensor", func(name string) {
		detected = true
	})

	// Simulate motion detection.
	motionSensor.DetectMotion()

	if !detected {
		t.Errorf("Motion sensor did not detect motion when it should have.")
	}

	if !motionSensor.MotionDetected {
		t.Errorf("MotionDetected should be true after motion is detected.")
	}
}

// TestDoorLock ensures that the door lock functions correctly.
func TestDoorLock(t *testing.T) {
	doorLock := security.NewDoorLock("TestDoorLock")
	if !doorLock.Locked {
		t.Errorf("New door lock should be locked by default.")
	}

	// Test unlocking the door.
	doorLock.Unlock()
	if doorLock.Locked {
		t.Errorf("Door lock should be unlocked after calling Unlock.")
	}

	// Test locking the door.
	doorLock.Lock()
	if !doorLock.Locked {
		t.Errorf("Door lock should be locked after calling Lock.")
	}
}

// TestCamera ensures that the camera functions correctly.
func TestCamera(t *testing.T) {
	camera := security.NewCamera("TestCamera")
	if camera.Active {
		t.Errorf("New camera should be inactive by default.")
	}

	// Test activating the camera.
	camera.Activate()
	if !camera.Active {
		t.Errorf("Camera should be active after calling Activate.")
	}

	// Test deactivating the camera.
	camera.Deactivate()
	if camera.Active {
		t.Errorf("Camera should be inactive after calling Deactivate.")
	}
}

// Remove TestUserJoiningProcess and TestMotionDetectionAndAlarmTrigger tests as TrustCenter and Alarm are no longer part of the package.
