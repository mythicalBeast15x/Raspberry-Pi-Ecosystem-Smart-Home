package databaseTests

import (
	"CMPSC488SP24SecThursday/UserProcessing"
	mongodb_dal "CMPSC488SP24SecThursday/bam"
	"testing"
)

func TestUpdateUser(t *testing.T) {
	username := "exampleUser1"
	password := "$2a$10$hCgo2q1eSYT6cA.VsAoDpOGy4WtlRNTK3q1loIorWC/savo3D79xO"
	filename := "testCurrentUser.json"

	testUser := mongodb_dal.SetUser(username, password)
	err := UserProcessing.GenerateandUpdateUserToken(testUser, filename)
	if err != nil {
		t.Errorf("error processing user info: %v", err)
	}
}
