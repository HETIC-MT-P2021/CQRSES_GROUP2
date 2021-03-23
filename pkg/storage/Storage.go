package storage

type Document struct {
	ID   string      `json:"id"`
	Body interface{} `json:"body"`
}

type Storage interface {
	Save(index string, document *Document) error
	Search(index string, objectID string) ([]*Document, error)
}
