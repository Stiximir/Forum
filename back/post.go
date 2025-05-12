package forum

import (
	"net/http"
	"path/filepath"
	"text/template"
)

func requestPost(Title string, content string) {

	DB, err := OpenDB()
	Error(err)
	defer DB.Close()
}

func CreatPost(w http.ResponseWriter, r *http.Request, templatePath string) {

	tmplPath := filepath.Join(templatePath, "html", "creatPost.html")

	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}
