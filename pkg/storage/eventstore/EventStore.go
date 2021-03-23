package eventstore

import (
	"sync"

	"cqrses/pkg/storage"
)

var once sync.Once
var es *EventStore

func GetInstance() *EventStore {
	once.Do(func() { es = new() })
	return es
}

type EventStore struct {
	storage storage.Storage
}

// new returns an initialized EventStore instance.
func new() *EventStore {
	es := &EventStore{}

	return es
}

func (es *EventStore) SetStorage(storage storage.Storage) {
	es.storage = storage
}

// Save creates a new object
func Save(index string, document *storage.Document) error { return es.Save(index, document) }
func (es *EventStore) Save(index string, document *storage.Document) error {
	return es.storage.Save(index, document)
}
