package main

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
)
func init() {
	clog.SetDefaultLog()
}

func main() {


	appController := controller.AppController{
		Organization:    controller.OrgController,
		UserPermissions: controller.UserPermController,
		User:            controller.UserController,
		Project:         controller.ProController,
		Did:             controller.DiController,
		SipTrunk:        controller.SipTController,
		Email:           controller.Email,
	}
	e := echo.New()
	loggerMiddleware := clog.LoggerMiddlewareWithOptions(logrus.StandardLogger(), &clog.LoggerMiddlewareOpts{
		DefaultMetrics:   true,
		MetricsNamespace: "cpaas",
		MetricsSubsystem: "webapi",
	})

	e.HTTPErrorHandler = clog.ErrHandler(logrus.StandardLogger())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(loggerMiddleware)
	e.Use(middleware.CORS())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	e.GET("/healthz", healthHandler)

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