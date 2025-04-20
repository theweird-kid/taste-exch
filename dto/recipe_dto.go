package dto

import (
	"strings"

	"github.com/theweird-kid/taste-exch/internal/database"
)

type NewRecipeRequest struct {
	UserID       int      `json:"user_id"`      // ID of the user creating the recipe
	Title        string   `json:"title"`        // Title of the recipe
	Description  string   `json:"description"`  // Description of the recipe
	Tags         []string `json:"tags"`         // Tags associated with the recipe
	Ingredients  []string `json:"ingredients"`  // List of ingredients
	Instructions []string `json:"instructions"` // Cooking instructions
	TotalTime    int      `json:"total_time"`   // Total time required (in minutes)
	Difficulty   string   `json:"difficulty"`   // Difficulty level (e.g., "easy", "medium", "hard")
	Servings     int      `json:"servings"`     // Number of servings
	PhotoURL     string   `json:"photo_url"`    // URL of the recipe photo
}

type MyRecipeResponse struct {
	ID           int    `json:"id"`           // ID of the recipe
	UserID       int    `json:"user_id"`      // ID of the user creating the recipe
	Title        string `json:"title"`        // Title of the recipe
	Description  string `json:"description"`  // Description of the recipe
	Tags         string `json:"tags"`         // Tags associated with the recipe
	Ingredients  string `json:"ingredients"`  // List of ingredients
	Instructions string `json:"instructions"` // Cooking instructions
	TotalTime    int    `json:"total_time"`   // Total time required (in minutes)
	Difficulty   string `json:"difficulty"`   // Difficulty level (e.g., "easy", "medium", "hard")
	Servings     int    `json:"servings"`     // Number of servings
	PhotoURL     string `json:"photo_url"`    // URL of the recipe photo
}

func RecipeResponseToDto(recipe database.Recipe) *MyRecipeResponse {

	return &MyRecipeResponse{
		ID:           int(recipe.ID),
		UserID:       int(recipe.UserID.Int32),
		Title:        recipe.Title,
		Description:  recipe.Description.String,
		Tags:         strings.Join(recipe.Tags, ", "),
		Ingredients:  strings.Join(recipe.Ingredients, ", "),
		Instructions: strings.Join(recipe.Instructions, ", "),
		TotalTime:    int(recipe.TotalTime.Int32),
		Difficulty:   recipe.Difficulty.String,
		Servings:     int(recipe.Servings.Int32),
		PhotoURL:     recipe.PhotoUrl.String,
	}
}

// Simplified response for a user's recipes
type RecipeResponse struct {
	ID          int    `json:"id"`          // ID of the recipe
	Title       string `json:"title"`       // Title of the recipe
	Description string `json:"description"` // Description of the recipe
	PhotoURL    string `json:"photo_url"`   // URL of the recipe photo
	Tags        string `json:"tags"`        // Tags associated with the recipe
}

// Converts a database row to MyRecipeResponse
func MyRecipeResponseFromRow(row database.GetRecipesByUserRow) RecipeResponse {
	return RecipeResponse{
		ID:          int(row.RecipeID),
		Title:       row.RecipeName,
		Description: row.RecipeDescription.String,
		PhotoURL:    row.PhotoUrl.String,
		Tags:        strings.Join(row.Tags, ", "),
	}
}

// Overload
func RecipeResponseFromRow(row database.GetRecipesRow) RecipeResponse {

	return RecipeResponse{
		ID:          int(row.RecipeID),
		Title:       row.RecipeName,
		Description: row.RecipeDescription.String,
		PhotoURL:    row.PhotoUrl.String,
		Tags:        strings.Join(row.Tags, ", "),
	}
}

// FavouriteRecipeResponseFromRow
func FavouriteRecipeResponseFromRow(rows []database.GetFavouriteRecipesRow) RecipesResponse {
	var resp RecipesResponse
	for _, row := range rows {

		resp.Recipes = append(resp.Recipes, RecipeResponse{
			ID:          int(row.RecipeID),
			Title:       row.RecipeName,
			Description: row.RecipeDescription.String,
			PhotoURL:    row.PhotoUrl.String,
			Tags:        strings.Join(row.Tags, ", "),
		})
	}

	return resp
}

// Wrapper for multiple recipes
type RecipesResponse struct {
	Recipes []RecipeResponse `json:"recipes"` // List of recipes
}
