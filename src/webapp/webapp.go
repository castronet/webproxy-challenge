package main

import (
	"fmt"
	"net/http"
	"os"
	"log"
)

const (
//	DEBUG = false
	DEBUG = true
)

// Main function
func main() {

	// Create server and listen on desired port with TLS
	// The certificate is self signed. A good idea is to use Let's Encrypt certificates.
	err := http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil)
	log.Fatal(err)

	// Handle specific request
	http.HandleFunc("/hello", handleRequest);
}

// Handle incoming requests
func handleRequest(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "{ \"hello\": \"world\" }")
}
