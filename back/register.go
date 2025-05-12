package forum

import (
	"net/http"
	"path/filepath"
	"text/template"
)

func Register(w http.ResponseWriter, r *http.Request, templatePath string) {

	tmplPath := filepath.Join(templatePath, "html", "register.html")

	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)

}
