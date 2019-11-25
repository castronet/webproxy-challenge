package main

import (
	"net/http"
	"io"
	"log"
)

const (
    LISTEN_PORT = "8080"
)

// Main function
func main() {
	// Handle specific request.
	// Set the configuration before start the server.
	http.HandleFunc("/hello", handleRequest);

	// Create server and listen on desired port.
	err := http.ListenAndServe(":"+LISTEN_PORT, nil)
	log.Fatal(err)
}

// Handle incoming requests
func handleRequest(w http.ResponseWriter, req *http.Request) {
	// As requested this endopoint just has to answer hello world.
	io.WriteString(w, "{ \"hello\": \"world\" }")
}
