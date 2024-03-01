package jwtTests

import (
	"CMPSC488SP24SecThursday/JWT"
	"github.com/golang-jwt/jwt/v5"
	"testing"
	"time"
)

func TestGenerateAndVerifyJWT(t *testing.T) {
	// Generate JWT token
	client := JWT.ClientInfo{Username: "Sam", Password: "123"}
	tokenString, err := JWT.GenerateJWT(client)
	if err != nil {
		t.Errorf("Error generating JWT token: %v", err)
	}

	// Verify JWT token
	token, err := JWT.VerifyJWT(tokenString)
	if err != nil {
		t.Errorf("Error verifying JWT token: %v", err)
	}

	// Check token validity
	if !token.Valid {
		t.Error("Generated JWT token is not valid")
	}

	// Check token claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		t.Error("Failed to extract claims from JWT token")
	}
	usernameClaim, ok := claims["username"].(string)
	if !ok {
		t.Error("Failed to extract username claim from JWT token")
	}
	if usernameClaim != client.Username {
		t.Errorf("Invalid username claim in JWT token: got %s, expected %s", usernameClaim, client.Username)
	}

	// Check token expiration
	expirationClaim, ok := claims["exp"].(float64)
	if !ok {
		t.Error("Failed to extract expiration claim from JWT token")
	}
	expirationTime := time.Unix(int64(expirationClaim), 0)
	if expirationTime.Before(time.Now()) {
		t.Error("JWT token has expired")
	}
}