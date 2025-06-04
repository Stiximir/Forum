package forum

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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

func UpdateProfilePicture(db *sql.DB, id string, path string) error {
	sql := `UPDATE users SET profile_picture = ? WHERE id = ?;`
	_, err := db.Exec(sql, path, id)
	return err
}

func ModificationProfil(w http.ResponseWriter, r *http.Request, templatePath string) {
	db, err := OpenDB()
	Error(err)
	defer db.Close()
	UserP = r.URL.Query().Get("userId")

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

		err = r.ParseMultipartForm(10 << 20)
		if err != nil {
			log.Println("Erreur form multipart:", err)
		} else {
			file, _, err := r.FormFile("profile_picture")
			if err == nil {
				defer file.Close()

				filePath := fmt.Sprintf("../uploads/%s.jpg", UserP)
				dst, err := os.Create(filePath)
				if err != nil {
					log.Println("Erreur création fichier:", err)
				} else {
					defer dst.Close()
					_, err = io.Copy(dst, file)
					if err != nil {
						log.Println("Erreur copie fichier:", err)
					} else {
						dbPath := fmt.Sprintf("/uploads/%s.jpg", UserP)
						err = UpdateProfilePicture(db, UserP, dbPath)
						if err != nil {
							log.Println("Erreur update DB image:", err)
						}
					}
				}
			}
		}
	}

	DBprofile(db, nil)
	RenderTemplate(w, "modificationProfil", data, templatePath)
}
