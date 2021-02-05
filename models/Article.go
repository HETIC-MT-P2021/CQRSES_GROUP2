package models

import (
	"time"
)

// Article defines the structure of the article entity
type Article struct {
	Title     string    `form:"title"`
	Content   string    `form:"content"`
	CreatedAt time.Time `form:"created_at"`
	CreatedBy uint      `form:"created_by"`
}

// ArticleData defines the structure of the article entity for the create command
type ArticleData struct {
	Title     string `form:"title"`
	Content   string `form:"content"`
	CreatedBy uint   `form:"created_by"`
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
