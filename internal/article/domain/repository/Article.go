package repository

import (
	"sync"
	"time"

	"cqrses/internal/article/domain/entities"
	"cqrses/internal/pkg/service"
	"cqrses/pkg/storage"
	"cqrses/pkg/storage/eventstore"
)

// ArticleRepository todo
type ArticleRepository struct {
	service.Repository
	mux sync.RWMutex
}

func NewArticleRepository() *ArticleRepository {
	r := new(ArticleRepository)
	r.Name = "article"

	return r
}

func (r ArticleRepository) Find() {
	panic("implement me")
}

// Create saves an article in es
func (r ArticleRepository) Create(article *entities.Article) (interface{}, error) {
	r.mux.RLock()
	defer r.mux.RUnlock()

	document := storage.Document{
		Body: eventstore.Event{
			EventName:      r.Name,
			EventType:      eventstore.Create,
			ObjectID:       storage.NewObjectID().String(),
			ObjectVersion:  storage.NewObjectVersion(0),
			Payload:        article,
			EventCreatedAt: time.Now(),
		},
	}

	err := eventstore.Save(r.Name, &document)
	if err != nil {
		return nil, err
	}

	return document, err
}

// Update saves an modified article in es
func (r ArticleRepository) Update(objectID string, article *entities.Article) (interface{}, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	objID, err := storage.ParseObjectID(objectID)
	if err != nil {
		return nil, err
	}

	document := storage.Document{
		Body: eventstore.Event{
			EventName:      r.Name,
			EventType:      eventstore.Update,
			ObjectID:       objID.String(),
			ObjectVersion:  storage.NewObjectVersion(0).Next(),
			Payload:        article,
			EventCreatedAt: time.Now(),
		},
	}

	err = eventstore.Save(r.Name, &document)
	if err != nil {
		return nil, err
	}

	return document, err
}

func (r ArticleRepository) Delete() {
	panic("implement me")
}
