package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/ziflex/lecho/v3"
)

func main() {

	logger := zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()
	echoLogger := lecho.From(logger)

	e := echo.New()
	e.HideBanner = true
	e.Logger = echoLogger

	// Install middleware
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(lecho.Middleware(lecho.Config{Logger: echoLogger}))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "http://localhost:8080"},
		AllowHeaders: []string{"*"},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/databases", func(c echo.Context) error {
		type db struct {
			Name    string `json:"name"`
			Version string `json:"version"`
			Engine  string `json:"engine"`
		}
		var dbs []db
		dbs = append(dbs, db{Name: "db1", Version: "1.0", Engine: "mysql"})
		dbs = append(dbs, db{Name: "db2", Version: "2.0", Engine: "postgres"})
		return c.JSON(http.StatusOK, dbs)
	})

	logger.Info().Msg("starting server")

	e.Logger.Fatal(e.Start(":1323"))
}
