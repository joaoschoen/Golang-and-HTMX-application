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
		var value float64

		err = rows.Scan(&id, &name, &value)

		if err != nil {
			context.String(500, "Error while building table")
		}
		newRow, err := tableRow(id, name, value)
		if err != nil {
			context.String(500, "Error while building table")
		}
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

func tableRow(id int16, name string, value float64) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return "", errors.New("internal server error")
	}
	data, err := os.ReadFile(cwd + "/controller/admin/table/table_row.html")
	if err != nil {
		println(err.Error())
		return "", errors.New("internal server error")

	}
	row := string(data)
	row = strings.Replace(row, "${id}", strconv.FormatInt(int64(id), 10), -1)
	row = strings.Replace(row, "${name}", name, 1)
	row = strings.Replace(row, "${value}", strconv.FormatFloat(value, 'f', -1, 32), -1)

	return row, nil
}
