package admin

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id    string `form:"id"`
	Name  string `form:"name"`
	Value string `form:"value"`
}

func AddValue(context echo.Context) error {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return context.String(500, "internal server error")
	}
	data, err := os.ReadFile(cwd + "/api_controller/admin/add_value.html")
	if err != nil {
		println(err.Error())
		return context.String(500, "internal server error")
	}
	content := string(data)
	var user User
	if err := context.Bind(&user); err != nil {
		return context.JSON(http.StatusBadRequest, "Error while parsing received data")
	}
	content = strings.Replace(content, "${name}", user.Name, 1)
	content = strings.Replace(content, "${value}", user.Value, 1)

	return context.HTML(200, content)
}
