package domain

import (
	"cqrses/internal/article/domain/command"
	"cqrses/internal/article/domain/query"
	"cqrses/pkg/cqrs"
)

// todo

// CommandBus define the command bus
var CommandBus *cqrs.CommandBus

// QueryBus define the query bus
var QueryBus *cqrs.QueryBus

func InitBusses() {
	CommandBus = cqrs.NewCommandBus()
	QueryBus = cqrs.NewQueryBus()

	SetArticleCommands(CommandBus)
	SetArticleQueries(QueryBus)
}

// SetArticleCommands defines all article commands.
func SetArticleCommands(cmdBus *cqrs.CommandBus) {
	_ = cmdBus.RegisterHandler(command.NewCreateArticleCommandHandler(), &command.CreateArticleCommandMessage{})
	_ = cmdBus.RegisterHandler(command.NewUpdateArticleCommandHandler(), &command.UpdateArticleCommandMessage{})
}

// SetArticleQueries defines all article queries.
func SetArticleQueries(queryBus *cqrs.QueryBus) {
	_ = queryBus.RegisterHandler(query.NewArticleQueryHandler(), &query.GetArticleQuery{})
}
