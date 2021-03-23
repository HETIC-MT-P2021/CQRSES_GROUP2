package query

import (
	"cqrses/internal/article/domain/repository"
	"cqrses/pkg/cqrs"
)

// GetArticleQuery todo
type GetArticleQuery struct {
	ObjectID string
}

// ArticleQueryHandler todo
type ArticleQueryHandler struct {
	repo *repository.ArticleRepository
}

// NewArticleQueryHandler create a new article query handler
func NewArticleQueryHandler() *ArticleQueryHandler {
	h := new(ArticleQueryHandler)
	h.repo = repository.NewArticleRepository()

	return h
}

// Handle the article query handler
func (aqh *ArticleQueryHandler) Handle(query cqrs.QueryMessage) (interface{}, error) {
	switch q := query.Payload().(type) {
	case *GetArticleQuery:
		document, err := aqh.repo.Find(q.ObjectID)
		if err != nil {
			return nil, err
		}

		return document, nil
	default:
		return nil, nil
	}
}
