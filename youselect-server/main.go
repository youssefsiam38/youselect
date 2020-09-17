package main

import (
	"github.com/youssefsiam38/youselect/handlers"
	"github.com/gin-contrib/cors"
	_ "github.com/joho/godotenv/autoload"
	"github.com/youssefsiam38/youselect/db"
	"github.com/youssefsiam38/youselect/middlewares"
	// "net/http"
	// "time"
	"os"
	// "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func main() {

	db.Setup()
	// gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("FE_DOMAIN")}, ///// env
		AllowMethods:     []string{"POST", "DELETE", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	r.GET("/categories", handlers.GetCategories)

	r.GET("/stores-names", handlers.GetStoresNames)

	r.GET("/all-products", handlers.AllProducts)

	r.GET("/products", handlers.Products)

	r.GET("/search", handlers.Search)

	r.GET("/login", handlers.Login)
	r.GET("/auth", middlewares.Auth)
	
	r.Run(`:` + os.Getenv("PORT"))

}
