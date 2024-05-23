package routers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/webbsalad/go-postgres-api/db"
	"github.com/webbsalad/go-postgres-api/db/operations"
)

func GetItemHandler(dbConn *db.DBConnection) gin.HandlerFunc {
	return func(c *gin.Context) {
		tableName := c.Param("table_name")

		var filters map[string]string
		if itemIDStr := c.Param("item_id"); itemIDStr != "" {
			itemID, err := strconv.Atoi(itemIDStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
				return
			}
			filters = map[string]string{"id": strconv.Itoa(itemID)}
		}

		itemsJSON, err := operations.FetchDataAsJSON(dbConn, tableName, filters, "")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if itemsJSON == "[]" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
			return
		}

		c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(itemsJSON[1:len(itemsJSON)-1]))
	}
}

func GetAllItemsHandler(dbConn *db.DBConnection) gin.HandlerFunc {
	return func(c *gin.Context) {
		tableName := c.Param("table_name")

		itemsJSON, err := operations.FetchDataAsJSON(dbConn, tableName, nil, "")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if itemsJSON == "[]" {
			c.JSON(http.StatusNotFound, gin.H{"error": "No items found"})
			return
		}

		c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(itemsJSON[1:len(itemsJSON)-1]))
	}
}
