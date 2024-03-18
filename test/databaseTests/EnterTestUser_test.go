package databaseTests

import (
	mongodb_dal "CMPSC488SP24SecThursday/bam"
	"errors"
	"fmt"
	"log"
	"testing"
)

func TestEnterTestUser(t *testing.T) {
	// Initialize MongoDBDAL

	// MongoURI MongoDB connection URI
	const MongoURI = "mongodb://localhost:27017"

	dbName := "test"
	collectionName := "users"
	dal, err := mongodb_dal.NewMongoDBDAL(MongoURI, dbName, collectionName)
	if err != nil {
		log.Fatal("Failed to initialize MongoDBDAL:", err)
	}
	defer dal.Close() // Defer closing the connection

	// Inserting a user
	hashedPassword1, err := mongodb_dal.HashPassword("skyGreen@123456")
	hashedPassword2, err := mongodb_dal.HashPassword("dogRed@456789")
	hashedPassword3, err := mongodb_dal.HashPassword("catBlue@789123")
	hashedPassword4, err := mongodb_dal.HashPassword("miceBrown@123456")
	if err != nil {
		log.Fatal(err)
	}

	user1 := mongodb_dal.SetUser("Tom", hashedPassword1)
	user2 := mongodb_dal.SetUser("Bob", hashedPassword2)
	user3 := mongodb_dal.SetUser("Sam", hashedPassword3)
	user4 := mongodb_dal.SetUser("David", hashedPassword4)

	fmt.Println("User insertion")
	fmt.Println("User to be inserted:", user1)
	fmt.Println("User to be inserted:", user2)
	fmt.Println("User to be inserted:", user3)
	fmt.Println("User to be inserted:", user4)
	fmt.Println()

	insertedID, err := dal.AddUser(user1)
	if err != nil && !errors.Is(err, mongodb_dal.ErrUsernameExists) {
		t.Errorf("Error inserting user1: %v", err)
	}
	fmt.Println("Inserted a single document: ", insertedID)

	insertedID2, err := dal.AddUser(user2)
	if err != nil && !errors.Is(err, mongodb_dal.ErrUsernameExists) {
		t.Errorf("Error inserting user1: %v", err)
	}
	fmt.Println("Inserted a single document: ", insertedID2)

	insertedID3, err := dal.AddUser(user3)
	if err != nil && !errors.Is(err, mongodb_dal.ErrUsernameExists) {
		t.Errorf("Error inserting user1: %v", err)
	}
	fmt.Println("Inserted a single document: ", insertedID3)

	insertedID4, err := dal.AddUser(user4)
	if err != nil && !errors.Is(err, mongodb_dal.ErrUsernameExists) {
		t.Errorf("Error inserting user1: %v", err)
	}
	fmt.Println("Inserted a single document: ", insertedID4)
}
