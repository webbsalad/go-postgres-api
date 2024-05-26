package routers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/webbsalad/go-postgres-api/db"
	"github.com/webbsalad/go-postgres-api/db/operations"
)

func PatchItemRouter(dbConn *db.DBConnection) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tableName := c.Params("table_name")
		itemID, err := strconv.Atoi(c.Params("item_id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid item ID"})
		}

		var newData map[string]interface{}
		if err := c.BodyParser(&newData); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
		}

		err = operations.UpdateItemStatus(dbConn, tableName, itemID, newData)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.SendString("Item updated successfully")
	}
}
