// Back-end API for a form management application
package main

import (
	"cqrses/app/config"
	"cqrses/app/database"
	"cqrses/app/database/factory"
	"cqrses/app/router"
	"cqrses/internal/article/domain"
	"cqrses/pkg/storage/connectors/elasticsearch"
	"cqrses/pkg/storage/eventstore"
	"cqrses/producer"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/olivere/elastic/v7"
	"github.com/spf13/viper"
)

// main launch all part of the project
func main() {
	config.SetConfig()

	database.Connect()
	factory.Migrate()
	factory.Seed()

	eventstore.GetInstance()

	elasticSearch := elasticsearch.NewConnector(
		elastic.SetSniff(true),
		elastic.SetURL(viper.GetString("ES_URL")),
	)

	eventstore.SetStorage(elasticSearch)

	producer.SendMessage()

	e := echo.New()

	router.InitRoutes(e)
	domain.InitBusses()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.Use(middleware.RequestID())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())

	e.Logger.Fatal(e.Start(":80"))
}
