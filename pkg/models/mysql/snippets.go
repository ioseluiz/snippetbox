package mysql

import (
	"database/sql"
	"errors"

	"github.com/ioseluiz/snippetbox_3/pkg/models"
)

// Define a SnippetModel type which wraps a sql.DB connection pool
type SnippetModel struct {
	DB *sql.DB
}

// This will insert a new snippet into the database
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	// Write the SQL statement we want to execute. I've split it over
	// two lines for readability.
	stmt := `INSERT INTO snippets (title, content, created, expires)
         VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	// Use the Exec() method on the embedded connection pool to execute the statement.
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	// Use the last insertId() on the result object to get the ID of our
	// newly inserted record in the snippets table
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// The ID returned has the type int64, so we convert it to an int type
	// before returning
	return int(id), nil

}

// This will return a specific snippet based on it's id.
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	// Write the SQL statement that we want to execute
	stmt := `SELECT id, title, content, created, expires FROM snippets
	         WHERE expires > UTC_TIMESTAMP() AND id = ?`

	// Use the QueryRow method on the connection pool to execute
	// our SQL statement, passing the untrusted id variable as the
	// value for the placeholder parameter. This returns a pointer to sql.Row
	// object which holds the result from the database.
	row := m.DB.QueryRow(stmt, id)

	// Initialize a pointer to a new zeroed Snippet struct
	s := &models.Snippet{}

	// Use the row.Scan() to copy the values from each field in sql.Row
	// to the corresponding field in the Snippet struct.
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		// If the query returns no rows, the row.Scan() will return a
		// sql.ErrNoRows error. We use the errors.Is() function to check
		// for that error specifically, and return our own models.ErrNoRecord
		// error instead.
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}

	}
	// If everything went OK then return the Snippet object
	return s, nil
}

// This will return the 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
