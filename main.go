package main

import (
	"covid-app/config"
	"covid-app/controllers"
	_ "covid-app/docs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
)


// @title Covid-19 Stats API
// @version 1.0
// @description This is a sample server Covid-19 stats.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /
func main() {
	cacheClient, _ := config.RedisConnect()
	config.MongoConnect()
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	e.Use(cacheClient.Middleware())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", welcome)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/fetch-data", controllers.FetchData)
	e.GET("/get-data", controllers.GetData)

	e.Logger.Fatal(e.Start(":8000"))
}

func welcome(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to Covid API")
}
