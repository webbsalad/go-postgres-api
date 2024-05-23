package operations

import (
	"context"
	"fmt"
	"strings"

	"github.com/webbsalad/go-postgres-api/db"
)

func UpdateItemStatus(dbConn *db.DBConnection, tableName string, itemID int, status map[string]interface{}) error {
	var setStatements []string
	var values []interface{}
	i := 1
	for key, value := range status {
		setStatements = append(setStatements, fmt.Sprintf("%s = $%d", key, i))
		values = append(values, value)
		i++
	}

	query := fmt.Sprintf(`UPDATE %s SET %s WHERE id = $%d`, tableName,
		strings.Join(setStatements, ", "), len(values)+1)
	values = append(values, itemID)

	_, err := dbConn.Conn.Exec(context.Background(), query, values...)
	if err != nil {
		return err
	}
	return nil
}
