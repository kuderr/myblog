package models

type Post struct {
	ID      int
	Title   string
	Summary string
	Body    string
	Author  string
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
