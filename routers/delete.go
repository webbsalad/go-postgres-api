package routers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/webbsalad/go-postgres-api/db"
	"github.com/webbsalad/go-postgres-api/db/operations"
)

func DeleteItemRouter(dbConn *db.DBConnection) gin.HandlerFunc {
	return func(con *gin.Context) {
		tableName := con.Param("table_name")
		itemID, err := strconv.Atoi(con.Param("item_id"))
		if err != nil {
			con.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
			return
		}

		err = operations.DeleteItemByID(dbConn, tableName, itemID)
		if err != nil {
			con.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		con.String(http.StatusOK, "Item deleted successfully")
	}
}
