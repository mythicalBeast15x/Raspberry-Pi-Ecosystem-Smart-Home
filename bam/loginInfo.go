package mongodb_dal

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"log"
)

// User represents the structure of a user in the database
type User struct {
	Username string `bson:"username"`
	Password string `bson:"password"`
	Token    string `bson:"token"`
}

// MongoURI MongoDB connection URI
const (
	MongoURI = "mongodb://localhost:27017"
)

func (d *MongoDBDAL) UsernameExists(username string) (bool, error) {
	// Define a filter to check if the username exists
	filter := bson.M{"username": username}

	// Perform a find operation with the specified filter
	result := d.Collection.FindOne(context.Background(), filter)

	// Check if any document was found
	if err := result.Err(); errors.Is(err, mongo.ErrNoDocuments) {
		// Username does not exist
		return false, nil
	} else if err != nil {
		// Error occurred during the find operation
		return false, err
	}

	// Username exists
	return true, nil
}

var ErrUsernameExists = errors.New("username already exists")

// AddUser inserts a new user into the database
func (d *MongoDBDAL) AddUser(user User) (interface{}, error) {
	// Check if the username already exists in the database
	exists, err := d.UsernameExists(user.Username)
	if err != nil {
		return nil, err
	}
	if exists {
		fmt.Println("Username:", user.Username, "is not available")
		return nil, ErrUsernameExists
	}

	// Insert the user into the database
	insertResult, err := d.InsertOne(user)
	if err != nil {
		return nil, err
	}
	return insertResult.InsertedID, nil
}

// ExtractUserInfo GetUsers retrieves all users from the database and returns them as a slice
func (d *MongoDBDAL) ExtractUserInfo() ([]User, error) {
	// Initialize a slice to hold the data
	var userList []User

	// Query all documents from the collection
	cursor, err := d.Collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	// Iterate over the cursor and append each document to the slice
	for cursor.Next(context.Background()) {
		var user User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		userList = append(userList, user)
	}

	// Check if there was an error during cursor iteration
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// Return the list of users
	return userList, nil
}

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

func (d *MongoDBDAL) DeleteUser(deleteFilter interface{}) error {
	// Delete the document
	deleteResult, err := d.DeleteOne(deleteFilter)
	if err != nil {
		return err
	}

	// Check if any document was deleted
	if deleteResult.DeletedCount == 0 {
		log.Println("No document deleted")
	} else {
		log.Println("Document deleted successfully")
	}

	return nil
}

func (d *MongoDBDAL) DeleteAllUsers() error {
	// Define an empty filter to match all documents
	filter := bson.D{}

	// Perform the delete operation
	deleteResult, err := d.Collection.DeleteMany(context.Background(), filter)
	if err != nil {
		return err
	}

	// Print delete result
	fmt.Println(deleteResult)
	fmt.Println("Deleted", deleteResult.DeletedCount, "documents")

	return nil
}
