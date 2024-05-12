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

	// Initialize a new http.Server struct.
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	// Call the ListenAndServe() method on our new http.Server struct
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
