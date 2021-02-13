package storage

var es *EventStore

func init() {
	es = New()
}

type Document struct {
	ID   string
	Body interface{}
}

type Storage interface {
	Save(index string, document *Document) error
}

type EventStore struct {
	storage Storage
}

// New returns an initialized EventStore instance.
func New() *EventStore {
	es := new(EventStore)

	return es
}

func SetStorage(storage Storage) { es.SetStorage(storage) }
func (es *EventStore) SetStorage(storage Storage) {
	es.storage = storage
}

// Save creates a new object
func Save(index string, document *Document) error { return es.Save(index, document) }
func (es *EventStore) Save(index string, document *Document) error {
	return es.storage.Save(index, document)
}
