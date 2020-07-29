package models

import (
	"database/sql"
	"time"
)

type Post struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Summary     string    `json:"summary"`
	Body        string    `json:"body"`
	Published   bool      `json:"published"`
	DateCreated time.Time `json:"dateCreated"`
	DateUpdated time.Time `json:"dateUpdated"`
	AuthorId    int       `json:"authorId"`
}

func (db *DB) AllPostsShorten() ([]Post, error) {
	rows, err := db.Query(`SELECT id, title, summary, date_created 
												 FROM posts 
												 WHERE published = true
												 ORDER BY date_created DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := []Post{}
	for rows.Next() {
		post := Post{}
		err := rows.Scan(&post.ID, &post.Title, &post.Summary, &post.DateCreated)
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

func (db *DB) AddPost(post *Post) (int, error) {
	row := db.QueryRow(`INSERT INTO posts (title, summary, body, author_id, published) 
										 VALUES ($1, $2, $3, $4, $5)
										 RETURNING id`,
		post.Title, post.Summary, post.Body, post.AuthorId, post.Published)

	lastInsertId := 0
	err := row.Scan(&lastInsertId)
	if err != nil {
		return -1, err
	}

	return lastInsertId, nil
}

func (db *DB) GetPostData(postId int) (Post, error) {
	row := db.QueryRow("SELECT * FROM posts WHERE id = $1", postId)
	var post Post
	err := row.Scan(&post.ID, &post.Title, &post.Summary, &post.Body, &post.DateCreated, &post.DateUpdated, &post.AuthorId, &post.Published)
	switch {
	case err == sql.ErrNoRows:
		return post, nil
	case err != nil:
		return post, err
	}

	return post, nil
}

func (db *DB) GetUserPosts(userId int) ([]Post, error) {
	rows, err := db.Query(`SELECT id, title, summary, date_created
												 FROM posts 
												 WHERE author_id = $1
												 ORDER BY date_created DESC`, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := []Post{}
	for rows.Next() {
		post := Post{}
		err := rows.Scan(&post.ID, &post.Title, &post.Summary, &post.DateCreated)
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

func (db *DB) UpdatePostPublishedStatus(postId int, published bool) error {
	_, err := db.Exec("UPDATE posts SET published = $1 WHERE id = $2", published, postId)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) DeletePost(postId int) error {
	_, err := db.Exec("DELETE FROM posts WHERE id = $1", postId)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) UpdatePost(postId int, post *Post) error {
	_, err := db.Exec(`UPDATE posts SET title = $1, summary = $2, body = $3, date_updated = $4
										WHERE id = $5`,
		post.Title, post.Summary, post.Body, time.Now(), postId)
	if err != nil {
		return err
	}

	return nil
}
