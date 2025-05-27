package forum

import (
	"database/sql"
	"fmt"
	"net/http"
)

func Update(db *sql.DB, id string, username string) (int64, error) {
	sql := `UPDATE users SET username = ? WHERE id = ?;`
	result, err := db.Exec(sql, username, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func UpdateEmail(db *sql.DB, id string, email string) (int64, error) {
	sql := `UPDATE users SET email = ? WHERE id = ?;`
	result, err := db.Exec(sql, email, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func ModificationProfil(w http.ResponseWriter, r *http.Request, templatePath string) {

	db, err := OpenDB()
	Error(err)
	defer db.Close()

	if r.Method == http.MethodPost {
		username := r.FormValue("pseudo")
		if username != "" {
			_, err = Update(db, UserP, username)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		email := r.FormValue("email")
		if email != "" {
			_, err = UpdateEmail(db, UserP, email)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}

	DBprofile(db, nil)

	RenderTemplate(w, "modificationProfil", data, templatePath)
}
