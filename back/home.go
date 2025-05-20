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

func Home(w http.ResponseWriter, r *http.Request, templatePath string) {
	data := map[string]interface{}{
		"Title": "Accueil",
	}
	RenderTemplate(w, "home", data, templatePath)
}
