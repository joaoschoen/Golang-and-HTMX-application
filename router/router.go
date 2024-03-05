package router

import (
	"restaurant/api_view"

	"github.com/labstack/echo/v4"
)

func InitRoutes(server *echo.Echo) {

	api_view.ComponentRoutes(server)
	// view.UIRoutes(server)
}
