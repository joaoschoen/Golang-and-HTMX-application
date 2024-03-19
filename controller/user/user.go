package user

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

type Update struct {
	id       string `form:"id"`
	newValue string `form:"newValue"`
}

type UserValue struct {
	Id    string `json:"id"`
	Value string `json:"value"`
}

func UpdateValue(context echo.Context) error {
	// Request Data
	formData := Update{
		id:       context.FormValue("id"),
		newValue: context.FormValue("newValue"),
	}
	additionalValue, err := strconv.ParseFloat(formData.newValue, 64)
	if err != nil {
		return context.String(http.StatusInternalServerError, "Error parsing float")
	}
	if len(formData.id) == 0 {
		return context.String(http.StatusInternalServerError, "Id in form data is empty")
	}
	// Open SQLite database
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		return context.String(http.StatusInternalServerError, "Error opening database")
	}
	defer db.Close()

	// Query user
	readQuery := fmt.Sprintf(`SELECT id,value FROM users WHERE id="%s" LIMIT 1;`, formData.id)
	findResult, err := db.Query(readQuery)
	if err != nil {
		return context.String(http.StatusInternalServerError, "Error querying database")
	}

	var user UserValue
	findResult.Next()
	err = findResult.Scan(&user.Id, &user.Value)
	if err != nil {
		return context.String(500, "Error while reading user data")
	}
	findResult.Close()

	currentValue, err := strconv.ParseFloat(user.Value, 64)
	if err != nil {
		return context.String(http.StatusInternalServerError, "Error parsing float")
	}
	updatedValue := currentValue + additionalValue
	updateQuery := fmt.Sprintf(`UPDATE users SET value = %.2f WHERE Id=%s;`, updatedValue, user.Id)
	println(updateQuery)
	updateResult, err := db.Exec(updateQuery)
	if err != nil {
		return context.String(http.StatusInternalServerError, err.Error())
	}
	affected, err := updateResult.RowsAffected()
	if err != nil {
		return context.String(500, "Error while updating user data")
	}
	if affected == 1 {
		// Read template file
		cwd, err := os.Getwd()
		if err != nil {
			return context.String(500, "Error while building table")
		}
		data, err := os.ReadFile(cwd + "/controller/admin/update_user_modal/table_row.html")
		if err != nil {
			return context.String(500, "Error while building table")

		}
		return context.String(http.StatusOK, "updated")
	}

	return context.String(500, "Error while reading user data")
}
