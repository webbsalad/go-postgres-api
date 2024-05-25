package operations

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/webbsalad/go-postgres-api/db"
)

func FetchDataAsJSON(dbConn *db.DBConnection, tableName string, filters map[string]string, sortBy string) (string, error) {
	whereClause := ""
	params := make([]interface{}, 0)

	if len(filters) > 0 {
		conditions := make([]string, 0)
		index := 1
		for key, value := range filters {
			if strings.Contains(value, "*") {
				value = strings.ReplaceAll(value, "*", "%")
				conditions = append(conditions, fmt.Sprintf(`"%s" ILIKE $%d`, key, index))
			} else {
				conditions = append(conditions, fmt.Sprintf(`"%s" = $%d`, key, index))
			}
			params = append(params, value)
			index++
		}
		whereClause = " WHERE " + strings.Join(conditions, " AND ")
	}

	validSortColumns := []string{"id", "name", "date"}
	isValidSortBy := false
	for _, col := range validSortColumns {
		if sortBy == col {
			isValidSortBy = true
			break
		}
	}
	if !isValidSortBy && sortBy != "" {
		return "", fmt.Errorf("invalid sortBy column: %s", sortBy)
	}

	orderClause := ""
	if sortBy != "" {
		orderClause = " ORDER BY " + sortBy
	}

	query := fmt.Sprintf(`SELECT * FROM "%s"%s%s`, tableName, whereClause, orderClause)
	rows, err := dbConn.Conn.Query(context.Background(), query, params...)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	columnNames := make([]string, 0)
	for _, fieldDescription := range rows.FieldDescriptions() {
		columnNames = append(columnNames, string(fieldDescription.Name))
	}

	data := make([]map[string]interface{}, 0)

	for rows.Next() {
		values := make([]interface{}, len(columnNames))
		valuePointers := make([]interface{}, len(columnNames))
		for i := range values {
			valuePointers[i] = &values[i]
		}
		if err := rows.Scan(valuePointers...); err != nil {
			return "", err
		}
		item := make(map[string]interface{})
		for i, col := range columnNames {
			item[col] = values[i]
		}
		data = append(data, item)
	}
	if err := rows.Err(); err != nil {
		return "", err
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return "[" + string(jsonData) + "]", nil
}
