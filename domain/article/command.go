package article

import (
	"cqrses/cqrs"
	"cqrses/models"
	"time"
)

// CreateArticleCommand is the struct we use to create a new command
type CreateArticleCommand struct {
	ArticleData models.ArticleData
}

// EditArticleCommand todo
type EditArticleCommand struct {
	AggregateID string
	ArticleData models.ArticleData
}

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
		article, err := BindArticleAndCreate(&cmd.ArticleData)
		return article, err
	case *EditArticleCommand:
		article, err := BindArticleAndUpdate(cmd.AggregateID, &cmd.ArticleData)
		return article, err
	default:
		return nil, nil
	}
}

// BindArticleAndCreate bind the ArticleData entity in the Article entity
func BindArticleAndCreate(articleData *models.ArticleData) (interface{}, error) {

	article := models.Article{
		Title:     articleData.Title,
		Content:   articleData.Content,
		CreatedAt: time.Now(),
		CreatedBy: articleData.CreatedBy,
	}

	document, err := models.StoreArticle(&article)
	if err != nil {
		return models.Article{}, err
	}

	return document, err
}

// BindArticleAndUpdate bind the ArticleData entity in the Article entity
func BindArticleAndUpdate(AggregateID string, articleData *models.ArticleData) (models.Article, error) {

	article := models.Article{
		Title:     articleData.Title,
		Content:   articleData.Content,
		CreatedAt: time.Now(),
		CreatedBy: articleData.CreatedBy,
	}

	if err := models.UpdateArticle(AggregateID, &article); err != nil {
		return models.Article{}, err
	}

	return article, nil
}
