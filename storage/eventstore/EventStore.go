package eventstore

import (
	"sync"

	"cqrses/storage"
)

var once sync.Once
var es *EventStore

func GetInstance() *EventStore {
	once.Do(func() { es = New() })
	return es
}

type EventStore struct {
	storage storage.Storage
}

// New returns an initialized EventStore instance.
func New() *EventStore {
	es := new(EventStore)

	return es
}

func SetStorage(storage storage.Storage) { es.SetStorage(storage) }
func (es *EventStore) SetStorage(storage storage.Storage) {
	es.storage = storage
}

// Save creates a new object
func Save(index string, document *storage.Document) error { return es.Save(index, document) }
func (es *EventStore) Save(index string, document *storage.Document) error {
	return es.storage.Save(index, document)
}
