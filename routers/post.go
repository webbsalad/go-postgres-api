package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/webbsalad/go-postgres-api/db"
	"github.com/webbsalad/go-postgres-api/db/operations"
)

func PostItemRouter(dbConn *db.DBConnection) gin.HandlerFunc {
	return func(c *gin.Context) {
		tableName := c.Param("table_name")
		var newItem map[string]interface{}

		if err := c.ShouldBindJSON(&newItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}

		err := operations.AddItem(dbConn, tableName, newItem)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, newItem)
	}
}
