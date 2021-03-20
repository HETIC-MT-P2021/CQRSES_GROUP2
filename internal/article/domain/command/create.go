package command

import (
	"errors"
	"time"

	"cqrses/internal/article/domain/entities"
	"cqrses/internal/article/domain/repository"
	"cqrses/pkg/cqrs"
)

// CreateArticleCommandMessage is the struct we use to action a new command
type CreateArticleCommandMessage struct {
	ArticleData entities.ArticleData
}

// CreateArticleCommandHandler todo
type CreateArticleCommandHandler struct {
	repo *repository.ArticleRepository
}

// NewCreateArticleCommandHandler action a new command handler for article
func NewCreateArticleCommandHandler() *CreateArticleCommandHandler {
	h := new(CreateArticleCommandHandler)
	h.repo = repository.NewArticleRepository()

	return h
}

// Handle handle the command
func (ach *CreateArticleCommandHandler) Handle(command cqrs.CommandMessage) (interface{}, error) {
	switch cmd := command.Payload().(type) {
	case *CreateArticleCommandMessage:
		article := entities.Article{
			Title:     cmd.ArticleData.Title,
			Content:   cmd.ArticleData.Content,
			CreatedAt: time.Now(),
			CreatedBy: cmd.ArticleData.CreatedBy,
		}

		document, err := ach.repo.Create(&article)
		if err != nil {
			return entities.Article{}, err
		}

		return document, err
	default:
		return nil, errors.New("bad command type")
	}
}
