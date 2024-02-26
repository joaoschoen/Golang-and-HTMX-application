package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RemoveTrailingSlash())

	// router.InitRoutes(e)
	e.Static("/public", "public")
	e.File("/", "public/index.html")

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
