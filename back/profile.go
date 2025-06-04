package forum

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type User struct {
	ID       int
	Username string
	Email    string
	Password string
}

type Likes struct {
	Likes int
}

type Posts struct {
	Posts int
}

var UserP string = "2"

func Delete(db *sql.DB, id string) (int64, error) {
	sql := `DELETE FROM users WHERE id = ?;`
	result, err := db.Exec(sql, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func GetRecentSearches(db *sql.DB) []User {
	var searches []User
	row, err := db.Query("SELECT id, username, email, password FROM users WHERE id = ?", UserP)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() {
		item := User{}
		err := row.Scan(&item.ID, &item.Username, &item.Email, &item.Password)
		if err != nil {
			log.Fatal(err)
		}
		searches = append(searches, item)
	}
	return searches
}

func GetLikes(db *sql.DB) []Likes {
	var searches []Likes
	row, err := db.Query("SELECT COUNT(*) FROM likes WHERE user_id = ?", UserP)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() {
		item := Likes{}
		err := row.Scan(&item.Likes)
		if err != nil {
			log.Fatal(err)
		}
		searches = append(searches, item)
	}
	return searches
}

func GetPosts(db *sql.DB) []Posts {
	var searches []Posts
	row, err := db.Query("SELECT COUNT(*) FROM posts WHERE user_id = ?", UserP)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() {
		item := Posts{}
		err := row.Scan(&item.Posts)
		if err != nil {
			log.Fatal(err)
		}
		searches = append(searches, item)
	}
	return searches
}

func DBprofile(db *sql.DB, nil error) {
	searches := GetRecentSearches(db)
	searchesLikes := GetLikes(db)
	searchesPosts := GetPosts(db)
	for _, item := range searches {
		data.ID = strconv.Itoa(item.ID)
		data.Username = item.Username
		data.Email = item.Email
	}
	for _, item := range searchesLikes {
		data.Likes = item.Likes
	}
	for _, item := range searchesPosts {
		data.Posts = item.Posts
	}
}

type Data struct {
	ID          string
	Username    string
	Email       string
	User        string
	Likes       int
	Posts       int
	UserP       string
	RealUser    string
	PostProfile []Post
}

var data = Data{}

func getPostProfile(userID string) []Post {
	DB, err := OpenDB()
	Error(err)
	defer DB.Close()

	rows, err := DB.Query("SELECT posts.title, users.username, posts.created_at, posts.id FROM posts JOIN users ON posts.user_id = users.id WHERE users.id = ? ORDER BY posts.created_at DESC", userID)
	Error(err)
	defer rows.Close()

	var Postlist []Post
	var dateStr string

	for rows.Next() {
		var p Post
		err := rows.Scan(&p.Title, &p.Pseudo, &dateStr, &p.Id)
		Error(err)

		t, err := time.Parse(time.RFC3339, dateStr)
		if err != nil {
			t = time.Now()
		}

		p.DateD = t.Format("2006/01/02")
		p.DateH = t.Format("15:04")

		p = getLastCom(p)
		Postlist = append(Postlist, p)
	}
	return Postlist
}

func Profile(w http.ResponseWriter, r *http.Request, templatePath string) {
	db, err := OpenDB()
	Error(err)
	defer db.Close()

	UserP = r.URL.Query().Get("userId")
	RealUser := GetCookie(r, "user").Cookie
	data.RealUser = RealUser
	data.UserP = UserP

	if r.Method == http.MethodPost {
		_, err = Delete(db, UserP)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	data.User = "1"
	DBprofile(db, nil)

	data.PostProfile = getPostProfile(UserP)

	RenderTemplate(w, "profile", data, templatePath)
}
