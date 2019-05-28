package main

import (
	"github.com/goji/httpauth"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"time"
)

var ScannedData wifiData

// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

// New returns an error that formats as the given text.
func New(text string) error {
	return &errorString{text}
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))

	initDB()
}

func main() {
	stopScanner := make(chan bool)

	r := mux.NewRouter()

	r.HandleFunc("/", home)
	r.HandleFunc("/status", status)
	r.HandleFunc("/data", data)
	r.NotFoundHandler = http.HandlerFunc(NotFound)
	http.Handle("/", httpauth.SimpleBasicAuth("user", "pass")(r))

	go Scanner(stopScanner)

	time.Sleep(5 * time.Second)
	stopScanner <- true

	http.ListenAndServe(":8080", nil)

}
