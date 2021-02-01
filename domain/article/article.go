package article

import (
	"cqrses/models"
	"time"
)

// BindArticleFrom bind the ArticleForm entity in the Article entity
func BindArticleFrom(articleForm *models.ArticleStore) (models.Article, error) {

	article := models.Article{
		AuthorID:  articleForm.AuthorID,
		Title:     articleForm.Title,
		Content:   articleForm.Content,
		CreatedAt: time.Now(),
	}

	if err := models.StoreArticle(&article); err != nil {
		return models.Article{}, err
	}

	return article, nil
}
