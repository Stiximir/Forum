package forum

import (
	"net/http"
	"time"
)

func DetailPost(w http.ResponseWriter, r *http.Request, templatePath string) {

	var Postlist []Post
	var p Post
	var data HomeData

	postId := r.URL.Query().Get("postId")
	data.User = GetCookie(r, "user").Cookie

	DB, err := OpenDB()
	Error(err)
	defer DB.Close()

	if r.Method == http.MethodPost {
		r.ParseForm()

		send := r.FormValue("message")

		_, err = DB.Exec("INSERT INTO comments(post_id,user_id,content) VALUES (?,?,?)", postId, data.User, send)
		Error(err)

	}

	//on récupère les info du post
	row, err := DB.Query("SELECT posts.title, posts.description, posts.created_at, users.username, posts.id FROM posts JOIN users ON posts.user_id = users.id WHERE posts.id = ?", postId)
	Error(err)
	defer row.Close()

	for row.Next() {
		err = row.Scan(&p.Title, &p.Content, &data.Date, &p.Pseudo, &p.Id)
		Error(err)

	}

	t, err := time.Parse(time.RFC3339, data.Date)
	if err != nil {
		panic(err)
	}

	p.DateD = t.Format("2006/01/02")
	p.DateH = t.Format("15:04")

	//on récupère tous les commentaire du post

	rows, err := DB.Query("SELECT comments.content, users.username FROM comments JOIN users ON comments.user_id = users.id WHERE comments.post_id = ?", postId)
	Error(err)
	defer rows.Close()

	for rows.Next() {
		var com Comment

		err = rows.Scan(&com.Content, &com.Pseudo)
		Error(err)

		p.Comment = append(p.Comment, com)
	}

	Postlist = append(Postlist, p)

	data.Post = Postlist

	RenderTemplate(w, "detailsPost", data, templatePath)

}
