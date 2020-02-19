package main

import (
	"github.com/deprecated/go-echo-boilerplate/cmd/server/routes"
)

func main() {


	/* e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `
			<h1>Welcome to Echo!</h1>
			<h3>TLS certificates automatically installed from Let's Encrypt :)</h3>
		`)
	}) */

	e := routes.InitRoutes()


	/* e.Logger.Fatal(e.StartTLS(":443", "cert.pem", "key.pem")) */
	e.Logger.Fatal(e.Start(":8080"))
}