package elasticsearch

import (
	"context"
	"fmt"

	"cqrses/storage"
	"github.com/olivere/elastic/v7"
)

type Storage interface {
	storage.Storage
}

type ElasticSearch struct {
	Context context.Context
	Client  *elastic.Client
}

// New returns an initialized ElasticSearch instance.
func New(client *elastic.Client) *ElasticSearch {
	es := new(ElasticSearch)
	es.Context = context.Background()
	es.Client = client

	return es
}

// createIndex returns a service to create a new index.
func (es *ElasticSearch) createIndex(index string) error {
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

// Save creates a new document
func (es *ElasticSearch) Save(index string, document *storage.Document) error {
	client := es.Client

	exists, _ := client.IndexExists(index).Do(es.Context)
	if !exists {
		if err := es.createIndex(index); err != nil {
			return err
		}
	}

	indexedDocument, err := client.Index().
		Index(index).
		BodyJson(document.Body).
		Do(es.Context)

	if err != nil {
		return fmt.Errorf("cannot add resource in index %s", index)
	}

	document.ID = indexedDocument.Id

	return nil
}
