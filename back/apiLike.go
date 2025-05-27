package forum

import (
	"encoding/json"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func ApiLike(w http.ResponseWriter, r *http.Request, templatePath string) {
	db, err := OpenDB()
	Error(err)
	defer db.Close()

	userID := GetCookie(r, "user").Cookie

	if r.Method == http.MethodPost {
		action := r.FormValue("action")
    	postID := r.FormValue("postID")
		if action == "like" {
    		_, err = db.Exec("INSERT OR IGNORE INTO likes (post_id, user_id) VALUES (?, ?)", postID, userID)
    		Error(err)
		} else if action == "unlike" {
    		_, err = db.Exec("DELETE FROM likes WHERE post_id = ? AND user_id = ?", postID, userID)
    		Error(err)
		}
	}

	postID := r.FormValue("postID")
	var likeCount int
	err = db.QueryRow("SELECT COUNT(*) FROM likes WHERE post_id = ?", postID).Scan(&likeCount)
	Error(err)
	var hasLiked bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM likes WHERE post_id = ? AND user_id = ?)", postID, userID).Scan(&hasLiked)
	if err != nil {
		if err.Error() != "Nothing found" {
			Error(err)
		}
	}

	data := struct {
		LikeCount int  `json:"likeCount"`
		HasLiked  bool `json:"hasLiked"`
	}{
		LikeCount: likeCount,
		HasLiked:  hasLiked,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
