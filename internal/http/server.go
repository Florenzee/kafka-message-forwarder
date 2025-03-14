package server

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"bitbucket.org/Amartha/go-x/log"
	"github.com/labstack/echo-contrib/pprof"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"bitbucket.org/Amartha/go-dlq-retrier/internal/handler"
)

type IHandler interface {
	Handle(c echo.Context) error
}

type server struct {
	echo *echo.Echo
	addr string
}

func New(
	handler handler.IMount[echo.Echo],
) *server {
	router := echo.New()
	router.HideBanner = true

	// options middleware
	router.Use(
		middleware.TimeoutWithConfig(
			middleware.TimeoutConfig{
				Timeout:      time.Second,
				ErrorMessage: "timeout",
				OnTimeoutRouteErrorHandler: func(err error, c echo.Context) {
					log.Error(c.Request().Context(), "request timeout", log.Err(err))
				},
				Skipper: func(c echo.Context) bool {
					return false
				},
			},
		),
		middleware.Recover(),
	)

	handler.Mount(router)

	pprof.Register(router)
	server := &server{
		echo: router,
		addr: fmt.Sprintf(":%d", 8080),
	}

	return server
}

func (s *server) Close(ctx context.Context) error {
	return s.echo.Shutdown(ctx)
}

func (s *server) Start(ctx context.Context) error {
	go func() {
		_ = s.echo.Start(s.addr)
	}()
	return nil
}

func (s *server) RoutesJSON() ([]byte, error) {
	return json.MarshalIndent(s.echo.Routes(), "", "  ")
}
