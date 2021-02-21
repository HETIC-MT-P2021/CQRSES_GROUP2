package command

import (
	"cqrses/internal/article/domain/entities"
	"cqrses/internal/article/domain/model"
	"time"
)

// EditArticleCommand todo
type EditArticleCommand struct {
	ObjectID    string
	ArticleData entities.ArticleData
}

// UpdateArticleCommandHandler action a new command handler for article
func UpdateArticleCommandHandler() *CommandHandler {
	return &CommandHandler{}
}

// BindArticleAndUpdate bind the ArticleData entities in the Article entities
func BindArticleAndUpdate(objectID string, articleData *entities.ArticleData) (interface{}, error) {
	article := entities.Article{
		Title:     articleData.Title,
		Content:   articleData.Content,
		CreatedAt: time.Now(),
		CreatedBy: articleData.CreatedBy,
	}

	document, err := model.UpdateArticle(objectID, &article)
	if err != nil {
		return entities.Article{}, err
	}

	return document, err
}
