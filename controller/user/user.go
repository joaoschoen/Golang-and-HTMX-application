package user

import "github.com/labstack/echo/v4"

func UpdateValue(context echo.Context) error {

	return context.HTML(200, "!")
}
