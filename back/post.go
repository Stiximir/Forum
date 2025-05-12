package forum

import (
	"net/http"
	"path/filepath"
	"text/template"
	"time"
)

func requestPost(Title string, content string) {

	DB, err := OpenDB()
	Error(err)
	defer DB.Close()
}

func CreatPost(w http.ResponseWriter, r *http.Request, templatePath string) {
	if r.Method == http.MethodPost {

		//pour la catégorie prendre l'ID pluto que le text
		cat := r.FormValue("cat")
		title := r.FormValue("title")
		content := r.FormValue("content")
		user := GetCookie(r, "user")
		current_time := time.Now()

		if cat != "" && title != "" && content != "" && user.Cookie != "" {

			DB, err := OpenDB()
			Error(err)
			defer DB.Close()

			_, err = DB.Exec("INSERT INTO posts(user_id , title , description , created_at , category_id) VALUES (?, ?, ?, ?, ?)", user.Cookie, title, content, current_time, cat)
			Error(err)
		}

	}
	tmplPath := filepath.Join(templatePath, "html", "creatPost.html")

	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}
