package main

import
(
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", homePage)
	http.HandleFunc("/data", sendData)
	http.HandleFunc("/startScan", startScanning)

	log.Fatal(http.ListenAndServeTLS(":8080", "docs/cert.pem", "docs/key.pem", nil))
}