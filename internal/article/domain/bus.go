package domain

import (
	"cqrses/internal/article/domain/command"
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
}

// SetArticleQueries defines all article queries.
func SetArticleQueries(queryBus *cqrs.QueryBus) {
	// TODO
}
