package router

import (
	view "restaurant/view"

	"github.com/labstack/echo/v4"
)

func InitRoutes(server *echo.Echo) {

	view.ComponentRoutes(server)
	view.AdminRoutes(server)
	// view.UIRoutes(server)
}
