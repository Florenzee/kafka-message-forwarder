package handler

import (
	"github.com/labstack/echo/v4"
)

type group struct {
	handlers []IMount[echo.Group]
}

func New(handlers ...IMount[echo.Group]) *group {

	return &group{
		handlers: handlers,
	}
}

func (g group) Mount(e *echo.Echo) {
	root := e.Group("/")
	for _, handler := range g.handlers {
		handler.Mount(root)
	}
}
