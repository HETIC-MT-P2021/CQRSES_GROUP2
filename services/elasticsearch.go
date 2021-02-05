package services

import (
	"context"
	"fmt"

	"github.com/olivere/elastic/v7"
)

type Document struct {
	ID   string
	Body interface{}
}

type ElasticInterface interface {
	CreateIndex(index string) error
	CreateNewDocument(index string, document *Document) error
}

type ElasticService struct {
	Context context.Context
	Client  *elastic.Client
}

func NewsElastic(client *elastic.Client) *ElasticService {
	es := new(ElasticService)
	es.Context = context.Background()
	es.Client = client

	return es
}

// CreateIndex returns a service to create a new index.
func (es ElasticService) CreateIndex(index string) error {
	client := es.Client

	// Use the IndexExists service to check if a specified index exists.
	exists, err := client.IndexExists(index).Do(es.Context)
	if err != nil {
		panic(err)
	}
	if exists {
		return fmt.Errorf("index already exists: %s", err)
	}

	createIndex, err := client.CreateIndex(index).Do(es.Context)
	if err != nil {
		return fmt.Errorf("error creating the indexer: %s", err)
	}

	if !createIndex.Acknowledged {
		return fmt.Errorf("index not acknowledged: %s", err)
	}

	return nil
}

//CreateNewDocument creates a new document
func (es ElasticService) CreateNewDocument(index string, document *Document) error {
	client := es.Client

	exists, err := client.IndexExists(index).Do(es.Context)
	if !exists {
		if err := es.CreateIndex(index); err != nil {
			return err
		}
	}

	_, err = client.Index().
		Index(index).
		BodyJson(document.Body).
		Do(es.Context)

	if err != nil {
		return fmt.Errorf("cannot add resource in index %s", index)
	}

	return nil
}
