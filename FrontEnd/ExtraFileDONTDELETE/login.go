package FrontEnd

import (
	"html/template"
	"net/http"
)

func startLogin() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("FrontEnd/login.html")
		err := t.Execute(w, nil)
		if err != nil {
			return
		}
	})

	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/forgot-password", forgotPasswordHandler)
	http.HandleFunc("/forgot-username", forgotUsernameHandler)
	http.HandleFunc("/signup", signupHandler)

	err := http.ListenAndServeTLS(":8080", "etc/server.crt", "etc/server.key", nil)
	if err != nil {
		return
	}

}

func loginHandler(http.ResponseWriter, *http.Request) {
}

func forgotPasswordHandler(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("Forgot Password?"))
	if err != nil {
		return
	}
}

func forgotUsernameHandler(w http.ResponseWriter, _ *http.Request) {

	_, err := w.Write([]byte("Forgot Username?"))
	if err != nil {
		return
	}
}

func signupHandler(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("Sign Up"))
	if err != nil {
		return
	}

}
