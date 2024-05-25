package handler

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/webbsalad/go-postgres-api/config"
	"github.com/webbsalad/go-postgres-api/db"
	"github.com/webbsalad/go-postgres-api/routers"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	rout := createRouter()
	if rout == nil {
		http.Error(w, "Failed to create router", http.StatusInternalServerError)
		return
	}

	ginWriter := &ginResponseWriter{
		ResponseWriter: w,
	}

	rout.ServeHTTP(ginWriter, r)
}

type ginResponseWriter struct {
	http.ResponseWriter
}

func (w *ginResponseWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *ginResponseWriter) Write(data []byte) (int, error) {
	return w.ResponseWriter.Write(data)
}

func createRouter() *gin.Engine {
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

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.Use(customHeadersMiddleware())

	r.GET("/:table_name/:item_id", func(c *gin.Context) {
		defer database.Close()
		routers.GetItemRouter(&database)(c)
	})
	r.GET("/:table_name/", func(c *gin.Context) {
		defer database.Close()
		routers.GetAllItemsRouter(&database)(c)
	})
	r.POST("/:table_name/", func(c *gin.Context) {
		defer database.Close()
		routers.PostItemRouter(&database)(c)
	})
	r.PATCH("/:table_name/:item_id", func(c *gin.Context) {
		defer database.Close()
		routers.PatchItemRouter(&database)(c)
	})
	r.DELETE("/:table_name/:item_id", func(c *gin.Context) {
		defer database.Close()
		routers.DeleteItemRouter(&database)(c)
	})

	return r
}

func customHeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
		c.Next()
	}
}
