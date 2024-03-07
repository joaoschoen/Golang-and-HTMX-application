package view

import (
	"restaurant/controller/admin"

	"github.com/labstack/echo/v4"
)

func ComponentRoutes(server *echo.Echo) {
	components := server.Group("/components")
	// Overlay
	overlayRoutes := components.Group("/overlay")
	overlayRoutes.POST("/add_value", admin.AddValue)
	// Admin
	adminRoutes := components.Group("/admin")
	adminRoutes.GET("/user_list", admin.GetUserList)
	// adminRoutes.GET("/add_value", admin.AddValue)

}
