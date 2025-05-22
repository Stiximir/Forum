package forum

import (
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

func comparehash(mdphash string, mdp string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(mdphash), []byte(mdp))
	Error(err)
	return true
}

func Login(w http.ResponseWriter, r *http.Request, templatePath string) {

	if r.Method == http.MethodPost {
		id := r.FormValue("identifiant")
		password := r.FormValue("password")

		DB, err := OpenDB()
		Error(err)
		defer DB.Close()

		userId, err := DB.Query("SELECT username FROM users WHERE username = ?", id)
		Error(err)
		defer userId.Close()

		email, err := DB.Query("SELECT email FROM users WHERE email = ?", id)
		Error(err)
		defer email.Close()

		var paswpseudo string
		err = DB.QueryRow("SELECT password FROM users WHERE username = ?", id).Scan((&paswpseudo))
		Error(err)

		var paswEmail string
		err = DB.QueryRow("SELECT password FROM users WHERE email = ? ", id).Scan((&paswEmail))
		Error(err)

		// (comparehash(paswpseudo, password) || comparehash(paswEmail, password))  remplace la vérification du mdp pas sa quand le hase sera fait

		if (userId.Next() || email.Next()) && (comparehash(paswpseudo, password) || comparehash(paswEmail, password)) {

			var userEmail string
			err = DB.QueryRow("SELECT id FROM users WHERE email = ? ", id).Scan((&userEmail))
			Error(err)

			var userUsername string
			err = DB.QueryRow("SELECT id FROM users WHERE username = ? ", id).Scan((&userUsername))
			Error(err)

			fmt.Println(userEmail)
			fmt.Println(userUsername)

			if userEmail == "" {
				SetCookie(w, "user", userUsername)
			} else {
				SetCookie(w, "user", userEmail)
			}

			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	data := map[string]interface{}{
		"Title": "Connexion",
	}
	RenderTemplate(w, "login", data, templatePath)

}
