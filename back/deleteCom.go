package forum

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func DeleteCom(w http.ResponseWriter, r *http.Request, templatePath string) {
	var c test
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, "Erreur de décodage JSON", http.StatusBadRequest)
		return
	}
	fmt.Println(c.Id)

	DB, err := OpenDB()
	Error(err)
	defer DB.Close()

	// Mise à jour du post dans la base de données
	_, err = DB.Exec("DELETE FROM comments  WHERE id = ?", c.Id)
	if err != nil {
		http.Error(w, "Erreur lors de la mise à jour du post", http.StatusInternalServerError)
		return
	}

}
