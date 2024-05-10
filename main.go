package main

import (
	"log"
	"net/http"
)

// Define a home handler function which returns a  byte slice
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func main() {
	// Register the home handler function for the "/" URL pattern
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
