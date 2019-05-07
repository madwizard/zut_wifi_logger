package main

import (
	"github.com/goji/httpauth"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", home)
	r.HandleFunc("/status", status)
	http.Handle("/", httpauth.SimpleBasicAuth("user", "pass")(r))
	// http.Handle("/status", httpauth.SimpleBasicAuth("user", "pass")(r))

	http.ListenAndServe(":8080", nil)

}