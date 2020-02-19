package admin

import (
	"net/http"
	"github.com/labstack/echo"
)

type User struct {
	Name  string `json:"name" xml:"name"`
	Email string `json:"email" xml:"email"`
  }
func login(c echo.Context) (err error) {
	u := new(User)
	if err = c.Bind(u); err != nil {
		response := make(map[string]string)
		response["message"] = "Bad Request"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	return c.JSON(http.StatusOK, u)
}

func Routes(e *echo.Echo) *echo.Group {
	admin := e.Group("/api")
	admin.POST("/v1/admin/login", login)

	admin.GET("/v1/admin/register", register)

	return admin
}

func register(c echo.Context) error {
	return c.String(http.StatusOK, "admin register")

}