package article

import (
	"cqrses/cqrs"
	"cqrses/models"
)

// CreateArticleCommand is the struct we use to create a new command
type CreateArticleCommand struct {
	ArticleStore models.ArticleStore
}

// CommandType todo
func (c CreateArticleCommand) CommandType() string {
	return "CreateArticleCommand"
}

// Payload todo
func (c CreateArticleCommand) Payload() interface{} {
	return &c
}

// CommandHandler todo
type CommandHandler struct{}

// NewArticleCommandHandler create a new command handler for article
func NewArticleCommandHandler() *CommandHandler {
	return &CommandHandler{}
}

// Handle handle the command
func (ach *CommandHandler) Handle(command cqrs.CommandMessage) (interface{}, error) {
	// Return command type
	switch command.CommandType() {
	case "CreateArticleCommand":
		payload := command.Payload().(*CreateArticleCommand)
		// article, err := validateAndPersistArticle(&payload.ArticleForm)
		// return article, err
		var err error
		return &payload.ArticleStore, err
	case "EditArticleCommand":
		return nil, nil
	default:
		return nil, nil
	}
}
