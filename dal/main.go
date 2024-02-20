package main

//go get go.mongodb.org/mongo-driver/mongo
//files to be imported: Appliances, lights, HVAC
import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Define a struct to represent your data model
type User struct {
	ID       string `bson:"_id,omitempty"`
	Username string `bson:"username"`
	Email    string `bson:"email"`
}

// MongoDB configuration
var (
	mongoURI       = "mongodb://localhost:27017" // MongoDB server URI
	dbName         = "mydb"                      // Database name
	collectionName = "users"                     // Collection name
)

// Connect to MongoDB and return a MongoDB client
func connectToMongoDB() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Ping the MongoDB server to verify the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")
	return client, nil
}

// Create a new user in the MongoDB database
func createUser(client *mongo.Client, user User) error {
	collection := client.Database(dbName).Collection(collectionName)
	_, err := collection.InsertOne(context.Background(), user)
	return err
}

////Lights communicates with Zigbee Devices & Adjusts device status
///* Purpose: Receive a request from the front-end,
//	Send that request to relevant Zigbee Device &
//	Return that Status to the Front End
//RequestedStatus: Takes String for Requested Status of Lights: ex: On/Off
//Returns Bool (True if request is successful, False if not)
//(could be broken into two functions)
//
//
//*/
//func Lights(RequestedStatus string) bool {
//	//if RequestedStatus = on
//	//if (Lighting->main.go->Turn On()) -> (if returns true)
//	//return response to front end ->encourage message
//	//if Requested Status = off
//	//if (Lighting -> main.go-> Turn Off()) -> if returns true
//	//return response to front end
//	//else
//	//false error handling
//
//	return false
//}
//
////setTemp communicates with Zigbee Devices to adjust HVAC temp
///* Purpose: Receive a request from the front-end,
//	Send that request to relevant Zigbee Device &
//	Return that Status to the Front End
//Temperature: Takes integer, front end should be able to limit range, but if not an error will return
//Returns T/F depending on response
//*/
//func setTemp(Temperature int) bool {
//	//if Temperature in range
//	//if ( HVAC -> set temp () ) -> if response success
//	//return to front
//	//else
//	return false
//
//}
////Should there be a list here?
//func checkApplianceStatus(appliancetocheck string) bool {
//	//if appliance in (appliance listings)
//		// if Appliance -> checkStatus() -> response(true)
//			//return to front end appliance status
//	// else -> return response to front end
//}

func main() {
	// Connect to MongoDB
	client, err := connectToMongoDB()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Example: Creating a new user
	newUser := User{
		Username: "john_doe",
		Email:    "john@example.com",
	}

	err = createUser(client, newUser)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User created successfully!")
}
