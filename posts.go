package main

import (
	"database/sql"
	"log"
	"time"
)

type Post struct {
	ID        int
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// CreatePost inserts a new post and returns it
func CreatePost(db *sql.DB, title, content string) (Post, error) {
	res, err := db.Exec(`
		INSERT INTO posts (title, content)
		VALUES (?, ?)
	`, title, content)
	if err != nil {

		log.Println(err)
		return Post{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
		return Post{}, err
	}

	var createdAtStr, updatedAtStr string
	err = db.QueryRow(`
		SELECT created_at, updated_at
		FROM posts
		WHERE id = ?
	`, id).Scan(&createdAtStr, &updatedAtStr)
	if err != nil {

		log.Println(err)
		return Post{}, err
	}

	createdAt, _ := time.Parse("2006-01-02 15:04:05", createdAtStr)
	updatedAt, _ := time.Parse("2006-01-02 15:04:05", updatedAtStr)

	return Post{
		ID:        int(id),
		Title:     title,
		Content:   content,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}

// GetAllPosts retrieves all blog posts from the database
func GetAllPosts(db *sql.DB) ([]Post, error) {
	rows, err := db.Query(`
		SELECT id, title, content, created_at, updated_at
		FROM posts
		ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post

	for rows.Next() {
		var p Post
		var createdAtStr, updatedAtStr string

		if err := rows.Scan(&p.ID, &p.Title, &p.Content, &createdAtStr, &updatedAtStr); err != nil {
			log.Println("row scan error:", err)
			continue
		}

		// Parse timestamps
		p.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAtStr)
		p.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", updatedAtStr)

		posts = append(posts, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}
