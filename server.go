package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"time"
)

var tmpl *template.Template

func init(){

	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	wwwServer()
}

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
	r.HandleFunc("/data", data)
	r.HandleFunc("/map", rendermap)
	r.NotFoundHandler = http.HandlerFunc(NotFound)
	http.Handle("/", r)

	log.Print("Starting server www")

	http.ListenAndServe(":8080", nil)
}