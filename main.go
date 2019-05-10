package main

import (
	"github.com/goji/httpauth"
	"github.com/gorilla/mux"
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

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", home)
	r.HandleFunc("/status", status)
	r.HandleFunc("/data", data)
	http.Handle("/", httpauth.SimpleBasicAuth("user", "pass")(r))

	var data wifiData

	WiFiParse("wlp4s0", &data)

	http.ListenAndServe(":8080", nil)

}