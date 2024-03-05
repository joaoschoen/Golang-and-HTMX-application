package api_view

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUser(echo echo.Context) error {
	// PARAM
	id := echo.Param("id")

	/**
		DATABASE REQUEST GOES HERE
	**/

	if id == "404" {
		return echo.JSON(http.StatusNotFound, "User not found.")
	}

	//BUILD RESPONSE

	return echo.JSON(http.StatusOK, "response")
}
