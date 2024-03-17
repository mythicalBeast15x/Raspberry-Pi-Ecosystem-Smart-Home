package FrontEnd

import (
	"html/template"
	"net/http"
)

// Hardcoded credentials
const (
	hardcodedUsername = "user"
	hardcodedPassword = "password"
)

// Handler for the login page
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Render the login template without any error message
		tmpl := template.Must(template.ParseFiles("templates/login.html"))
		tmpl.Execute(w, nil)
	} else if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Check if the provided credentials match the hardcoded values
		if username == hardcodedUsername && password == hardcodedPassword {
			// Successful login, redirect to the dashboard
			http.Redirect(w, r, "/dashboard?username="+username, http.StatusSeeOther)
		} else {
			// Invalid credentials, render login template with error message
			data := map[string]interface{}{
				"Error": true,
			}
			tmpl := template.Must(template.ParseFiles("templates/login.html"))
			tmpl.Execute(w, data)
		}
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
	// Clear any session or authentication tokens here
	// For example, if using session, you may clear session data
	// If using tokens, you may revoke the token

	// Redirect the user to the login page after logout
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func main() {
	// Register handlers
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/dashboard", dashboardHandler)
	http.HandleFunc("/logout", logoutHandler)

	// Serve static files if any
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Start the server
	http.ListenAndServe(":8080", nil)
}
