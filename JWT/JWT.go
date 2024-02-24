package JWT

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// jwtSecret the secret key used for token verification
var jwtSecret = []byte("your-secret-key") // To be encrypted

// ClientInfo stub for holding client info
type ClientInfo struct {
	Username string
	Password string
}

// jwtExpiration shelf life of the token
const jwtExpiration = 24 * time.Second

// Stub client verification process start ----
// verifiedClients stub for the database of verified clients
var verifiedClients = []ClientInfo{
	{Username: "Sam", Password: "123"},
	{Username: "Bob", Password: "456"},
	{Username: "Tom", Password: "789"}}

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
func generateProfile(username string, password string) ClientInfo {
	profile := ClientInfo{Username: username, Password: password}
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
		verifiedClients = append(verifiedClients, newClient)
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
func isValidUsername(clientInfo ClientInfo) (ClientInfo, error) {
	for _, i := range verifiedClients {
		if clientInfo.Username == i.Username && clientInfo.Password == i.Password {
			return clientInfo, nil
		}
	}
	return ClientInfo{}, errors.New("client can't be verified")
}

// Stub client verification process end ----

// GenerateJWT convert ClientInfo object into a token
/*
Purpose: Convert ClientINfo object into token to verify the client.
	client: ClientInfo object
	return: Token string and error report
*/
func GenerateJWT(client ClientInfo) (string, error) {
	// Create a new token object
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	validClient, err := isValidUsername(client)

	if err != nil {
		return "", err
	}

	// Set claims (payload) for the token
	claims["username"] = validClient.Username
	claims["exp"] = time.Now().Add(jwtExpiration).Unix() // Token expires in 24 hours

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// VerifyJWT to verify if the token is valid
/*
Purpose: User to have been verified inorder to gain access
	tokenString: User token
	return: token object and error report
*/
func VerifyJWT(tokenString string) (*jwt.Token, error) {
	// Parse the token string
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verify the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Return the secret key for validation
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	// Verify if the token is valid
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}
