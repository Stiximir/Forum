package forum

import (
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3" // Nécessaire pour sqlite3
)

func Error(r error) {
	if r != nil {
		fmt.Println("Cette erreur est survenue", r)
	}
}

type Post struct {
	Title   string
	Content string
}

func Home(w http.ResponseWriter, r *http.Request, templatePath string) {

	DB, err := OpenDB()
	Error(err)
	defer DB.Close()

	rows, err := DB.Query("SELECT title, description FROM posts ORDER BY id DESC")
	Error(err)
	defer rows.Close()

	var Postlist []Post

	for rows.Next() {
		var p Post

		err := rows.Scan(&p.Title, &p.Content)
		Error(err)

		Postlist = append(Postlist, p)

	}

	RenderTemplate(w, "home", Postlist, templatePath)
}
