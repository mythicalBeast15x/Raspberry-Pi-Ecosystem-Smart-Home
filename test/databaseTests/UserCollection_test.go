package databaseTests

import (
	mongodb_dal "CMPSC488SP24SecThursday/bam"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"testing"
)

func TestUserCollectionOperation(t *testing.T) {
	// Initialize MongoDBDAL

	dal, err := mongodb_dal.ConnectToDatabase()
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

	// Retrieve all accounts in the database and have then in a list
	userList, err := dal.ExtractUserInfo()
	if err != nil {
		log.Fatal("Failed to retrieve users:", err)
	}

	fmt.Println()

	// Print the list of users
	fmt.Println("List of Users:")
	for _, user := range userList {
		fmt.Println("User:", user)
	}

	fmt.Println()

	replaceErr := dal.ReplaceUsers(userList)
	if err != nil {
		fmt.Println("Error replacing users:", replaceErr)
	} else {
		fmt.Println("Users replaced successfully")
	}

	fmt.Println()

	// Serialize user list
	fileName := "users.json"
	err = mongodb_dal.SerializeUsersToJSON(userList, fileName)
	if err != nil {
		log.Fatal(err)
	}

	// Deserialize user list
	deserializeUserList, err := mongodb_dal.DeserializeUsersFromJSON(fileName)
	if err != nil {
		log.Fatal(err)
	}
	if !mongodb_dal.UserListMatch(userList, deserializeUserList) {
		log.Fatal("Users list does not match")
	} else {
		fmt.Println("Users list match")
	}

	fmt.Println()

	// Finding a user on the database
	var result mongodb_dal.User
	filter := bson.D{{Key: "username", Value: "exampleUser1"}}
	//filter := bam.User{Username: "exampleUser1"}
	err = dal.FindOne(filter, &result)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Found user:", result)

	// Comparing password
	err = mongodb_dal.ComparePasswords(result.Password, "secret_password_1")
	if err != nil {
		fmt.Println("Password does not match")
	} else {
		fmt.Println("Password matches")
	}

	// Test document deletion and dropping db
	deleteFilter := bson.D{{Key: "username", Value: "exampleUser1"}}
	deleteErr := dal.DeleteUser(deleteFilter)
	if deleteErr != nil {
		log.Fatal("Failed to delete document:", err)
	}

	// Call the DeleteAllUsers method to delete all documents in the user collection
	err = dal.DeleteAllUsers()
	if err != nil {
		log.Fatal(err)
	}

	// Drop the entire database
	err = dal.DropDatabase()
	if err != nil {
		log.Fatal("\nFailed to drop the database:", err)
	}
	fmt.Println("\nDatabase dropped successfully.")

	// Retrieve all accounts in the database and have then in a list
	userList1, err := dal.ExtractUserInfo()
	if err != nil {
		log.Fatal("Failed to retrieve users:", err)
	}

	fmt.Println()

	// Print the list of users
	fmt.Println("List of Users:")
	for _, user := range userList1 {
		fmt.Println("User:", user)
	}

	fmt.Println()

}
