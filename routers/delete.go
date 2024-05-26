package routers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/webbsalad/go-postgres-api/db"
	"github.com/webbsalad/go-postgres-api/db/operations"
)

func DeleteItemRouter(dbConn *db.DBConnection) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tableName := c.Params("table_name")
		itemID, err := strconv.Atoi(c.Params("item_id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid item ID"})
		}

		err = operations.DeleteItemByID(dbConn, tableName, itemID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.SendString("Item deleted successfully")
	}
}
