package blockchainTests

import (
	dal "CMPSC488SP24SecThursday/dal"
	light "CMPSC488SP24SecThursday/lighting"
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
