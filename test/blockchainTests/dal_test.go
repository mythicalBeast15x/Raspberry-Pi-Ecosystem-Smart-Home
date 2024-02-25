package blockchainTests

import (
	dal "CMPSC488SP24SecThursday/dal"
	hvac "CMPSC488SP24SecThursday/hvac"
	light "CMPSC488SP24SecThursday/lighting"
	"CMPSC488SP24SecThursday/security"
	"testing"
)

// testing the block creations
func TestLights(t *testing.T) {
	testLight := light.NewLighting("livingroom", false)
	testResult := dal.Light(0, testLight, 100, "white")
	//Test for Validating a correct block
	if testResult != true {
		t.Errorf("Expected: %t, Got: %t", true, testResult)
	}
}

func TestHVAC(t *testing.T) {
	testHVAC := hvac.NewHVAC("fan")
	testResult := dal.HVAC(9, testHVAC, 100, 100, "cool")
	if testResult != true {
		t.Errorf("Expected: %t, Got: %t", true, testResult)
	}

}

func TestSecurity(t *testing.T) {
	testSecurity := security.NewAlarm("Fire")
	testResult := dal.SecuritySystem(6, testSecurity)
	if testResult != true {
		t.Errorf("Expected: %t, Got: %t", true, testResult)
	}
}
