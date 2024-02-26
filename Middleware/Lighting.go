package main

import (
	"CMPSC488SP24SecThursday/lighting"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	//Commenting out until branches get merged because JWT is only in BE/DB Branch
	//"CMPSC488SP24SecThursday/JWT"
)

// Represents a user for testing purposes
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Hardcoded user credentials for demo
var users = map[string]string{
	"admin": "password",
}

// LightControlRequest represents the payload expected when a light control request is made
type LightControlRequest struct {
	LightID string `json:"light_id"` // ID of the light to control
	Action  string `json:"action"`   // Action to perform: "on" or "off"
	Token   string `json:"token"`    // JWT token for authentication
}

// LightControlHandler handles requests to control the lighting.
/*
Purpose: This function processes incoming HTTP POST requests to control lighting devices. It decodes the request body
into a `LightControlRequest` struct, validating the request data. If the request is malformed (cannot be decoded),
it responds with a "Bad request" error.

Params:
    w http.ResponseWriter - The response writer to send responses to the client.
    r *http.Request - The request object containing all the details of the HTTP request made by the client.
*/
func LightControlHandler(w http.ResponseWriter, r *http.Request) {
	var req LightControlRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// token validation from JWT package in Database/Backend
	// so need to merge branches to verify this works as intended
	_, err := JWT.VerifyJWT(req.Token)
	if err != nil {
		// If token is invalid, check if user's credentials are provided
		var user User
		err = json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Bad request: "+err.Error(), http.StatusBadRequest)
			return
		}

		// Validate user credentials (this is a simplified version without secure password handling)
		password, ok := users[user.Username]
		if !ok || password != user.Password {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		// Generate a new token for the user
		clientInfo := JWT.ClientInfo{Username: user.Username, Password: user.Password}
		newToken, err := JWT.GenerateJWT(clientInfo)
		if err != nil {
			http.Error(w, "Failed to generate new token: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Send the new token back to the user
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"new_token": newToken})
		return

	}

	// create light object to act upon
	light := lighting.NewLighting(req.LightID, false)

	// Control the light based on the action specified in the request.
	switch req.Action {
	case "on":
		light.TurnOn()
	case "off":
		light.TurnOff()
	default:
		http.Error(w, "Invalid action", http.StatusBadRequest)
		return
	}

	// Respond to the request indicating the action was successful
	json.NewEncoder(w).Encode(map[string]string{"result": "success"})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/control-light", LightControlHandler).Methods("POST") // New handler for light control

	// Start server with TLS
	log.Println("Server listening on port 8080 with TLS...")
	// logging to cert and key in etc folder
	err := http.ListenAndServeTLS(":8080", "etc/server.crt", "etc/server.key", r)
	if err != nil {
		log.Fatal("ListenAndServeTLS: ", err)
	}
}
