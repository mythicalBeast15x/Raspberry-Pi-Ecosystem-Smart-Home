package mongodb_dal

import (
	"encoding/json"
	"errors"
	"os"
)

// SerializeUsersToJSON serializes the list of users to a JSON file
/*
Purpose: to turn the user list in to json file
users: the user list
filename: the Json file to store the serialized user list
return: error
*/
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
/*
Purpose: to turn the serialized user list back to user list
filename: the json file
return: the user list and error
*/
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
/*
Purpose: to test the serialized process work properly
usersList1: list of user
usersList2: list of user
return: boolean
*/
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

// UpdateUserToken update user token
/*
Purpose: to update user token
userList: list of users
userName: targeted user
newToken: updated token
return: error
*/
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

// UpdateUsername update user username
/*
Purpose: to update user token
userList: list of users
oldUsername: targeted user
newUsername: updated user name
return: error
*/
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

// UpdatePassword update user password
/*
Purpose: to update user token
userList: list of users
username: targeted user
newPassword: updated password
return: error
*/
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

// SerializeUserToJSON serializes a single user to a JSON file
/*
Purpose: Serialize a single user to a JSON file
user: The user to be serialized
filename: The JSON file to store the serialized user
return: error
*/
func SerializeUserToJSON(user User, filename string) error {
	// Open the file for writing
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a JSON encoder
	encoder := json.NewEncoder(file)

	// Write the user to the file
	if err := encoder.Encode(user); err != nil {
		return err
	}

	return nil
}

// DeserializeUserFromJSON deserializes user data from a JSON file
/*
Purpose: Deserialize user data from a JSON file
filename: The JSON file containing the serialized user
return: The user and error
*/
func DeserializeUserFromJSON(filename string) (User, error) {
	// Open the file for reading
	file, err := os.Open(filename)
	if err != nil {
		return User{}, err
	}
	defer file.Close()

	// Create a variable to hold the decoded user
	var user User

	// Create a JSON decoder
	decoder := json.NewDecoder(file)

	// Decode the JSON data into the user variable
	if err := decoder.Decode(&user); err != nil {
		return User{}, err
	}

	return user, nil
}
