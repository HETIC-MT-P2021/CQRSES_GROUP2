package storage

type Document struct {
	ID   string
	Body interface{}
}

type Storage interface {
	Save(index string, document *Document) error
}
