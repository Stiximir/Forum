package forum

import (
	"database/sql"
	"fmt"
)

func OpenDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "../data_DB/forum.db")
	if err != nil {
		return nil, err
	}

	fmt.Println("Db ouverte")
	return db, nil
}
