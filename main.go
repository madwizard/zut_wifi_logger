package main

import (
	"github.com/goji/httpauth"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

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
}

func main() {

	// This needs to be configurable
	// By arguments, config file or DB
	WIFI := "wlp0s20f3"

	r := mux.NewRouter()

	r.HandleFunc("/", home)
	r.HandleFunc("/status", status)
	r.HandleFunc("/data", data)
	r.NotFoundHandler = http.HandlerFunc(NotFound)
	http.Handle("/", httpauth.SimpleBasicAuth("user", "pass")(r))

	var data wifiData

	WiFiParse(WIFI, &data)

	http.ListenAndServe(":8080", nil)

}
