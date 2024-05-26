package routers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/webbsalad/go-postgres-api/db"
	"github.com/webbsalad/go-postgres-api/db/operations"
)

func GetItemRouter(dbConn *db.DBConnection) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tableName := c.Params("table_name")

		filters := make(map[string]string)
		if itemIDStr := c.Params("item_id"); itemIDStr != "" {
			itemID, err := strconv.Atoi(itemIDStr)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid item ID"})
			}
			filters["id"] = strconv.Itoa(itemID)
		}

		c.QueryParser(&filters)

		sortBy := c.Query("sortBy", "")

		itemsJSON, err := operations.FetchDataAsJSON(dbConn, tableName, filters, sortBy)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		if itemsJSON == "[]" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Item not found"})
		}

		return c.Status(fiber.StatusOK).SendString(itemsJSON[1 : len(itemsJSON)-1])
	}
}

func GetAllItemsRouter(dbConn *db.DBConnection) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tableName := c.Params("table_name")

		filters := make(map[string]string)
		c.QueryParser(&filters)

		sortBy := c.Query("sortBy", "")

		itemsJSON, err := operations.FetchDataAsJSON(dbConn, tableName, filters, sortBy)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		if itemsJSON == "[]" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "No items found"})
		}

		return c.Status(fiber.StatusOK).SendString(itemsJSON[1 : len(itemsJSON)-1])
	}
}
