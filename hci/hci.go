package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type HomeData struct {
	Appliances  string
	Lights      string
	HVAC        string
	Security    string
	Usage       string
	Time        string
	DarkMode    bool
	Fan         string
	DoorLock    string
	VacuumRobot string
	MusicSystem string
}

var homeData HomeData

func main() {
	homeData = HomeData{
		"Connected",
		"On",
		"22°C",
		"Secure",
		"Low",
		time.Now().Format(time.Stamp),
		false, // Default to light mode
		"On",  // Fan status
		"Off", // Door Lock status
		"On",  // Vacuum Robot status
		"Off", // Music System status
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", homeHandler)

	fmt.Println(http.ListenAndServe(":8000", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		updateHomeData(r)
	}

	tmpl, err := template.ParseFiles("template/template.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, homeData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func updateHomeData(r *http.Request) {
	homeData.Appliances = r.FormValue("appliances")
	homeData.Lights = r.FormValue("lights")
	homeData.HVAC = r.FormValue("hvac")
	homeData.Security = r.FormValue("security")
	homeData.Usage = r.FormValue("usage")

	homeData.DarkMode = r.FormValue("darkMode") == "on"

	homeData.Fan = r.FormValue("fan")
	homeData.DoorLock = r.FormValue("doorLock")
	homeData.VacuumRobot = r.FormValue("vacuumRobot")
	homeData.MusicSystem = r.FormValue("musicSystem")

	homeData.Time = time.Now().Format(time.Stamp)
}
