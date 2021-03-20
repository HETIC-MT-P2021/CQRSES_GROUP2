package command

import (
	"errors"
	"time"

	"cqrses/internal/article/domain/entities"
	"cqrses/internal/article/domain/repository"
	"cqrses/pkg/cqrs"
)

// UpdateArticleCommandMessage todo
type UpdateArticleCommandMessage struct {
	ObjectID    string
	ArticleData entities.ArticleData
}

// UpdateArticleCommandHandler todo
type UpdateArticleCommandHandler struct {
	repo *repository.ArticleRepository
}

// UpdateArticleCommandHandler action a new command handler for article
func NewUpdateArticleCommandHandler() *UpdateArticleCommandHandler {
	h := new(UpdateArticleCommandHandler)
	h.repo = repository.NewArticleRepository()

	return h
}

// Handle handle the command
func (uch *UpdateArticleCommandHandler) Handle(command cqrs.CommandMessage) (interface{}, error) {
	switch cmd := command.Payload().(type) {
	case *UpdateArticleCommandMessage:
		article := entities.Article{
			Title:     cmd.ArticleData.Title,
			Content:   cmd.ArticleData.Content,
			CreatedAt: time.Now(),
			CreatedBy: cmd.ArticleData.CreatedBy,
		}

		document, err := uch.repo.Update(cmd.ObjectID, &article)
		if err != nil {
			return entities.Article{}, err
		}

		return document, err
	default:
		return nil, errors.New("bad command type")
	}
}
