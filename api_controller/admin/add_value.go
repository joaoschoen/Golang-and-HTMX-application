package admin

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
)

type User struct {
	id    int16
	name  string
	value float32
}

func AddValue(context echo.Context) error {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return context.String(500, "internal server error")
	}
	content, err := os.ReadFile(cwd + "/public/overlay/add_value.html")
	if err != nil {
		println(err.Error())
		return context.String(500, "internal server error")
	}
	user := new(User)
	println(user.id)
	println(user.name)
	println(user.value)
	if err := context.Bind(user); err != nil {
		return context.String(500, "internal server error")
	}
	return context.HTML(200, string(content))

	// overlay := `<div id="overlay" class="flex flex-col absolute -z-10 items-center justify-center w-[100vw] h-[100vh]">`
	// wrapper := "<div w-1/5 h-/5 bg->"
	// form := "<form>" +
	// 	"<div>Nome:${name}</div>" +
	// 	"<div>Valor:${value}</div>" +
	// 	`<input type="number" placeholder="Valor">` +
	// 	"</form>"
	// closeDiv := "</div>"
	// component := overlay + wrapper + form + closeDiv + closeDiv
	// return context.HTML(200, component)
}
