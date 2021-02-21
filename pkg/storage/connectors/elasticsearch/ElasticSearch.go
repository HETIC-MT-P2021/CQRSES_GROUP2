package elasticsearch

import (
	"context"
	"fmt"

	"cqrses/pkg/storage"

	"github.com/olivere/elastic/v7"
)

type Connector interface {
	storage.Storage
	connect(options ...elastic.ClientOptionFunc) *elastic.Client
}

type ElasticSearch struct {
	Context context.Context
	Client  *elastic.Client
}

// NewConnector create a new ElasticSearch connector instance.
func NewConnector(options ...elastic.ClientOptionFunc) *ElasticSearch {
	es := new(ElasticSearch)
	es.Context = context.Background()
	es.Client = es.connect(options...)

	return es
}

// connect establish a Elasticsearch connection.
func (es *ElasticSearch) connect(options ...elastic.ClientOptionFunc) *elastic.Client {
	client, err := elastic.NewClient(options...)

	if err != nil {
		panic(fmt.Errorf("error creating the client: %s", err))
	}

	return client
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
