package domain

import (
	"cqrses/cqrs"

	"cqrses/domain/article"
)

// Cb define the command bus
var Cb *cqrs.CommandBus

// Qb define the query bus
var Qb *cqrs.QueryBus

// InitBus initializes the command and query buses
func InitBus() {
	Cb = cqrs.NewCommandBus()
	Qb = cqrs.NewQueryBus()

	_ = Cb.RegisterHandler(article.NewArticleCommandHandler(), &article.CreateArticleCommand{})
}
