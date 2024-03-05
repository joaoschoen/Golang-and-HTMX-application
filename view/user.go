package api_view

import (
	"restaurant/controller/admin"

	"github.com/labstack/echo/v4"
)

func AdminRoutes(server *echo.Echo) {
	router := server.Group("/api")
	// Overlay
	userRoutes := router.Group("/user")
	userRoutes.GET("/update", admin.UpdateValue)
	// adminRoutes.GET("/add_value", admin.AddValue)

}
