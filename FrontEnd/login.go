package main

import (
	"html/template"
	"net/http"
)

// User represents a user in the system
type User struct {
	Username string
	Password string
}

// Hardcoded users (for demonstration purposes)
var users = []User{
	{Username: "user1", Password: "password1"},
	{Username: "deep", Password: "Deeppatel"},
}

// Handler for the login page
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Render the login template without any error message
		tmpl := template.Must(template.ParseFiles("templates/login.html"))
		tmpl.Execute(w, nil)
	} else if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Check if the provided credentials match any user
		for _, user := range users {
			if username == user.Username && password == user.Password {
				// Successful login, redirect to the dashboard
				http.Redirect(w, r, "/dashboard?username="+username, http.StatusSeeOther)
				return
			}
		}

		// Invalid credentials, render login template with error message
		data := map[string]interface{}{
			"Error": true,
		}
		tmpl := template.Must(template.ParseFiles("templates/login.html"))
		tmpl.Execute(w, data)
	}
}

// Handler for the dashboard
func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the user is authenticated
	username := r.FormValue("username")
	if username == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Render the dashboard template
	tmpl := template.Must(template.ParseFiles("templates/dashboard.html"))
	tmpl.Execute(w, username)
}

// Handler for logout
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	// Redirect the user to the login page after logout
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// Handler for the forgot password page
func forgotPasswordHandler(w http.ResponseWriter, r *http.Request) {
	// Render the forgot password template
	tmpl := template.Must(template.ParseFiles("templates/forgotpassword.html"))
	tmpl.Execute(w, nil)
}

// Handler for the signup page
func signupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Render the signup template for GET requests
		tmpl := template.Must(template.ParseFiles("templates/signup.html"))
		tmpl.Execute(w, nil)
	} else if r.Method == "POST" {
		// Process the signup form submission for POST requests
		// Your signup logic goes here
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	}
}

// Handler for 404 page
func pageNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	tmpl := template.Must(template.ParseFiles("templates/static/404.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		// Handle error if template execution fails
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	// Register handlers
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/dashboard", dashboardHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/forgotpassword.html", forgotPasswordHandler)
	http.HandleFunc("/signup.html", signupHandler)
	http.HandleFunc("/", pageNotFoundHandler)

	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("templates/static"))))
	http.Handle("/terms/", http.StripPrefix("/terms/", http.FileServer(http.Dir("static"))))
	http.Handle("/privacy/", http.StripPrefix("/privacy/", http.FileServer(http.Dir("static"))))

	// Start the server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
