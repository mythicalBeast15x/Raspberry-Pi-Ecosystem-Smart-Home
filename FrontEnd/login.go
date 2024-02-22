package FrontEnd

import (
    "html/template"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        t, _ := template.ParseFiles("login.html")
        t.Execute(w, nil)
    })

    http.HandleFunc("/login", loginHandler)
    http.HandleFunc("/forgot-password", forgotPasswordHandler)
    http.HandleFunc("/forgot-username", forgotUsernameHandler)
    http.HandleFunc("/signup", signupHandler)

    http.ListenAndServe(":8080", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
}

func forgotPasswordHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Forgot Password?"))
}

func forgotUsernameHandler(w http.ResponseWriter, r *http.Request) {
   
    w.Write([]byte("Forgot Username?"))
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Sign Up"))
}
