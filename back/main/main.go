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
	templatePath, err := filepath.Abs(filepath.Join("..", "..", "template/html"))
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	mux.Handle("/html/", http.StripPrefix("/html", http.FileServer(http.Dir(filepath.Join(templatePath, "html")))))
	mux.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir(filepath.Join(templatePath, "css")))))
	mux.Handle("/picture/", http.StripPrefix("/picture", http.FileServer(http.Dir(filepath.Join(templatePath, "picture")))))
	mux.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("../uploads"))))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		forum.Home(w, r, templatePath)
	})
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		forum.Login(w, r, templatePath)
	})
	mux.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		forum.Register(w, r, templatePath)
	})
	mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		forum.Profile(w, r, templatePath)
	})
	mux.HandleFunc("/modificationProfil", func(w http.ResponseWriter, r *http.Request) {
		forum.ModificationProfil(w, r, templatePath)
	})
	mux.HandleFunc("/creatPost", func(w http.ResponseWriter, r *http.Request) {
		forum.CreatPost(w, r, templatePath)
	})
	mux.HandleFunc("/detailPost", func(w http.ResponseWriter, r *http.Request) {
		forum.DetailPost(w, r, templatePath)
	})

	mux.HandleFunc("/apiLike", func(w http.ResponseWriter, r *http.Request) {
		forum.ApiLike(w, r, templatePath)
	})

	//updatePost
	mux.HandleFunc("/updatePost", func(w http.ResponseWriter, r *http.Request) {
		forum.UpdatePost(w, r, templatePath)
	})

	//updateComment
	mux.HandleFunc("/updateComment", func(w http.ResponseWriter, r *http.Request) {
		forum.UpdateComment(w, r, templatePath)
	})

	mux.HandleFunc("/deleteCom", func(w http.ResponseWriter, r *http.Request) {
		forum.DeleteCom(w, r, templatePath)
	})

	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
