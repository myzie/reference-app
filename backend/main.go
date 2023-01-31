package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/ziflex/lecho/v3"
)

type Database struct {
	Name          string `json:"name"`
	EngineVersion string `json:"engine_version"`
	NodeCount     int    `json:"node_count"`
}

var databases []*Database

func main() {

	logger := zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()
	echoLogger := lecho.From(logger)

	addDatabase := func(db *Database) error {
		databases = append(databases, db)
		return nil
	}

	addDatabase(&Database{Name: "dev-1", EngineVersion: "15.1", NodeCount: 2})
	addDatabase(&Database{Name: "dev-2", EngineVersion: "14.6", NodeCount: 3})
	addDatabase(&Database{Name: "staging", EngineVersion: "15.1", NodeCount: 5})
	addDatabase(&Database{Name: "production", EngineVersion: "14.6", NodeCount: 5})

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
		return c.JSON(http.StatusOK, databases)
	})

	e.POST("/databases", func(c echo.Context) error {
		type input struct {
			Name          string `json:"name"`
			EngineVersion string `json:"engine_version"`
			NodeCount     int    `json:"node_count"`
		}
		var in input
		if err := c.Bind(&in); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, &Database{
			Name:          in.Name,
			EngineVersion: in.EngineVersion,
			NodeCount:     in.NodeCount,
		})
	})

	logger.Info().Msg("starting server")

	e.Logger.Fatal(e.Start(":1323"))
}
