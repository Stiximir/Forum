package forum

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	ID       int
	Username string
	Email    string
	Password string
}

var UserP int = 1

func GetRecentSearches(db *sql.DB) []User {
	var searches []User
	row, err := db.Query("SELECT * FROM users WHERE id = ?", UserP)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		item := User{}
		err := row.Scan(&item.ID, &item.Username, &item.Email, &item.Password)
		if err != nil {
			log.Fatal(err)
		}
		searches = append(searches, item)
	}
	return searches
}

func DBprofile(db *sql.DB, nil error) {
	searches := GetRecentSearches(db)
	fmt.Printf("ID\tUsername\tEmail\n")
	for _, item := range searches {
		fmt.Printf("%d\t%s\t%s\n", item.ID, item.Username, item.Email)
		data.ID = item.ID
		data.Username = item.Username
		data.Email = item.Email
	}
}

type Data struct {
	ID       int
	Username string
	Email    string
}

var data = Data{}

func Profile(w http.ResponseWriter, r *http.Request, templatePath string) {

	db, err := OpenDB()
	Error(err)
	defer db.Close()
	//DBprofile(db, nil)

	RenderTemplate(w, "profile", data, templatePath)
}
