package api_view

import (
	"restaurant/api_controller/admin"

	"github.com/labstack/echo/v4"
)

func ComponentRoutes(server *echo.Echo) {
	components := server.Group("/components")
	// Overlay
	// overlayRoutes := components.Group("/overlay")
	// overlayRoutes.POST("/addValue", admin.GetUserList)
	// Admin
	adminRoutes := components.Group("/admin")
	adminRoutes.GET("/user_list", admin.GetUserList)

}
