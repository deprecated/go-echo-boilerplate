package routes

import (
	"github.com/deprecated/go-echo-boilerplate/api/admin"
	"github.com/deprecated/go-echo-boilerplate/api/user"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// InitRoutes initializes all of our api routes
func InitRoutes() *echo.Echo {
	e := echo.New()
	// Cache certificates
	/* e.Pre(middleware.HTTPSNonWWWRedirect()) */
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	  }), middleware.Recover()) 

	admin.Routes(e)
	user.Routes(e)

	

	return e
}