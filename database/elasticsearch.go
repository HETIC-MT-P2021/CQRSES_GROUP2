package database

import (
	"fmt"

	"github.com/olivere/elastic/v7"
	"github.com/spf13/viper"
)

// ElasticConn public declaration
var ElasticConn *elastic.Client

// GetElasticCon creates a new client Elasticsearch.
func GetElasticCon() {
	client, err := elastic.NewClient(
		elastic.SetSniff(true),
		elastic.SetURL(viper.GetString("ES_URL")),
	)

	if err != nil {
		panic(fmt.Errorf("error creating the client: %s", err))
	}

	ElasticConn = client
}
