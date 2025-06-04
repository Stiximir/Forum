package forum

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func RenderTemplate(w http.ResponseWriter, page string, data any, templatePath string) {
	layout := filepath.Join(templatePath, "layout.html")
	navbar := filepath.Join(templatePath, "navbar.html")
	footer := filepath.Join(templatePath, "footer.html")
	content := filepath.Join(templatePath, page+".html")

	tmpl, err := template.ParseFiles(layout, navbar, footer, content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Error RenderTemplate:", err)
	}
}
