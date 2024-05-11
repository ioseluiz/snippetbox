package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// Define a new comnnad line flag with the name 'addr'
	addr := flag.String("addr", ":4000", "HTTP network address")
	//Parse the command line flag with flag.Parse() function
	flag.Parse()
	// Register the home handler function for the "/" URL pattern
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// Create a file server which serves files out of the "/ui/static" directory
	// Note that the path given to the http.Dir function is relative to the project
	// directory root
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	// Register the file server as the handler
	// for all the URL paths that start with "/static/". For matching paths
	// "/static" prefix before the request reaches the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
