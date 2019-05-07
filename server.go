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

var Username = string("user")
var Password = string("pass")

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