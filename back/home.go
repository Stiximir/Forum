package forum

import (
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"

	_ "github.com/mattn/go-sqlite3" // Nécessaire pour sqlite3
)

func Error(r error) {
	if r != nil {
		fmt.Println("Cette erreur est survenue", r)
	}
}

func Home(w http.ResponseWriter, r *http.Request, templatePath string) {

	// test base de donnée
	// var test string
	// DB, err := OpenDB()
	// Error(err)

	// err = DB.QueryRow("SELECT username FROM users WHERE id = ?", 1).Scan(&test)
	// Error(err)

	// fmt.Println(test)

	tmplPath := filepath.Join(templatePath, "html", "home.html")

	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)

}
