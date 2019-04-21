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

	log.Fatal(http.ListenAndServe(":8080", nil))
}