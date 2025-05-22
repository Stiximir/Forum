package forum

import (
	"fmt"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3" // Nécessaire pour sqlite3
)

func Error(r error) {
	if r != nil {
		fmt.Println("Cette erreur est survenue", r)
	}
}

type Post struct {
	Title   string
	Content string
	Comment []Comment
	Id      string
	Pseudo  string
	DateD   string
	DateH   string
}

type Comment struct {
	Pseudo  string
	content string
}

type Filter struct {
	Cat string
}

type HomeData struct {
	Post    []Post
	Content Filter
	Date    string
}

func Home(w http.ResponseWriter, r *http.Request, templatePath string) {

	var Postlist []Post

	var HomeData HomeData

	DB, err := OpenDB()
	Error(err)
	defer DB.Close()

	if r.Method == http.MethodPost {

		r.ParseForm()

		category := r.FormValue("cat")

		HomeData.Content.Cat = category

		if category == "" {
			rows, err := DB.Query("SELECT posts.title,users.username, posts.created_at, posts.id FROM posts JOIN users ON posts.user_id = users.id ORDER BY posts.created_at DESC")
			Error(err)
			defer rows.Close()

			for rows.Next() {
				var p Post

				err := rows.Scan(&p.Title, &p.Pseudo, &HomeData.Date, &p.Id)
				Error(err)

				t, err := time.Parse(time.RFC3339, HomeData.Date)
				if err != nil {
					panic(err)
				}

				p.DateD = t.Format("2006/01/02")
				p.DateH = t.Format("15:04")

				Postlist = append(Postlist, p)
			}

		} else {

			// changer requete poura fficher le pseudo
			rows, err := DB.Query("SELECT posts.title ,users.username, posts.created_at, posts.id FROM posts JOIN users ON posts.user_id = users.id WHERE category_id = ?  ORDER BY posts.created_at DESC", category)
			Error(err)
			defer rows.Close()

			for rows.Next() {
				var p Post

				err := rows.Scan(&p.Title, &p.Pseudo, &HomeData.Date, &p.Id)
				Error(err)

				t, err := time.Parse(time.RFC3339, HomeData.Date)
				if err != nil {
					panic(err)
				}

				p.DateD = t.Format("2006/01/02")
				p.DateH = t.Format("15:04")
				Postlist = append(Postlist, p)
			}
		}

		HomeData.Post = Postlist

		RenderTemplate(w, "home", HomeData, templatePath)
		return
	}

	rows, err := DB.Query("SELECT posts.title ,users.username, posts.created_at, posts.id FROM posts JOIN users ON posts.user_id = users.id ORDER BY posts.created_at DESC")
	Error(err)
	defer rows.Close()

	for rows.Next() {
		var p Post

		err := rows.Scan(&p.Title, &p.Pseudo, &HomeData.Date, &p.Id)
		Error(err)

		t, err := time.Parse(time.RFC3339, HomeData.Date)
		if err != nil {
			panic(err)
		}

		p.DateD = t.Format("2006/01/02")
		p.DateH = t.Format("15:04")

		Postlist = append(Postlist, p)

	}

	HomeData.Post = Postlist

	RenderTemplate(w, "home", HomeData, templatePath)
}
