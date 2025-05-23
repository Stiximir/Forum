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
	templatePath, err := filepath.Abs(filepath.Join("..","..", "template/html"))
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/html/", http.StripPrefix("/html", http.FileServer(http.Dir(filepath.Join(templatePath, "html")))))
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir(filepath.Join(templatePath, "css")))))
	http.Handle("/picture/", http.StripPrefix("/picture", http.FileServer(http.Dir(filepath.Join(templatePath, "picture")))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		forum.Home(w, r, templatePath)
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		forum.Login(w, r, templatePath)
	})

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		forum.Register(w, r, templatePath)
	})

	http.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		forum.Profile(w, r, templatePath)
	})

	http.HandleFunc("/modificationProfil", func(w http.ResponseWriter, r *http.Request) {
		forum.ModificationProfil(w, r, templatePath)
	})

	//création de post
	http.HandleFunc("/creatPost", func(w http.ResponseWriter, r *http.Request) {
		forum.CreatPost(w, r, templatePath)
	})

	//détail des posts
	http.HandleFunc("/DetailPost", func(w http.ResponseWriter, r *http.Request) {
		forum.DetailPost(w, r, templatePath)
	})

	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
