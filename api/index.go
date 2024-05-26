package handler

import (
	"log"
	"net/http"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/webbsalad/go-postgres-api/config"
	"github.com/webbsalad/go-postgres-api/db"
	"github.com/webbsalad/go-postgres-api/routers"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	app := createApp()
	if app == nil {
		http.Error(w, "Failed to create Fiber app", http.StatusInternalServerError)
		return
	}

	adaptor.FiberApp(app).ServeHTTP(w, r)
}

func createApp() *fiber.App {
	cfgDB, err := config.LoadConfig()
	if err != nil {
		log.Printf("Ошибка при чтении переменных окружения: %v\n", err)
		return nil
	}

	database := db.DBConnection{Config: cfgDB}

	if err := database.Connect(); err != nil {
		log.Printf("Ошибка при подключении к PostgreSQL: %v\n", err)
		return nil
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST, PUT, PATCH, DELETE, OPTIONS",
		AllowHeaders: "*",
	}))

	app.All("/*", func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Set("Access-Control-Allow-Headers", "*")
		return c.SendStatus(fiber.StatusNoContent)
	})

	app.Use(customHeadersMiddleware())

	app.Get("/:table_name/:item_id", func(c *fiber.Ctx) error {
		defer database.Close()
		return routers.GetItemRouter(&database)(c)
	})
	app.Get("/:table_name/", func(c *fiber.Ctx) error {
		defer database.Close()
		return routers.GetAllItemsRouter(&database)(c)
	})
	app.Post("/:table_name/", func(c *fiber.Ctx) error {
		defer database.Close()
		return routers.PostItemRouter(&database)(c)
	})
	app.Patch("/:table_name/:item_id", func(c *fiber.Ctx) error {
		defer database.Close()
		return routers.PatchItemRouter(&database)(c)
	})
	app.Delete("/:table_name/:item_id", func(c *fiber.Ctx) error {
		defer database.Close()
		return routers.DeleteItemRouter(&database)(c)
	})

	return app
}

func customHeadersMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Set("Access-Control-Allow-Headers", "*")
		return c.Next()
	}
}
