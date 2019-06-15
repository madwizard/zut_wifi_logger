package main

import (
	"github.com/goji/httpauth"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
	stopGpsScanner := make(chan bool)

	// Gracefully close down
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go Scanner(stopScanner)
	go GpsScanner(stopGpsScanner)

	go func() {
		r := mux.NewRouter()

		r.HandleFunc("/", home)
		r.HandleFunc("/status", status)
		r.HandleFunc("/data", data)
		r.NotFoundHandler = http.HandlerFunc(NotFound)
		http.Handle("/", httpauth.SimpleBasicAuth("user", "pass")(r))

		log.Print("Starting server www")

		http.ListenAndServe(":8080", nil)
	}()
	sig := <- sigs
	log.Printf("main: Got signal: %+v", sig)
	stopScanner <- true
	stopGpsScanner <- true
	log.Printf("main: Shutting down")
}
