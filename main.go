package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Define a home handler function which returns a  byte slice
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

// Add a showSnippet handler function
func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// Add a createSnippet handler function
func createSnippet(w http.ResponseWriter, r *http.Request) {
	// Use r.Method to check wether the request is using POST or not.
	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
		return
	}

	w.Write([]byte("Create a new Snippet"))
}

func main() {
	// Register the home handler function for the "/" URL pattern
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
