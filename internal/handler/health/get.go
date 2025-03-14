package health

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Get struct{}

func NewGet() *Get {
	return &Get{}
}

func (get *Get) Handle(c echo.Context) error {
	type responseData struct {
		ServiceTime time.Time `json:"service_time"`
	}

	return c.JSON(http.StatusOK, responseData{ServiceTime: time.Now()})
}

func (get *Get) Mount(e *echo.Group) {
	e.GET("health", get.Handle)
}
