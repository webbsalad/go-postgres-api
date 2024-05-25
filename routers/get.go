package routers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/webbsalad/go-postgres-api/db"
	"github.com/webbsalad/go-postgres-api/db/operations"
)

func GetItemRouter(dbConn *db.DBConnection) gin.HandlerFunc {
	return func(c *gin.Context) {
		tableName := c.Param("table_name")

		filters := make(map[string]string)
		if itemIDStr := c.Param("item_id"); itemIDStr != "" {
			itemID, err := strconv.Atoi(itemIDStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
				return
			}
			filters["id"] = strconv.Itoa(itemID)
		}

		for key, value := range c.Request.URL.Query() {
			if key != "sortBy" && len(value) > 0 {
				filters[key] = value[0]
			}
		}

		sortBy := c.DefaultQuery("sortBy", "")

		itemsJSON, err := operations.FetchDataAsJSON(dbConn, tableName, filters, sortBy)
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

func GetAllItemsRouter(dbConn *db.DBConnection) gin.HandlerFunc {
	return func(c *gin.Context) {
		tableName := c.Param("table_name")

		filters := make(map[string]string)
		for key, value := range c.Request.URL.Query() {
			if key != "sortBy" && len(value) > 0 {
				filters[key] = value[0]
			}
		}

		sortBy := c.DefaultQuery("sortBy", "")

		itemsJSON, err := operations.FetchDataAsJSON(dbConn, tableName, filters, sortBy)
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
