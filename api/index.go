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

	app.Get("/:table_name/:item_id", func(ctx *fiber.Ctx) error {
		defer database.Close()
		routers.GetItemRouter(&database)(ctx)
		return routers.GetItemRouter(&database)(ctx)
	})

	app.Get("/:table_name/", func(ctx *fiber.Ctx) error {
		defer database.Close()
		routers.GetAllItemsRouter(&database)(ctx)
		return routers.GetAllItemsRouter(&database)(ctx)
	})
	app.Post("/:table_name/", func(ctx *fiber.Ctx) error {
		defer database.Close()
		routers.PostItemRouter(&database)(ctx)
		return routers.PostItemRouter(&database)(ctx)
	})
	app.Patch("/:table_name/:item_id", func(ctx *fiber.Ctx) error {
		defer database.Close()
		routers.PatchItemRouter(&database)(ctx)
		return routers.PatchItemRouter(&database)(ctx)
	})
	app.Delete("/:table_name/:item_id", func(ctx *fiber.Ctx) error {
		defer database.Close()
		routers.DeleteItemRouter(&database)(ctx)
		return routers.DeleteItemRouter(&database)(ctx)
	})

	return adaptor.FiberApp(app)
}
