package main

import (
	"errors"
	"fmt"

	// "html/template"
	"net/http"
	"strconv"

	"github.com/ioseluiz/snippetbox_3/pkg/models"
)

// Define a home handler function which returns a  byte slice
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w) // Use the notFound() helper
		return
	}

	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
	}

	for _, snippet := range s {
		fmt.Fprintf(w, "%v\n", snippet)
	}

	// Initialize a slice containing the paths to the two files. Note that the
	// home.page.tmpl file must be the "first" file in the slice.
	// files := []string{
	// 	"./ui/html/home.page.tmpl",
	// 	"./ui/html/base.layout.tmpl",
	// 	"./ui/html/footer.partial.tmpl",
	// }

	// // Use the template.ParseFiles function to read the template file into a
	// // a template set. If there's an error, we log the detailed error message and use the
	// // http.Error() function to send a generic 500 Internal Server Error
	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	app.serverError(w, err) // Use the serverError() helper
	// 	return
	// }

	// //We then use the Execute() method on the template set to write the template
	// // content as the response body. The last parameter to Execute() represents any
	// // dynamic data that we want to pass in.
	// err = ts.Execute(w, nil)
	// if err != nil {
	// 	app.serverError(w, err) // Use the serverError() helper
	// }
}

// Add a showSnippet handler function
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w) // Use the not found() helper
		return
	}

	// Use the SnippetModel object's Get method to retrieve the data for a
	// specific record based on it's ID. If no matching record is found,
	// return a 404 Not Found response.
	s, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	// Write the snippet data as a plain-text HTTP response body.
	fmt.Fprintf(w, "%v", s)
}

// Add a createSnippet handler function
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	// Use r.Method to check wether the request is using POST or not.
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed) // Use the clientError() helper
		return
	}

	// Create some variables holding dummy data. We'll remove these later on
	// during the build.
	title := "0 snail"
	content := "0 snail\nClimb Mount Fuji, \nBut slowly, slowly!\n\n- Kobayashi Issa"
	expires := "7"

	// Pass the data to the SnippetModel.Insert() method, receiving the ID
	// of the new record back.
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Redirect the user to the relevant page for the snippet...
	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}
