package http

import (
	"context"
	container "gochat/app"
	"gochat/app/http/middlewares"
	"gochat/infra/config"
	"gochat/infra/logger"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
)

func Start() {
	e := echo.New()

	if err := middlewares.Attach(e); err != nil {
		logger.Error("error occur when attaching middlewares", err)
		os.Exit(1)
	}

	// routes for documentation
	dg := e.Group("docs")
	dg.GET("/swagger", echo.WrapHandler(middlewares.SwaggerDocs()))
	dg.GET("/redoc", echo.WrapHandler(middlewares.ReDocDocs()))
	dg.GET("/rapidoc", echo.WrapHandler(middlewares.RapiDocs()))
	e.Static("/swagger.yaml", "./swagger.yaml")

	// Create Prometheus server and Middleware
	echoPrometheus := echo.New()
	echoPrometheus.HideBanner = true
	prom := prometheus.NewPrometheus("echo", nil)

	// Scrape metrics from Main Server
	e.Use(prom.HandlerFunc)
	// Setup metrics endpoint at another server
	prom.SetMetricsPath(echoPrometheus)

	go func() {
		echoPrometheus.Logger.Fatal(echoPrometheus.Start(":" + config.App().MetricsPort))

		// graceful shutdown prometheus server
		GracefulShutdown(echoPrometheus)
	}()

	container.Init(e.Group("api"))

	// start http server
	go func() {
		e.Logger.Fatal(e.Start(":" + config.App().Port))
	}()

	// graceful shutdown
	GracefulShutdown(e)
}

// server will gracefully shutdown within 5 sec
func GracefulShutdown(e *echo.Echo) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	logger.Info("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_ = e.Shutdown(ctx)
	logger.Info("server shutdowns gracefully")
}
