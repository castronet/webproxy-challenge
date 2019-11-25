package main

import (
//	"fmt"
	"net/http"
	"io"
	"log"
)

// Main function
func main() {
	// Handle specific request
	// Set the configuration before start the server
	http.HandleFunc("/hello", handleRequest);

	// Create server and listen on desired port with TLS
	// The certificate is self signed. A good idea is to use Let's Encrypt certificates.
	err := http.ListenAndServeTLS(":8080", "server.crt", "server.key", nil)
	log.Fatal(err)
}

// Handle incoming requests
func handleRequest(w http.ResponseWriter, req *http.Request) {
	// As requested this endopoint just has to answer hello world
	io.WriteString(w, "{ \"hello\": \"world\" }")
}
