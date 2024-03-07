package view

import (
	"restaurant/controller/user"

	"github.com/labstack/echo/v4"
)

func AdminRoutes(server *echo.Echo) {
	router := server.Group("/api")
	// Overlay
	userRoutes := router.Group("/user")
	userRoutes.GET("/update", user.UpdateValue)
	// adminRoutes.GET("/add_value", admin.AddValue)

}
