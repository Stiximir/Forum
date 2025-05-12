package main

import (
	"fmt"
	"forum"
	"log"
	"net/http"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	templatePath, err := filepath.Abs(filepath.Join("..", "..", "template"))
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/html/", http.StripPrefix("/html", http.FileServer(http.Dir(filepath.Join(templatePath, "html")))))
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir(filepath.Join(templatePath, "css")))))
	http.Handle("/picture/", http.StripPrefix("/picture", http.FileServer(http.Dir(filepath.Join(templatePath, "picture")))))

	//page acceuil
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		forum.Home(w, r, templatePath)
	})

	//page login
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		forum.Login(w, r, templatePath)
	})

	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
