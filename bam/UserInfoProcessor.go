package mongodb_dal

import (
	"encoding/json"
	"errors"
	"os"
)

// SerializeUsersToJSON serializes the list of users to a JSON file
func SerializeUsersToJSON(users []User, filename string) error {
	// Open the file for writing
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a JSON encoder
	encoder := json.NewEncoder(file)

	// Write the users slice to the file
	if err := encoder.Encode(users); err != nil {
		return err
	}

	return nil
}

// DeserializeUsersFromJSON deserializes user data from a JSON file
func DeserializeUsersFromJSON(filename string) ([]User, error) {
	// Open the file for reading
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a slice to hold the decoded users
	var users []User

	// Create a JSON decoder
	decoder := json.NewDecoder(file)

	// Decode the JSON data into the users slice
	if err := decoder.Decode(&users); err != nil {
		return nil, err
	}

	return users, nil
}

// UserListMatch compares two slices of User structs and returns true if they are identical, false otherwise
func UserListMatch(usersList1, userList2 []User) bool {
	// Check if the lengths of the slices are equal
	if len(usersList1) != len(userList2) {
		return false
	}

	// Iterate over each element in the slices
	for i := range usersList1 {
		// Compare the Username, Password, and Token fields
		if usersList1[i].Username != userList2[i].Username ||
			usersList1[i].Password != userList2[i].Password ||
			usersList1[i].Token != userList2[i].Token {
			return false
		}
	}

	// If all elements are identical, return true
	return true
}

func UpdateUserToken(userList []User, username, newToken string) error {
	// Search for the user by username
	for i, user := range userList {
		if user.Username == username {
			// Update the token
			userList[i].Token = newToken
			return nil // User found and token updated, return nil (no error)
		}
	}
	// If the loop completes without finding the user, return an error
	return errors.New("user not found")
}

func UpdateUsername(userList []User, oldUsername, newUsername string) error {
	// Search for the user by old username
	for i, user := range userList {
		if user.Username == oldUsername {
			// Update the username
			userList[i].Username = newUsername
			return nil // User found and username updated, return nil (no error)
		}
	}
	// If the loop completes without finding the user, return an error
	return errors.New("user not found")
}

func UpdatePassword(userList []User, username, newPassword string) error {
	// Search for the user by username
	for i, user := range userList {
		if user.Username == username {
			// Update the password (You may want to hash the new password before updating)
			userList[i].Password = newPassword
			return nil // User found and password updated, return nil (no error)
		}
	}
	// If the loop completes without finding the user, return an error
	return errors.New("user not found")
}
