package main

import (
	"log"
	"net/http"
	"time"
)


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

	err := tmpl.ExecuteTemplate(w, "data.html", ScannedData)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
	log.Print(ScannedData)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "404.html", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func Scanner(stop chan bool) {
	stopscanner := false
	log.Println("Scanner starting")
	for {
		WIFI, err := setWiFiInterface("config")
		if err != nil {
			log.Fatal("Can't read config file!")
		}
		WiFiParse(WIFI, &ScannedData)
		log.Println(ScannedData.MAC, ScannedData.ESSID)
		writeDB(ScannedData)
		stopscanner = <- stop
		if stopscanner == true {
			log.Println("Scanner stopping")
			break
		}
	}
}