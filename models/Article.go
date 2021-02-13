package models

import (
	"time"

	"cqrses/storage"
	"cqrses/storage/eventstore"

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
func StoreArticle(article *Article) (interface{}, error) {
	eventName := "article"

	document := storage.Document{
		Body: eventstore.Event{
			Name:      eventName,
			Typology:  eventstore.Create,
			ObjectID:  uuid.NewString(),
			Payload:   article,
			CreatedAt: time.Now(),
		},
	}

	err := eventstore.Save(eventName, &document)
	if err != nil {
		return nil, err
	}

	return document, err
}

// UpdateArticle saves an modified article in es
func UpdateArticle(objectID string, article *Article) (interface{}, error) {
	eventName := "article"

	document := storage.Document{
		Body: eventstore.Event{
			Name:      eventName,
			Typology:  eventstore.Update,
			ObjectID:  objectID,
			Payload:   article,
			CreatedAt: time.Now(),
		},
	}

	err := eventstore.Save(eventName, &document)
	if err != nil {
		return nil, err
	}

	return document, err
}
