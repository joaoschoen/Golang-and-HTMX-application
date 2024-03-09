package admin

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

func GetUserList(context echo.Context) error {
	var table string
	// Read table head template
	table, err := tableHead()
	if err != nil {
		return context.String(http.StatusInternalServerError, "Error reading template")
	}

	// Open SQLite database
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		return context.String(http.StatusInternalServerError, "Error opening database")
	}
	defer db.Close()

	// Get user list information
	rows, err := db.Query("SELECT id,name,value FROM users")
	if err != nil {
		return context.String(http.StatusInternalServerError, "Error querying database")
	}
	defer rows.Close()

	// Read template file
	cwd, err := os.Getwd()
	if err != nil {
		return context.String(500, "Error while building table")
	}
	data, err := os.ReadFile(cwd + "/controller/admin/table/table_row.html")
	if err != nil {
		return context.String(500, "Error while building table")

	}
	template := string(data)
	// Collumns
	for rows.Next() {
		var id int16
		var name string
		var value float64

		err = rows.Scan(&id, &name, &value)

		if err != nil {
			context.String(500, "Error while building table")
		}
		newRow := tableRow(template, id, name, value)
		table += newRow
	}

	// Set Content-Type header
	context.Response().Header().Set(echo.HeaderContentType, echo.MIMETextPlain)
	// Send CSV data as text
	return context.HTML(http.StatusOK, table)
}

func tableHead() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return "", errors.New("internal server error")
	}
	data, err := os.ReadFile(cwd + "/controller/admin/table/table_head.html")
	if err != nil {
		println(err.Error())
		return "", errors.New("internal server error")

	}
	return string(data), nil
}

func tableRow(template string, id int16, name string, value float64) string {
	newRow := template
	newRow = strings.Replace(newRow, "${id}", strconv.FormatInt(int64(id), 10), -1)
	newRow = strings.Replace(newRow, "${name}", name, -1)
	newRow = strings.Replace(newRow, "${value}", strconv.FormatFloat(value, 'f', 2, 32), -1)

	return newRow
}
