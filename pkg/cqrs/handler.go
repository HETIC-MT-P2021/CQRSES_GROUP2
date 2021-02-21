package cqrs

// CommandHandler define the structure of the command handler interface
type CommandHandler interface {
	Handle(CommandMessage) (interface{}, error)
}

// QueryHandler define the structure of the query handler interface
type QueryHandler interface {
	Handle(QueryMessage) (interface{}, error)
}
