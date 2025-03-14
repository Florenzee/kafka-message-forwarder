package handler

import "github.com/labstack/echo/v4"

type IMount[T echo.Group | echo.Echo] interface {
	Mount(*T)
}
