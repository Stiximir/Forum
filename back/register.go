package forum

import (
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil

}

func Register(w http.ResponseWriter, r *http.Request, templatePath string) {
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		pseudo := r.FormValue("username")
		password := r.FormValue("password")

		current_time := time.Now()

		DB, err := OpenDB()
		Error(err)
		defer DB.Close()

		rows, err := DB.Query("SELECT * FROM users WHERE username = ?", pseudo)
		if err != nil {
			log.Println("Erreur requête pseudo:", err)
			http.Error(w, "Erreur serveur", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		erows, err := DB.Query("SELECT * FROM users WHERE email = ?", email)
		Error(err)
		defer erows.Close()

		if !rows.Next() && !erows.Next() {

			hashedPassword, err := hashPassword(password)
			if err != nil {
				log.Println("Erreur lors du hashage du mot de passe", err)
				http.Error(w, "Erreur serveur", http.StatusInternalServerError)
				return
			}

			_, err = DB.Exec("INSERT INTO users(email, password, username, created_at, profile_picture) VALUES (?, ?, ?, ?, ?)", email, hashedPassword, pseudo, current_time, "/uploads/default.jpg")
			if err != nil {
				log.Println("Erreur insertion:", err)
				http.Error(w, "Erreur serveur", http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		} else {
			log.Println("Utilisateur ou email déjà existant.")
			http.Redirect(w, r, "/register", http.StatusSeeOther)
			return
		}
	}

	data := map[string]interface{}{
		"Title": "Connexion",
		"User":  "",
	}

	RenderTemplate(w, "register", data, templatePath)
}
