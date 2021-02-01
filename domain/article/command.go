package article

import (
	"cqrses/cqrs"
	"cqrses/models"
)

// CreateArticleCommand is the struct we use to create a new command
type CreateArticleCommand struct {
	ArticleStore models.ArticleStore
}

// EditArticleCommand todo
type EditArticleCommand struct{}

// CommandHandler todo
type CommandHandler struct{}

// NewArticleCommandHandler create a new command handler for article
func NewArticleCommandHandler() *CommandHandler {
	return &CommandHandler{}
}

// Handle handle the command
func (ach *CommandHandler) Handle(command cqrs.CommandMessage) (interface{}, error) {
	switch cmd := command.Payload().(type) {
	case *CreateArticleCommand:
		article, err := BindArticleFrom(&cmd.ArticleStore)
		return article, err
	case *EditArticleCommand:
		return nil, nil
	default:
		return nil, nil
	}
}
