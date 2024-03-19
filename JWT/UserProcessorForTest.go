package JWT

import (
	mongodb_dal "CMPSC488SP24SecThursday/bam"
	"errors"
	"log"
)

var verifiedUsers []mongodb_dal.User // Declare verifiedUsers at package level

func init() {
	dal, err := mongodb_dal.ConnectToDatabase()
	if err != nil {
		log.Fatal("Failed to initialize MongoDBDAL:", err)
	}
	defer dal.Close() // Defer closing the connection
	verifiedUsers, err = dal.ExtractUserInfo()
	if err != nil {
		log.Fatal("Failed to retrieve users:", err)
	}

}

// Process to verify username possibly other info as well.
// verifyUser stub to verify the entered username and password
/*
Purpose: to verify if the username and password combo are registered
	username: for username
	password: for password
	return: bool
*/
func verifyUser(username string, password string) bool {
	if username == username && password == password {
		return true
	}
	return false
}

// generateProfile helper function to run username and password to a ClientInfo object
/*
Purpose: to turn username and password combo to a ClientInfo object for better handling
	username: for username
	password: for password
	return: ClientInfo object
*/
func generateProfile(username string, password string) mongodb_dal.User {
	profile := mongodb_dal.User{Username: username, Password: password, Token: ""}
	return profile
}

// appendVerifiedClient to append verified clients into the stub database
/*
Purpose to append verified clients into stub database
	username: for username
	password: for password
	return: void when if the username and password combo are valid, error if not.
*/
func appendVerifiedClient(username string, password string) error {
	if verifyUser(username, password) {
		newClient := generateProfile(username, password)
		verifiedUsers = append(verifiedUsers, newClient)
	}
	return errors.New("username could not be verified")
}

// Stub verification
// isValidUsername check if the client info is valid
/*
Purpose: Check if the clientInfo is in the verified client database.
	clientInfo: ClientInfo object
	return: ClientInfo object, and error report
*/
func isValidUsername(user mongodb_dal.User) (mongodb_dal.User, error) {
	for _, i := range verifiedUsers {
		if user.Username == i.Username && user.Password == i.Password {
			return user, nil
		}
	}
	return mongodb_dal.User{}, errors.New("invalid username or password")
}

// Stub client verification process end ----
