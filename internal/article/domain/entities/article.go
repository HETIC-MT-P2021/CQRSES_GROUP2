package entities

import "time"

// Article defines the structure of the article entities
type Article struct {
	Title     string    `form:"title"`
	Content   string    `form:"content"`
	CreatedAt time.Time `form:"created_at"`
	CreatedBy uint      `form:"created_by"`
}

// ArticleData defines the structure of the article entities for the action command
type ArticleData struct {
	Title     string `form:"title"`
	Content   string `form:"content"`
	CreatedBy uint   `form:"created_by"`
}
