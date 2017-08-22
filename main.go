package main

import (
	"github.com/david-torres/webview-app/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	middleware "github.com/labstack/echo/middleware"
	"github.com/zserge/webview"
)

func main() {
	// init the web server
	e := echo.New()

	// init app-wide middleware
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())

	// security measures
	e.Use(middleware.BodyLimit("2M")) // sets maximum request body size
	e.Use(middleware.CSRF())          // default protection against session riding
	e.Use(middleware.Secure())        // default protection against XSS, content sniffing, clickjacking, and other infections

	// init static assets
	e.Static("/", "public")

	// routes
	e.File("/", "public/index.html")
	e.GET("/ws", standard.WrapHandler(controllers.Socket()))

	go func() {
		// start the server
		e.Run(standard.New(":3000"))
	}()

	// start webview UI
	webview.Open("Hello World", "http://localhost:3000", 400, 300, true)
}
