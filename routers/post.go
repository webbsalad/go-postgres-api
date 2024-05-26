package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/webbsalad/go-postgres-api/db"
	"github.com/webbsalad/go-postgres-api/db/operations"
)

func PostItemRouter(dbConn *db.DBConnection) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tableName := c.Params("table_name")
		var newItem map[string]interface{}

		if err := c.BodyParser(&newItem); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
		}

		err := operations.AddItem(dbConn, tableName, newItem)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(fiber.StatusOK).JSON(newItem)
	}
}
