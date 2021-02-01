package article

import (
	"cqrses/models"
	"time"
)

// BindArticleStore bind the ArticleStore entity in the Article entity
func BindArticleStore(articleStore *models.ArticleStore) (models.Article, error) {

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
