package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/theweird-kid/taste-exch/db"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		os.Exit(-1)
	}

	dsn := os.Getenv("DSN")

	// Server Instance
	r := gin.Default()

	//Database Instance
	_, err = db.NewDB(dsn)
	if err != nil {
		log.Fatal(err)
	}

	RegisterRoutes(r)
	r.Run("0.0.0.0:8080")
}

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
}
