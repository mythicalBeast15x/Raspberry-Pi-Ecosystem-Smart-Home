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

// SetUser to set up user data structure
/*
Purpose:
username: the name of the user
password: the password of the user
return: user data structure
*/
func SetUser(username string, password string) User {
	newUser := User{
		Username: username,
		Password: password,
		Token:    "",
	}
	return newUser
}

// UsernameExists to check if a user with the username already exit.
/*
Purpose: Organize a database ensure no user has the same name.
username: The name to check
return: boolean and err
*/
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
/*
Purpose: take record of the user in the database
User: the user to add to the database
Return: result and error
*/
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
/*
Purpose: to be able to send to another pi
Return: List of users and error
*/
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
/*
Purpose: to protect password
Return: hashed password
*/
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// ComparePasswords compares a hashed password with its possible plaintext equivalent
/*
Purpose: to verify password
hashedPassword: hashed password
password: the correct password
return error
*/
func ComparePasswords(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err
}

// DeleteUser to remove user
/*
Purpose: manage database
deleteFilter: to target the user
return: error
*/
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

// DeleteAllUsers to remove all users from the database
/*
Purpose: to manage a database
return: error
*/
func (d *MongoDBDAL) DeleteAllUsers() error {
	// Define an empty filter to match all documents
	filter := bson.D{}

	// Perform the delete operation
	deleteResult, err := d.Collection.DeleteMany(context.Background(), filter)
	if err != nil {
		return err
	}

	// Print delete result
	fmt.Println("Deleted", deleteResult.DeletedCount, "documents")

	return nil
}

// ReplaceUsers to refill the database with the updated userlist
/*
Purpose: to update the database
userList: the updated user list
return: error
*/
func (d *MongoDBDAL) ReplaceUsers(userList []User) error {
	// Delete all existing users
	if err := d.DeleteAllUsers(); err != nil {
		return err
	}

	// Add new users from the provided user list
	for _, user := range userList {
		_, err := d.AddUser(user)
		if err != nil {
			return err
		}
	}

	return nil
}
