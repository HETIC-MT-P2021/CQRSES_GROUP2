package service

// RepositoryInterface todo
type RepositoryInterface interface {
	Find()
	Create()
	Update()
	Delete()
}

// Repository todo
type Repository struct {
	Name string
	// mux  sync.RWMutex // must be re implemented in repository
}
