package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Hardcoded credentials for demonstration purposes
const (
	username = "admin"
	password = "password123"
)

// generateToken simulates token generation by using the current date and a renewal period of 30 days
func generateToken() string {
	currentDate := time.Now()
	renewalDate := currentDate.AddDate(0, 0, 30) // Adds 30 days to simulate token renewal
	token := fmt.Sprintf("%d%d", renewalDate.YearDay(), renewalDate.Hour())
	return token
}

// loginHandler processes login requests. It validates the provided username and password.
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	formUsername := r.FormValue("username")
	formPassword := r.FormValue("password")

	if formUsername != username || formPassword != password {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	token := generateToken()
	fmt.Fprintf(w, "Login Successful. Token: %s", token)
}

// LoggingMiddleware logs the HTTP method and URI of requests
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)

		log.Printf("Completed in %v", time.Since(start))
	})
}

func main() {
	mux := http.NewServeMux()

	// Wrap the loginHandler with the LoggingMiddleware for logging request details
	mux.Handle("/login", LoggingMiddleware(http.HandlerFunc(loginHandler)))

	log.Println("Server listening on :8080...")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
