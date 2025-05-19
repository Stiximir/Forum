package forum

import (
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"

	_ "github.com/mattn/go-sqlite3" // Nécessaire pour sqlite3
)

type Post struct {
	Title   string
	Content string
}

func Error(r error) {
	if r != nil {
		fmt.Println("Cette erreur est survenue", r)
	}
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

	fmt.Println(Postlist)

	tmplPath := filepath.Join(templatePath, "html", "home.html")

	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, Postlist)

}
