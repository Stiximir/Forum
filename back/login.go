package forum

import (
	"net/http"
	"path/filepath"
	"text/template"
)

func Login(w http.ResponseWriter, r *http.Request, templatePath string) {

	tmplPath := filepath.Join(templatePath, "html", "login.html")

	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)

}
