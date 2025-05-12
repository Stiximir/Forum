package forum

import (
	"database/sql"
	"fmt"
	"log")


func OpenDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "../data_DB/forum.db") // OpenDatabase function
	if err != nil {
		return nil, err
	}

	fmt.Println("Db ouverte")
	return db, nil
}
func CloseDatabase(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Db fermée")
}

//Users
func GetUser(db *sql.DB, username string) (string, error) {
	var result string
	query := "SELECT * FROM users WHERE username = ?"
	err := db.QueryRow(query, username).Scan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("user not found")
		}
		return "", err
	}
	return result, nil
}
func GetAllUsers(db *sql.DB) ([]string, error) {
	var users []string
	query := "SELECT username FROM users"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var username string
		if err := rows.Scan(&username); err != nil {
			return nil, err
		}
		users = append(users, username)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

//Posts
func GetAllPosts(db *sql.DB) ([]string, error) {
	var posts []string
	query := "SELECT * FROM posts"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var post string
		if err := rows.Scan(&post); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}
func GetPostsByUser(db *sql.DB, userId int) ([]string, error) {
	var posts []string
	query := "SELECT * FROM posts WHERE user_id = ?"
	rows, err := db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var post string
		if err := rows.Scan(&post); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}

//Comments
func GetAllComments(db *sql.DB, postId int) ([]string, error) {
	var comments []string
	query := "SELECT * FROM comments WHERE post_id = ?"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var comment string
		if err := rows.Scan(&comment); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return comments, nil
}
func GetCommentsByUser(db *sql.DB, userId int) ([]string, error) {
	var comments []string
	query := "SELECT * FROM comments WHERE user_id = ?"
	rows, err := db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var comment string
		if err := rows.Scan(&comment); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return comments, nil
}

//Categories
func GetAllCategories(db *sql.DB) ([]string, error) {
	var categories []string
	query := "SELECT * FROM categories"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var category string
		if err := rows.Scan(&category); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return categories, nil
}