package models

import (
	"time"
)

// Article defines the structure of the article entity
type Article struct {
	AuthorID  uint      `form:"author_id"`
	Title     string    `form:"title"`
	Content   string    `form:"content"`
	CreatedAt time.Time `form:"created_at"`
}

// ArticleStore defines the structure of the article entity for the create command
type ArticleStore struct {
	AuthorID uint   `form:"author_id"`
	Title    string `form:"title"`
	Content  string `form:"content"`
}
