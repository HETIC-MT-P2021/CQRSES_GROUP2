package models

import (
	db "cqrses/database"
	"cqrses/services"
	"time"

	"github.com/google/uuid"
)

// Article defines the structure of the article entity
type Article struct {
	Title     string    `form:"title"`
	Content   string    `form:"content"`
	CreatedAt time.Time `form:"created_at"`
	CreatedBy uint      `form:"created_by"`
}

// ArticleData defines the structure of the article entity for the create command
type ArticleData struct {
	Title     string `form:"title"`
	Content   string `form:"content"`
	CreatedBy uint   `form:"created_by"`
}

// StoreArticle saves an article in es
func StoreArticle(article *Article) error {
	eventName := "article"
	es := services.NewsElastic(db.ElasticConn)

	document := services.Document{
		Body: Event{
			Name:      eventName,
			Typology:  Create,
			ObjectID:  uuid.NewString(),
			Payload:   article,
			CreatedAt: time.Now(),
		},
	}

	err := es.CreateNewDocument(eventName, &document)
	if err != nil {
		return err
	}
	return nil
}

// UpdateArticle saves an modificated article in es
func UpdateArticle(AggregateID string, article *Article) error {
	return nil
}
