package model

import (
	"cqrses/internal/article/domain/entities"
	"time"

	"cqrses/pkg/storage"
	"cqrses/pkg/storage/eventstore"
)

// StoreArticle saves an article in es
func StoreArticle(article *entities.Article) (interface{}, error) {
	eventName := "article"

	document := storage.Document{
		Body: eventstore.Event{
			EventName:      eventName,
			EventType:      eventstore.Create,
			ObjectID:       storage.NewObjectID().String(),
			Payload:        article,
			EventCreatedAt: time.Now(),
		},
	}

	err := eventstore.Save(eventName, &document)
	if err != nil {
		return nil, err
	}

	return document, err
}

// UpdateArticle saves an modified article in es
func UpdateArticle(objectID string, article *entities.Article) (interface{}, error) {
	eventName := "article"
	objID, err := storage.ParseObjectID(objectID)
	if err != nil {
		return nil, err
	}

	document := storage.Document{
		Body: eventstore.Event{
			EventName:      eventName,
			EventType:      eventstore.Update,
			ObjectID:       objID.String(),
			Payload:        article,
			EventCreatedAt: time.Now(),
		},
	}

	err = eventstore.Save(eventName, &document)
	if err != nil {
		return nil, err
	}

	return document, err
}
