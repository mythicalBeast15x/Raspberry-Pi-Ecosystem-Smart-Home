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
	hashedPassword1, err := mongodb_dal.HashPassword("secret_password_1")
	hashedPassword2, err := mongodb_dal.HashPassword("secret_password_2")
	hashedPassword3, err := mongodb_dal.HashPassword("secret_password_3")
	hashedPassword4, err := mongodb_dal.HashPassword("secret_password_4")
	if err != nil {
		log.Fatal(err)
	}

	user1 := mongodb_dal.SetUser("exampleUser1", hashedPassword1)
	user2 := mongodb_dal.SetUser("exampleUser2", hashedPassword2)
	user3 := mongodb_dal.SetUser("exampleUser3", hashedPassword3)
	user4 := mongodb_dal.SetUser("exampleUser4", hashedPassword4)

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
