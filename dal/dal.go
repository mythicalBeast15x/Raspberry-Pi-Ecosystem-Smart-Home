package dal

//go get go.mongodb.org/mongo-driver/mongo
//files to be imported: Appliances, lights, HVAC
import (
	hvac "CMPSC488SP24SecThursday/hvac"
	light "CMPSC488SP24SecThursday/lighting"
	security "CMPSC488SP24SecThursday/security"

	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// Key/Value Demo of OIDs
var myMap = map[string][]int{
	"Light":          {1, 2, 3, 4},
	"SecuritySystem": {5, 6, 7},
	"HVAC ":          {9, 10, 11},
}

//Request Disperser
/*Purpose: Takes in Request and depending on the domain, calls domain-specific function w/ args from request command block
 	r-> request block
	returns bool from domain-function ex: Lights
{Light : {Li1: "Turn On" }}
*/
//func RequestDisperser(r request.Request) {
//	if r.Domain == "Light" {
//
//	}
//	//if r.Domain == "HVAC" {
//	//	return
//	//}
//	//if r.Domain == "Security" {
//	//	return
//	//}
//
//}

// //Lights communicates with Zigbee Devices & Adjusts device status
/* Purpose: Receive a request from the front-end,
 Send that request to relevant Zigbee Device &
	Return that Status to the Front End

	OID: Operation ID to dictate what occurs
	l: Existing lighting struct
	brightness: float32 between 0-100
	color: color name
	Returns Bool (True if request is successful, False if not)*/

func Light(OID int, l *light.Lighting, brightness float32, color string) bool {
	switch OID {
	case 0: //turn on lights
		if l.TurnOn() && l.SetBrightness(brightness, true) {
			return true
		} else {
			panic("Lights failed to turn on")
			return false
		}

	case 1: //turn off lights
		if l.TurnOff() && l.SetBrightness(brightness, true) {
			return true
		} else {
			panic("Lights failed to turn off")
			return false
		}
	case 2: //set brightness
		if l.TurnOn() && l.SetBrightness(brightness, true) {
			return true
		} else {
			panic("Lights failed to turn on")
			return false
		}
	case 3: //set color
		if l.SetColor(color, true) {
			return true
		} else {
			panic("Color not set")
			return false
		}
	case 4: //Adjust Brightness over time
		if l.SetBrightness(brightness, true) {
			return true
		} else {
			panic("Brightness not set")
			return false
		}
	}
	return false
}

func SecuritySystem(OID int, s *security.Alarm) bool {
	switch OID {
	case 5: //arm system
		if s.Arm() {
			return true
		} else {
			panic("System not Armed")
			return false
		}
	case 6: //disarm system
		if s.Disarm() {
			return true
		} else {
			panic("System unable to be disarmed")
			return false
		}
	case 7: //Trigger Alarm
		if s.Trigger() {
			return true
		} else {
			panic("Alarm cannot sound")
			return false
		}
	}
	return false
}
func HVAC(OID int, h *hvac.HVAC, temperature int, fanspeed int, mode string) bool {
	switch OID {
	case 8: //set temperature
		if h.SetTemperature(temperature) {
			return true
		} else {
			panic("Temperature failure")
			return false
		}
	case 9: //set fanspeed
		if h.SetFanSpeed(fanspeed) {
			return true
		} else {
			return false
		}
	case 10: //set Mode
		if h.SetMode(mode) {
			return true
		} else {
			return false
		}
	}
	return false
}

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
//	//if (Lighting->networkAdapter.go->Turn On()) -> (if returns true)
//	//return response to front end ->encourage message
//	//if Requested Status = off
//	//if (Lighting -> networkAdapter.go-> Turn Off()) -> if returns true
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
