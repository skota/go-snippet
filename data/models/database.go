package models

import (
	"database/sql"
)

/*
  this is a comment
*/
type Database struct {
	*sql.DB
}

/*
	this is a comment
*/
func (db *Database) GetSnippet(id int) (*Snippet, error) {
	// if id == 123 {
	// 	snippet := &Snippet{
	// 		Id:      id,
	// 		Title:   "Example title",
	// 		Content: "Example content",
	// 		Created: time.Now(),
	// 		Expires: time.Now(),
	// 	}
	// 	return snippet, nil
	// }
	stmt := `SELECT id, title, content, created, expires
	FROM SNIPPETS WHERE id =?`

	row := db.QueryRow(stmt, id)

	s := &Snippet{}

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return s, nil
}
