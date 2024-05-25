package operations

import (
	"context"
	"fmt"
	"strconv"

	"github.com/webbsalad/go-postgres-api/db"
	"github.com/webbsalad/go-postgres-api/utils"
)

func AddItem(dbConn *db.DBConnection, tableName string, newItem map[string]interface{}) error {
	maxID, err := utils.GetMaxID(dbConn, tableName)
	if err != nil {
		return err
	}

	newItem["id"] = maxID + 1

	columns := ""
	values := ""
	var args []interface{}

	index := 0
	for key, value := range newItem {
		if index != 0 {
			columns += ", "
			values += ", "
		}
		columns += key
		values += "$" + strconv.Itoa(index+1)
		args = append(args, value)
		index++
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, columns, values)
	_, err = dbConn.Conn.Exec(context.Background(), query, args...)
	if err != nil {
		return err
	}

	return nil
}
