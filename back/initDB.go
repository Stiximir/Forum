package forum

import (
	"database/sql"
)

func OpenDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "../data_DB/forum.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}
