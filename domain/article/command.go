package article

import (
	"cqrses/cqrs"
	"cqrses/models"
	"time"
)

// CreateArticleCommand is the struct we use to create a new command
type CreateArticleCommand struct {
	ArticleStore models.ArticleStore
}

// EditArticleCommand todo
type EditArticleCommand struct {
	AggregateID  string
	ArticleStore models.ArticleStore
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
		article, err := BindArticleAndCreate(&cmd.ArticleStore)
		return article, err
	case *EditArticleCommand:
		article, err := BindArticleAndUpdate(cmd.AggregateID, &cmd.ArticleStore)
		return article, err
	default:
		return nil, nil
	}
}

// BindArticleAndCreate bind the ArticleStore entity in the Article entity
func BindArticleAndCreate(articleStore *models.ArticleStore) (models.Article, error) {

	article := models.Article{
		AuthorID:  articleStore.AuthorID,
		Title:     articleStore.Title,
		Content:   articleStore.Content,
		CreatedAt: time.Now(),
	}

	if err := models.StoreArticle(&article); err != nil {
		return models.Article{}, err
	}

	return article, nil
}

// BindArticleAndUpdate bind the ArticleStore entity in the Article entity
func BindArticleAndUpdate(AggregateID string, articleStore *models.ArticleStore) (models.Article, error) {

	article := models.Article{
		AuthorID:  articleStore.AuthorID,
		Title:     articleStore.Title,
		Content:   articleStore.Content,
		CreatedAt: time.Now(),
	}

	if err := models.UpdateArticle(AggregateID, &article); err != nil {
		return models.Article{}, err
	}

	return article, nil
}
