package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// Define an application struct to hold the application wide dependencies
// for the web application.
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	// Define a new comnnad line flag with the name 'addr'
	addr := flag.String("addr", ":4000", "HTTP network address")
	//Parse the command line flag with flag.Parse() function
	flag.Parse()
	// Use log.New() to create a logger for writing information messages.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// Create a logger for writing error messages.
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Initialize  a new instance of application containing the dependencies
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

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

	// Initialize a new http.Server struct.
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Starting server on %s", *addr)
	// Call the ListenAndServe() method on our new http.Server struct
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
