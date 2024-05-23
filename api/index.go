package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/webbsalad/go-postgres-api/config"
	"github.com/webbsalad/go-postgres-api/db"
	"github.com/webbsalad/go-postgres-api/routers"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	rout := handler()

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

func handler() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	cfgDB, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Error reading environment variables: %v\n", err)
		return nil
	}

	database := db.DBConnection{Config: cfgDB}
	defer database.Close()

	if err := database.Connect(); err != nil {
		fmt.Printf("Error while connecting to PostgreSQL: %v\n", err)
		return nil
	}

	r := gin.New()

	r.GET("/:table_name/:item_id", routers.GetItemHandler(&database))
	r.GET("/:table_name/", routers.GetAllItemsHandler(&database))
	r.DELETE("/:table_name/:item_id", routers.DeleteItemHandler(&database))

	return r
}
