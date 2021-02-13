// Back-end API for a form management application
package main

import (
	"cqrses/config"
	"cqrses/database"
	"cqrses/database/factory"
	"cqrses/domain"
	"cqrses/router"
	"cqrses/storage"
	"cqrses/storage/elasticsearch"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// main launch all part of the project
func main() {

	config.SetConfig()

	database.Connect()
	factory.Migrate()
	factory.Seed()

	database.GetElasticCon()
	elastic := elasticsearch.New(database.ElasticConn)
	storage.SetStorage(elastic)

	e := echo.New()

	router.InitRoutes(e)
	domain.InitBus()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.Use(middleware.RequestID())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())

	e.Logger.Fatal(e.Start(":80"))
}
