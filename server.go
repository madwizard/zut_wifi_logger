package main

import (
	"html/template"
	"net/http"
	"time"
)

// HTTP handler functions

type ServerStatus struct {
	WiFiScanUp bool
	GPSScanUp bool
	Timestamp time.Time
}

type scannedData struct {
	Scanned string
}

func home(w http.ResponseWriter, r *http.Request) {
	data := ServerStatus{true, true, time.Now()}
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, data)

}

func status(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/status.html"))
	serverStatus := ServerStatus{
		WiFiScanUp: true,
		GPSScanUp: true,
	}

	tmpl.Execute(w, serverStatus)
}

func data(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/data.html"))

	scannedData := scannedData {
		Scanned: "wlp4s0    Scan completed :\n	Cell 01 - Address: 88:AD:43:F9:6A:BC",
	}
	tmpl.Execute(w, scannedData)
}