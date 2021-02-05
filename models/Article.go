package models

import (
	"time"
)

// Article defines the structure of the article entity
type Article struct {
	AuthorID  uint      `form:"author_id"`
	Title     string    `form:"title"`
	Content   string    `form:"content"`
	CreatedAt time.Time `form:"created_at"`
}

// ArticleStore defines the structure of the article entity for the create command
type ArticleStore struct {
	AuthorID uint   `form:"author_id"`
	Title    string `form:"title"`
	Content  string `form:"content"`
}

// StoreArticle saves an article in es
func StoreArticle(article *Article) error {

	// document := services.Document{
	// 	Body: es.Event{
	//	    AggregateID: uuid.NewV4().String(),
	// 		Typology:  es.Create,
	// 		Payload:   article,
	// 		CreatedAt: time.Now(),
	// 		Index:     1,
	// 	},
	// }

	// err := services.CreateNewDocumentInIndex("article", &document)
	// if err != nil {
	// 	log.Error("Error while creating event : ", err)
	// 	return err
	// }
	return nil
}

// UpdateArticle saves an modificated article in es
func UpdateArticle(AggregateID string, article *Article) error {
	return nil
}
