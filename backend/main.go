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

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	logger.Info().Msg("starting server")

	e.Logger.Fatal(e.Start(":1323"))
}
