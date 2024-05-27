package handler

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/webbsalad/go-postgres-api/config"
	"github.com/webbsalad/go-postgres-api/db"
	"github.com/webbsalad/go-postgres-api/routers"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	r.RequestURI = r.URL.String()
	createApp().ServeHTTP(w, r)
}

func createApp() http.HandlerFunc {
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

	app.Use(func(c *fiber.Ctx) error {
		origin := c.Get("Origin")
		allowedOrigins := []string{"https://db-flask-test.vercel.app", "http://127.0.0.1:8080"}

		if contains(allowedOrigins, origin) {
			c.Set("Access-Control-Allow-Origin", origin)
		} else {
			c.Set("Access-Control-Allow-Origin", "*")
		}
		c.Set("Access-Control-Allow-Methods", "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Origin,Content-Type,Accept,Authorization")
		if c.Method() == fiber.MethodOptions {
			return c.SendStatus(fiber.StatusNoContent)
		}
		return c.Next()
	})

	app.Get("/:table_name/:item_id", func(ctx *fiber.Ctx) error {
		return routers.GetItemRouter(&database)(ctx)
	})

	app.Get("/:table_name/", func(ctx *fiber.Ctx) error {
		return routers.GetAllItemsRouter(&database)(ctx)
	})
	app.Post("/:table_name/", func(ctx *fiber.Ctx) error {
		return routers.PostItemRouter(&database)(ctx)
	})
	app.Patch("/:table_name/:item_id", func(ctx *fiber.Ctx) error {
		return routers.PatchItemRouter(&database)(ctx)
	})
	app.Delete("/:table_name/:item_id", func(ctx *fiber.Ctx) error {
		return routers.DeleteItemRouter(&database)(ctx)
	})

	return adaptor.FiberApp(app)
}

func contains(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}
