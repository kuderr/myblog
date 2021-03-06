package posts

import (
	"backend/config"
	"database/sql"
	"time"
)

type Post struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Summary     string    `json:"summary"`
	Image       string    `json:"img"`
	Color       string    `json:"color"`
	Body        string    `json:"body"`
	Published   bool      `json:"published"`
	DateCreated time.Time `json:"dateCreated"`
	DateUpdated time.Time `json:"dateUpdated"`
	AuthorId    int       `json:"authorId"`
	Views				int 			`json:"views"`
}

type PostShorten struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Summary     string    `json:"summary"`
	DateCreated time.Time `json:"dateCreated"`
	Image       string    `json:"img"`
	Color       string    `json:"color"`
	Views				int 			`json:"views"`
}

func allPostsShorten() ([]PostShorten, error) {
	rows, err := config.DB.Query(`SELECT id, title, summary, date_created, image, color
												 FROM posts 
												 WHERE published = true
												 ORDER BY date_created DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := []PostShorten{}
	for rows.Next() {
		post := PostShorten{}
		err := rows.Scan(&post.ID, &post.Title, &post.Summary, &post.DateCreated,
			&post.Image, &post.Color)
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

func addPost(post *Post) (int, error) {
	row := config.DB.QueryRow(`INSERT INTO posts (title, summary, author_id) 
										 VALUES ($1, $2, $3)
										 RETURNING id`,
		post.Title, post.Summary, post.AuthorId)

	lastInsertId := 0
	err := row.Scan(&lastInsertId)
	if err != nil {
		return -1, err
	}

	return lastInsertId, nil
}

func getPostData(postId int) (Post, error) {
	row := config.DB.QueryRow("SELECT * FROM posts WHERE id = $1", postId)
	var post Post
	err := row.Scan(&post.ID, &post.Title, &post.Summary, &post.Body,
		&post.DateCreated, &post.DateUpdated, &post.AuthorId, &post.Published,
		&post.Image, &post.Color, &post.Views)
	switch {
	case err == sql.ErrNoRows:
		return post, nil
	case err != nil:
		return post, err
	}

	return post, nil
}

func getUserPosts(userId int) ([]Post, error) {
	rows, err := config.DB.Query(`SELECT id, title, summary, date_created, image, color
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
		err := rows.Scan(&post.ID, &post.Title, &post.Summary, &post.DateCreated,
			&post.Image, &post.Color)
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

func updatePostPublishedStatus(postId int, published bool) error {
	_, err := config.DB.Exec("UPDATE posts SET published = $1 WHERE id = $2", published, postId)
	if err != nil {
		return err
	}

	return nil
}

func deletePost(postId int) error {
	_, err := config.DB.Exec("DELETE FROM posts WHERE id = $1", postId)
	if err != nil {
		return err
	}

	return nil
}

func updatePost(postId int, post *Post) error {
	_, err := config.DB.Exec(`UPDATE posts SET title = $1, summary = $2, body = $3, date_updated = $4,
										image = $5, color = $6
										WHERE id = $7`,
		post.Title, post.Summary, post.Body, time.Now(), post.Image, post.Color, postId)
	if err != nil {
		return err
	}

	return nil
}

func addView(postId int) error {
	_, err := config.DB.Exec(`UPDATE posts SET views = views + 1
														WHERE id = $1`, postId)
	if err != nil{
		return err
	}

	return nil
}