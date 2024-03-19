package JWT

import (
	mongodb_dal "CMPSC488SP24SecThursday/bam"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// jwtSecret the secret key used for token verification
var jwtSecret = []byte("your-secret-key") // To be encrypted

// jwtExpiration shelf life of the token
const jwtExpiration = 24 * time.Second

// GenerateJWT convert ClientInfo object into a token
/*
Purpose: Convert ClientINfo object into token to verify the client.
	client: ClientInfo object
	return: Token string and error report
*/
func GenerateJWT(user mongodb_dal.User) (string, error) {
	// Create a new token object
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	validClient, err := isValidUsername(user) // Verification Process Stub

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
