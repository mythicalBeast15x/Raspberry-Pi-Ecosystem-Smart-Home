package FrontEnd

//Change the above to package main for testing

import (
	message "CMPSC488SP24SecThursday/messaging"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	_ "github.com/gorilla/websocket"
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

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { // Allow all connections for simplicity
			return true
		},
	}
	clients = make(map[*websocket.Conn]bool) // Keep track of connected WebSocket clients
)

/*
// ----UNCOMMENT FOR TESTING-----

	func main() {
		// Create a file server handler to serve static files from the same directory
		fs := http.FileServer(http.Dir("."))

		// Register the file server handler with the root URL path "/"
		http.Handle("/", fs)

		// Register the devicesHandler function with the "/devices" URL path
		http.HandleFunc("/devices", devicesHandler)
		http.HandleFunc("/ws", handleWebSocket)

		// Start the server
		log.Println("Server is starting and listening on port 8080...")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}
*/
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

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer conn.Close()
	clients[conn] = true

	for {
		var msg message.Message // Corrected to use the import alias 'message'
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("WebSocket read error: %v", err)
			delete(clients, conn)
			break
		}
		broadcastToDeviceClients(msg)
	}
}

func broadcastToDeviceClients(msg message.Message) { // Corrected to use the import alias 'message'
	for client := range clients {
		err := client.WriteJSON(msg)
		if err != nil {
			log.Printf("WebSocket write error: %v", err)
			client.Close()
			delete(clients, client)
		}
	}
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
