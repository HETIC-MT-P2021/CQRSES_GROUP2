package command

import (
	"errors"
	"time"

	"cqrses/internal/article/domain/entities"
	"cqrses/internal/article/domain/model"
	"cqrses/pkg/cqrs"
)

// ArticleCommandMessage is the struct we use to action a new command
type ArticleCommandMessage struct {
	ArticleData entities.ArticleData
}

// CommandHandler todo
type CommandHandler struct{}

// UpdateArticleCommandHandler action a new command handler for article
func CreateArticleCommandHandler() *CommandHandler {
	return &CommandHandler{}
}

// Handle handle the command
func (ach *CommandHandler) Handle(command cqrs.CommandMessage) (interface{}, error) {
	switch cmd := command.Payload().(type) {
	case *ArticleCommandMessage:
		article, err := BindArticleAndCreate(&cmd.ArticleData)
		return article, err
	default:
		return nil, errors.New("bad command type")
	}
}

// BindArticleAndCreate bind the ArticleData entities in the Article entities
func BindArticleAndCreate(articleData *entities.ArticleData) (interface{}, error) {
	article := entities.Article{
		Title:     articleData.Title,
		Content:   articleData.Content,
		CreatedAt: time.Now(),
		CreatedBy: articleData.CreatedBy,
	}

	document, err := model.StoreArticle(&article)
	if err != nil {
		return entities.Article{}, err
	}

	return document, err
}
