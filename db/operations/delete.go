package operations

import (
	"context"
	"fmt"

	"github.com/webbsalad/go-postgres-api/db"
)

func DeleteItemByID(dbConn *db.DBConnection, tableName string, itemID int) error {
	_, err := dbConn.Conn.Exec(context.Background(), fmt.Sprintf(`DELETE FROM "%s" WHERE id = $1`, tableName), itemID)
	if err != nil {
		return err
	}
	return nil
}
