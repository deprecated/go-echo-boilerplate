package user

import (
	"net/http"
	"github.com/labstack/echo"
)

func Routes(e *echo.Echo) *echo.Group {
	user := e.Group("/api")
	user.GET("/v1/user", func (c echo.Context) error {
		return c.String(http.StatusOK, "user route")

	})

	return user
}