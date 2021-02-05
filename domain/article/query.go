package article

import "cqrses/cqrs"

// GetArticleQuery todo
type GetArticleQuery struct{}

// QueryHandler todo
type QueryHandler struct{}

// NewArticleQueryHandler create a new artcile query handler
func NewArticleQueryHandler() *QueryHandler {
	return &QueryHandler{}
}

// Handle the article query handler
func (aqh *QueryHandler) Handle(command cqrs.QueryMessage) (interface{}, error) {
	switch command.Payload() {
	case GetArticleQuery{}:
		return nil, nil
	default:
		return nil, nil
	}
}
