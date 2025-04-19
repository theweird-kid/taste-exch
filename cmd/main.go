package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/theweird-kid/taste-exch/db"
	"github.com/theweird-kid/taste-exch/handlers"
	"github.com/theweird-kid/taste-exch/internal/database"
	"github.com/theweird-kid/taste-exch/utils"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		os.Exit(-1)
	}

	// Get the DSN (Database Source Name) from the environment
	dsn := os.Getenv("DSN")

	// Initialize the database
	db, err := db.NewDB(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a new Gin server instance
	r := gin.Default()

	// Initialize the handler with the database
	handler := &handlers.Handler{
		Queries: database.New(db),
	}

	// Register routes
	RegisterRoutes(r, handler)

	// Start the server
	r.Run("0.0.0.0:8080")
}

func RegisterRoutes(r *gin.Engine, handler *handlers.Handler) {
	// Add a health check route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	// Group all routes under /api
	api := r.Group("/api")

	// Public Routes (No Authentication Required)
	api.POST("/signin", handler.SignIn)
	api.POST("/register", handler.Register)
	api.GET("/get_recipes/page/:page", handler.GetRecipes)
	api.GET("/get_recipe/id/:id", handler.GetRecipe)
	api.GET("/get_user/id/:id", handler.GetUser)

	// Authenticated Routes (Require Authentication)
	auth := api.Group("/")
	auth.Use(utils.AuthMiddleware())
	auth.POST("/update_user", handler.UpdateUser)
	auth.GET("/favourite_recipes", handler.GetFavouriteRecipes)
	auth.GET("/my_recipes", handler.GetMyRecipes)
	auth.POST("/favourite_recipe", handler.FavouriteRecipe)
	auth.POST("/new_recipe", handler.NewRecipe)
	auth.POST("/like_recipe", handler.LikeRecipe)
}
