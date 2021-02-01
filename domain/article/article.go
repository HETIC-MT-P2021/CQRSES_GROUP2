package article

import (
	"cqrses/models"
	"time"
)

// BindArticleStore bind the ArticleForm entity in the Article entity
func BindArticleStore(articleForm *models.ArticleStore) (models.Article, error) {

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
