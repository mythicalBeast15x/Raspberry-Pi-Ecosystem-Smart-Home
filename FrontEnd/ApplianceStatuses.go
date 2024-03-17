package main

//Change the above to package main for testing

import (
	message "CMPSC488SP24SecThursday/messaging"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Device struct {
	Name     string `json:"Name"`
	Category string `json:"Category"`
	Status   string `json:"status"`
}

//{ name: 'Appliances', category: 'Home Appliances', status: 'off' },

//if you wish to test this file uncomment the below main && FOR NOW comment out login.go's main function
//build/run this file only
//this will set up the server on localhost:8080
//go to localhost:8080 in your browser and navigate FrontEnd/Templates/dashboard.html (easiest to just click through links on browser)
//navigate to devices tab and click any status to switch on or off
//this will send the entire list of devices and use the devices Handler to send a list of devices to the server

// NOTE FROM CAIT:
// I have set both this file and the login.go file to the main package.  This is not ideal and will be changed in the future.
// I have also uncommented the below main func and commented out the main func in login.go.  Lastly, I changed the map to
// generate device IDs and return their statuses as per Love's request.  In the future, we will store this map in the Pi's
// structure (created on startup) and ensure that all Pi's have the same map via Zigbee message.
// If you wish to revert my changes, just set these files back to package FrontEnd, and comment out the main func in this file
// while uncommenting the main func in login.go.

func main() {
	// Create a file server handler to serve static files from the same directory
	fs := http.FileServer(http.Dir("."))

	// Register the file server handler with the root URL path "/"
	http.Handle("/", fs)

	// Register the devicesHandler function with the "/devices" URL path
	http.HandleFunc("/devices", devicesHandler)

	// Start the server
	log.Println("Server is starting and listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func devicesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse JSON request body
	var devices []Device
	if err := json.NewDecoder(r.Body).Decode(&devices); err != nil {

		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	deviceMap := structToMap(devices)
	fmt.Println(devices)
	// Respond with the devices received
	responseJSON, err := json.Marshal(devices)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	//this line is for debugging purposes
	fmt.Println(deviceMap)

	message.NewMessage("senderdemo1", "all", "general", "21", deviceMap, &message.OpenMessages{}, &message.MessageQueue{})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)

}

// helper function to convert struct into map for message structure
func structToMap(devices []Device) map[string]interface{} {
	deviceMap := make(map[string]interface{})
	deviceID := 0
	for _, device := range devices {
		deviceMap[strconv.Itoa(deviceID)] = device.Status
		deviceID++

		//deviceMap[deviceID] = map[string]interface{}{
		//	"Category": device.Category,
		//	"Status":   device.Status,
		//}
	}
	return deviceMap
}
