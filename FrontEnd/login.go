package main

import (
    "html/template"
    "net/http"
    "fmt"
)

func main() {
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/login", loginHandler)
    fmt.Println("Server starting on port :8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Printf("Failed to start server: %v\n", err)
    }
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("templates/login.html")
    t.Execute(w, nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        r.ParseForm()
        username := r.FormValue("username")
        password := r.FormValue("password")
     
        fmt.Printf("Login attempt with Username: %s Password: %s\n", username, password)

  
        http.Redirect(w, r, "/", http.StatusSeeOther)
    } else {
       
        http.Redirect(w, r, "/", http.StatusSeeOther)
    }
}

