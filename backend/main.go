package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ZachIgarz/test-api-rest/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {

	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.CORS())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	//e.GET("/healthz", healthHandler)

	e = router.NewRouter(e, appController)

	go func() {
		if err := e.Start(config.HTTPListener()); err != nil {
			log.Error("shutting down the server")
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		clog.Panic(log, merry.Wrap(err))
	}
}
