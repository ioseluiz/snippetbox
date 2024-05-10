package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// Define a home handler function which returns a  byte slice
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Initialize a slice containing the paths to the two files. Note that the
	// home.page.tmpl file must be the "first" file in the slice.
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}

	// Use the template.ParseFiles function to read the template file into a
	// a template set. If there's an error, we log the detailed error message and use the
	// http.Error() function to send a generic 500 Internal Server Error
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	//We then use the Execute() method on the template set to write the template
	// content as the response body. The last parameter to Execute() represents any
	// dynamic data that we want to pass in.
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
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
