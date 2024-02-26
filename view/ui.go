package view

import (
	"github.com/labstack/echo/v4"
)

func UIRoutes(server *echo.Echo) {
	group := server.Group("/public/ui")

	group.File("/header", "public/ui/header.html")
}
