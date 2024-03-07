package main

import (
	"encoding/json"
	"os"
	"restaurant/router"

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

	// main app
	e.Static("/public", "public")
	e.File("/", "public/index.html")

	// board
	// e.Static("/board", "board")
	// e.File("/boardPage", "board/index.html")
	// // board
	// e.Static("/board2", "board2")
	// e.File("/boardPage2", "board2/index.html")
	// e.GET("/table", func(c echo.Context) error {
	// 	table := board.GenBoard()
	// 	return c.HTML(200, table)
	// })

	router.InitRoutes(e)

	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		panic(err)
	}
	os.WriteFile("routes.json", data, 0644)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
