package datastore

import (
	"sync"

	"cqrses/pkg/storage"
)

var once sync.Once
var ds *DataStore

func GetInstance() *DataStore {
	once.Do(func() { ds = new() })
	return ds
}

type DataStore struct {
	storage storage.Storage
}

// new returns an initialized DataStore instance.
func new() *DataStore {
	ds := &DataStore{}

	return ds
}

func (ds *DataStore) SetStorage(storage storage.Storage) {
	ds.storage = storage
}

// Save creates a new object
func Save(index string, document *storage.Document) error { return ds.Save(index, document) }
func (ds *DataStore) Save(index string, document *storage.Document) error {
	return ds.storage.Save(index, document)
}

func Search(index string, objectID string) (interface{}, error) { return ds.Search(index, objectID) }
func (ds *DataStore) Search(index string, objectID string) (interface{}, error) {
	return ds.storage.Search(index, objectID)
}