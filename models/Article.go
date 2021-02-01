package models

import (
	"time"
)

// Article defines the structure of the article entity
type Article struct {
	AuthorID  uint64
	Title     string
	Content   string
	CreatedAt time.Time
}

// ArticleStore defines the structure of the article entity for the create command
type ArticleStore struct {
	AuthorID uint64
	Title    string
	Content  string
}
