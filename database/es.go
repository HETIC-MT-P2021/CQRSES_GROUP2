package database

import (
	"fmt"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/spf13/viper"
)

// Es public declaration
var EsClient *elasticsearch.Client

// GetEsConfig return config variables of the elasticsearch.
func GetEsConfig() elasticsearch.Config {
	return elasticsearch.Config{
		Addresses: strings.Split(viper.GetString("ES_URL"), ","),
	}
}

// EsConnect connect to the elasticsearch.
func EsConnect() {
	cfg := GetEsConfig()

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic(fmt.Errorf("error creating the client: %s", err))
	} else {
		log.Println(es.Info())
	}

	EsClient = es
}
