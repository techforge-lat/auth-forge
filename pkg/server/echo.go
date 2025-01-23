package server

import (
	"cloud-crm-backend/internal/shared/domain"
	"cloud-crm-backend/internal/shared/domain/ports/out"
	"cloud-crm-backend/internal/shared/middlewares"
	"cloud-crm-backend/pkg/database"
	"cloud-crm-backend/pkg/dependency"
	"cloud-crm-backend/pkg/telemetry"
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/techforge-lat/errortrace/v2"
	"github.com/techforge-lat/linkit"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
)

type Server struct {
	ServiceName string
	Echo        *echo.Echo
	Config      domain.Configuration
	Container   *linkit.DependencyContainer
	Logger      out.Logger
	Database    *database.Adapter
}

func New(container *linkit.DependencyContainer, serviceName string) (*Server, error) {
	config, err := linkit.Resolve[domain.Configuration](container, dependency.LocalConfig)
	if err != nil {
		return nil, errortrace.OnError(err)
	}

	logger, err := linkit.Resolve[out.Logger](container, dependency.Logger)
	if err != nil {
		return nil, errortrace.OnError(err)
	}

	database, err := linkit.Resolve[*database.Adapter](container, dependency.Database)
	if err != nil {
		return nil, errortrace.OnError(err)
	}

	e := echo.New()

	// NOTE: the open telemetry middleware must be the first middleware to be registered in the echo instance to ensure that all the requests are traced
	e.Use(otelecho.Middleware(serviceName))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())

	// CORS restricted
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: config.AllowedOrigins,
		AllowMethods: config.AllowedMethods,
	}))

	e.HTTPErrorHandler = middlewares.HTTPErrorHandler(logger)

	return &Server{
		ServiceName: serviceName,
		Echo:        e,
		Container:   container,
		Config:      config,
		Logger:      logger,
		Database:    database,
	}, nil
}

func (s Server) Start() error {
	// Configures the timezone for the hole application

	// loc, err := time.LoadLocation(s.Config.Timezone)
	// if err != nil {
	// 	return errortrace.OnError(err)
	// }
	//
	// time.Local = loc

	// Handle SIGINT (CTRL+C) gracefully
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	otelShutdown, err := telemetry.NewOpenTelemetry(s.ServiceName, s.Config.OTEL.CollectorEndpoint).Execute(ctx)
	if err != nil {
		return errortrace.OnError(err)
	}

	// Handle shutdown properly so nothing leaks.
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	s.Echo.Server.BaseContext = func(listener net.Listener) context.Context {
		return ctx
	}

	srvErr := make(chan error, 1)

	// Start server
	go func() {
		srvErr <- s.Echo.Start(fmt.Sprintf(":%d", s.Config.PortHTTP))
	}()

	// Wait for interruption.
	select {
	case err = <-srvErr:
		// Error when starting HTTP server.
		return err
	case <-ctx.Done():
		// Wait for first CTRL+C.
		// Stop receiving signal notifications as soon as possible.
		stop()
	}

	if err := s.Echo.Shutdown(context.Background()); err != nil {
		return errortrace.OnError(err)
	}

	return nil
}

func (s Server) HealthCheckController(c echo.Context) error {
	if err := s.Database.Ping(c.Request().Context()); err != nil {
		s.Logger.Error(c.Request().Context(), "err when pinging databases", "error", err.Error(), "server_time", time.Now(), "service_name", s.ServiceName)

		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":       "error",
			"error":        err.Error(),
			"server_time":  time.Now(),
			"service_name": s.ServiceName,
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"status":       "ok",
		"server_time":  time.Now(),
		"service_name": s.ServiceName,
	})
}
