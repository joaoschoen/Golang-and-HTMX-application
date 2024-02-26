package router

import (
	"restaurant/view"

	"github.com/labstack/echo/v4"
)

func InitRoutes(server *echo.Echo) {

	// view.UserRoutes(server)
	view.UIRoutes(server)
}
