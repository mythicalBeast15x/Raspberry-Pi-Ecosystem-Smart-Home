package mongodb_dal

import (
	"golang.org/x/crypto/bcrypt"
)

// User represents the structure of a user in the database
type User struct {
	Username string `bson:"username"`
	Password string `bson:"password"`
}

// MongoURI MongoDB connection URI
const (
	MongoURI = "mongodb://localhost:27017"
)

// HashPassword hashes the given password using bcrypt
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// ComparePasswords compares a hashed password with its possible plaintext equivalent
func ComparePasswords(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err
}
