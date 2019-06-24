package main

import (
	"github.com/goji/httpauth"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type webdata struct {
	Timestamp string `json:"Timestamp"`
	ESSID string `json:"ESSID"` 		// ESSID
	MAC string	`json:"MAC"`			// Address
	Freq string `json:"freq"`			// Frequency
	SigLvl string `json:"siglvl"`		// SignalLevel
	Qual string `json:"qual"`			// Quality
	Enc string `json:"enc"`				// Encryption key
	Channel int `json:"channel"`		// Channel
	Mode string `json:"mode"`			// Mode
	IEEE string `json:"IEEE"`			// IEEE
	Bitrates string `json:"bitrates"`	// bitrates
	WPA string `json:"wpa"`				// WPA version
	Latitude string `json:"Latitude"`
	Longitude string `json:"Longitude"`
}

// HTTP handler functions

type ServerStatus struct {
	WiFiScanUp bool
	GPSScanUp bool
	Timestamp time.Time
}

func home(w http.ResponseWriter, r *http.Request) {
	data := ServerStatus{true, true, time.Now()}
	err := tmpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func status(w http.ResponseWriter, r *http.Request) {
	serverStatus := ServerStatus{
		WiFiScanUp: true,
		GPSScanUp: true,
	}

	err := tmpl.ExecuteTemplate(w, "status.html", serverStatus)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

// Handle data page
// In final version read data from database
func data(w http.ResponseWriter, r *http.Request) {

	ScannedData := readDB()
	err := tmpl.ExecuteTemplate(w, "data.html", ScannedData)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func rendermap(w http.ResponseWriter, r *http.Request) {

	ScannedData := readDB()
	err := tmpl.ExecuteTemplate(w, "map.html", ScannedData)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}


func NotFound(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "404.html", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func wwwServer() {
	r := mux.NewRouter()

	r.HandleFunc("/", home)
	r.HandleFunc("/status", status)
	r.HandleFunc("/data", data)
	r.HandleFunc("/map", rendermap)
	r.NotFoundHandler = http.HandlerFunc(NotFound)
	http.Handle("/", httpauth.SimpleBasicAuth("user", "pass")(r))

	log.Print("Starting server www")

	http.ListenAndServe(":8080", nil)
}