package utils

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/webbsalad/go-postgres-api/db"
)

func GetMaxID(dbConn *db.DBConnection, tableName string) (int, error) {
	var maxID int
	err := dbConn.Conn.QueryRow(context.Background(), fmt.Sprintf("SELECT COALESCE(MAX(id), 0) FROM %s", tableName)).Scan(&maxID)
	if err != nil {
		return 0, err
	}
	return maxID, nil
}

type ArrayType []interface{}

func (a ArrayType) Value() (driver.Value, error) {
	jsonArray, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}
	return string(jsonArray), nil
}
