package forum

import (
	"encoding/json"
	"net/http"
)

func UpdateComment(w http.ResponseWriter, r *http.Request, templatePath string) {
	var c test
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, "Erreur de décodage JSON", http.StatusBadRequest)
		return
	}

	DB, err := OpenDB()
	Error(err)
	defer DB.Close()

	// Mise à jour du commentaire dans la base de données
	_, err = DB.Exec("UPDATE comments SET content = ? WHERE id = ?", c.Text, c.Id)
	if err != nil {
		http.Error(w, "Erreur lors de la mise à jour du commentaire", http.StatusInternalServerError)
		return
	}
}
