package models

import (
	"time"
)

type Post struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Summary     string    `json:"summary"`
	Body        string    `json:"body"`
	DateCreated time.Time `json:"dateCreated"`
	DateUpdated time.Time `json:"dateUpdated"`
	AuthorId    int       `json:"authorId"`
}

func (db *DB) AllPostsShorten() ([]Post, error) {
	rows, err := db.Query("SELECT id, title, summary FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := []Post{}
	for rows.Next() {
		post := Post{}
		err := rows.Scan(&post.ID, &post.Title, &post.Summary)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (db *DB) AddPost(post *Post) error {
	_, err := db.Exec("INSERT INTO posts (title, summary, body, author_id) VALUES ($1, $2, $3, $4)",
		post.Title, post.Summary, post.Body, post.AuthorId)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) GetPostData(postId int) (Post, error) {
	row := db.QueryRow("SELECT * FROM posts WHERE id = $1", postId)
	var post Post
	err := row.Scan(&post.ID, &post.Title, &post.Summary, &post.Body, &post.DateCreated, &post.DateUpdated, &post.AuthorId)
	if err != nil {
		return post, err
	}
	return post, nil
}
