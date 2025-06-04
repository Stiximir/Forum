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
	Title          string
	Content        string
	Comment        []Comment
	ProfilePicture string
	CreatorId      string
	Id             string
	Pseudo         string
	DateD          string
	DateH          string
	LikeCount      int
	HasLiked       bool
}

type Comment struct {
	Pseudo         string
	Content        string
	CreatorId      string
	Id             string
	ProfilePicture string
}

type Filter struct {
	Cat string
}

type HomeData struct {
	Post    []Post
	Content Filter
	Date    string
	User    string
}

func Home(w http.ResponseWriter, r *http.Request, templatePath string) {

	var Postlist []Post

	var HomeData HomeData

	HomeData.User = GetCookie(r, "user").Cookie

	if r.Method == http.MethodPost {

		r.ParseForm()

		category := r.FormValue("cat")

		HomeData.Content.Cat = category

		if category == "" {
			HomeData, Postlist = getPost(HomeData, Postlist)
		} else {
			HomeData, Postlist = getPostFilter(HomeData, Postlist, category)
		}

		HomeData.Post = Postlist

		RenderTemplate(w, "home", HomeData, templatePath)
		return
	}

	HomeData, Postlist = getPost(HomeData, Postlist)

	HomeData.Post = Postlist

	RenderTemplate(w, "home", HomeData, templatePath)
}

func getLastCom(p Post) Post {
	DB, err := OpenDB()
	Error(err)
	defer DB.Close()

	rows, err := DB.Query("SELECT comments.content, users.username FROM comments JOIN users ON comments.user_id = users.id WHERE comments.post_id = ? ORDER BY comments.created_at DESC LIMIT 1", p.Id)
	Error(err)
	defer rows.Close()

	for rows.Next() {
		var com Comment

		err = rows.Scan(&com.Content, &com.Pseudo)
		Error(err)

		p.Comment = append(p.Comment, com)
	}

	return p
}

func getPost(HomeData HomeData, Postlist []Post) (HomeData, []Post) {
	DB, err := OpenDB()
	Error(err)
	defer DB.Close()

	rows, err := DB.Query("SELECT posts.title ,users.username, posts.created_at, posts.id , users.profile_picture FROM posts JOIN users ON posts.user_id = users.id ORDER BY posts.created_at DESC")
	Error(err)
	defer rows.Close()

	for rows.Next() {
		var p Post

		err := rows.Scan(&p.Title, &p.Pseudo, &HomeData.Date, &p.Id, &p.ProfilePicture)
		Error(err)

		t, err := time.Parse(time.RFC3339, HomeData.Date)
		if err != nil {
			panic(err)
		}

		p.DateD = t.Format("2006/01/02")
		p.DateH = t.Format("15:04")

		p = getLastCom(p)

		err = DB.QueryRow("SELECT COUNT(*) FROM likes WHERE post_id = ?", p.Id).Scan(&p.LikeCount)
		Error(err)
		userID := HomeData.User
		if userID != "" {
			err = DB.QueryRow("SELECT EXISTS(SELECT 1 FROM likes WHERE post_id = ? AND user_id = ?)", p.Id, userID).Scan(&p.HasLiked)
			Error(err)
		}

		Postlist = append(Postlist, p)

	}
	return HomeData, Postlist
}

func getPostFilter(HomeData HomeData, Postlist []Post, category string) (HomeData, []Post) {
	DB, err := OpenDB()
	Error(err)
	defer DB.Close()

	rows, err := DB.Query("SELECT posts.title ,users.username, posts.created_at, posts.id , users.profile_picture FROM posts JOIN users ON posts.user_id = users.id WHERE category_id = ?  ORDER BY posts.created_at DESC", category)
	Error(err)
	defer rows.Close()

	for rows.Next() {
		var p Post

		err := rows.Scan(&p.Title, &p.Pseudo, &HomeData.Date, &p.Id, &p.ProfilePicture)
		Error(err)

		t, err := time.Parse(time.RFC3339, HomeData.Date)
		if err != nil {
			panic(err)
		}

		p.DateD = t.Format("2006/01/02")
		p.DateH = t.Format("15:04")

		p = getLastCom(p)

		err = DB.QueryRow("SELECT COUNT(*) FROM likes WHERE post_id = ?", p.Id).Scan(&p.LikeCount)
		Error(err)
		userID := HomeData.User
		if userID != "" {
			err = DB.QueryRow("SELECT EXISTS(SELECT 1 FROM likes WHERE post_id = ? AND user_id = ?)", p.Id, userID).Scan(&p.HasLiked)
			Error(err)
		}

		Postlist = append(Postlist, p)
	}
	return HomeData, Postlist
}
