package article

import (
	"cqrses/models"
	"time"
)

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
