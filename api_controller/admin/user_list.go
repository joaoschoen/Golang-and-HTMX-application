package admin

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

func GetUserList(context echo.Context) error {
	println("??")
	var tableRows string = tableHead()

	// Open SQLite database
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		return context.String(http.StatusInternalServerError, "Error opening database")
	}
	defer db.Close()

	// Perform a query
	rows, err := db.Query("SELECT id,name,value FROM users")
	if err != nil {
		return context.String(http.StatusInternalServerError, "Error querying database")
	}
	defer rows.Close()

	// Collumns
	for rows.Next() {
		var id int16
		var name string
		var value float32

		err = rows.Scan(&id, &name, &value)

		if err != nil {
			log.Fatal(err)
		}
		tableRows += tableRow(id, name, value)
	}

	// Set Content-Type header
	context.Response().Header().Set(echo.HeaderContentType, echo.MIMETextPlain)
	// Send CSV data as text
	return context.HTML(http.StatusOK, tableRows)
}

func tableHead() string {
	tableHead := "<tr>\n" +
		"<th class=\"w-[10%]\">Id</th>\n" +
		"<th class=\"w-[50%]\">Nome</th>\n" +
		"<th class=\"w-[30%]\">Valor</th>\n" +
		"<th class=\"w-[10%]\">Ações</th>\n" +
		"</tr>\n"
	return tableHead
}

func tableRow(id int16, name string, value float32) string {
	tableHead :=
		"<tr>\n" +
			fmt.Sprintf("<td>%d</td>\n", id) +
			fmt.Sprintf("<td>%s</td>\n", name) +
			fmt.Sprintf("<td>%.2f</td>\n", value) +
			`<td>
				<button 
					hx-post="/components/overlay/addValue"
					hx-target="#overlay"
					hx-swap="innerHTML" ` +
			fmt.Sprintf(`hx-vals='{"id":%d}'`, id) +
			`class="w-32 h-32 rounded-full active:bg-slate-500 hover:bg-slate-400 bg-slate-300 text-black"
				>
					+
				</button>
			</td>
		</tr>`
	return tableHead
}
