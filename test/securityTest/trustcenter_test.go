package securityTest

/*
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

// TestAlarm checks the functionality of arming, disarming, and triggering the alarm.
func TestAlarm(t *testing.T) {
	alarm := security.NewAlarm("TestAlarm")

	// Test arming the alarm.
	alarm.Arm()
	if !alarm.Armed {
		t.Errorf("Alarm should be armed after calling Arm.")
	}

	// Test triggering the alarm when armed.
	alarm.Trigger()
	if !alarm.Sounded {
		t.Errorf("Alarm should be sounded after calling Trigger when armed.")
	}

	// Reset alarm
	alarm.Sounded = false
	alarm.Disarm()
	if alarm.Armed {
		t.Errorf("Alarm should be disarmed after calling Disarm.")
	}

	// Test triggering the alarm when disarmed.
	alarm.Trigger()
	if alarm.Sounded {
		t.Errorf("Alarm should not be sounded after calling Trigger when disarmed.")
	}
}

// TestTrustCenter checks the user authorization process.
func TestTrustCenter(t *testing.T) {
	tc := security.NewTrustCenter()

	// Authorize a user and check authorization.
	userName := "AuthorizedUser"
	tc.AuthorizeUser(userName)
	if !tc.IsUserAuthorized(userName) {
		t.Errorf("User should be authorized after calling AuthorizeUser.")
	}

	// Attempt to join with an authorized user.
	joinMsg := tc.AttemptJoin(userName)
	expectedMsg := userName + " has successfully joined the network.\n"
	if joinMsg != expectedMsg {
		t.Errorf("Expected '%s', got '%s'", expectedMsg, joinMsg)
	}

	// Attempt to join with an unauthorized user.
	unauthUserName := "UnauthorizedUser"
	joinMsg = tc.AttemptJoin(unauthUserName)
	expectedMsg = unauthUserName + " is not authorized to join the network.\n"
	if joinMsg != expectedMsg {
		t.Errorf("Expected '%s', got '%s'", expectedMsg, joinMsg)
	}
}
*/
