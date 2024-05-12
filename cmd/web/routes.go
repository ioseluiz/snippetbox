package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	// Register the home handler function for the "/" URL pattern
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	// Create a file server which serves files out of the "/ui/static" directory
	// Note that the path given to the http.Dir function is relative to the project
	// directory root
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Register the file server as the handler
	// for all the URL paths that start with "/static/". For matching paths
	// "/static" prefix before the request reaches the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux

}
