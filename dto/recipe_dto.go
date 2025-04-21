package dto

import (
	"encoding/json"
	"strings"

	"github.com/theweird-kid/taste-exch/internal/database"
)

// Struct for creating a new recipe
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

// Struct for returning a detailed recipe response
type MyRecipeResponse struct {
	ID           int      `json:"id"`           // ID of the recipe
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

// RecipeResponseToDto converts a database.Recipe object to a MyRecipeResponse DTO.
// It decodes the inner JSON arrays for Tags, Ingredients and Instructions if needed.
func RecipeResponseToDto(recipe database.Recipe) *MyRecipeResponse {
	return &MyRecipeResponse{
		ID:           int(recipe.ID),
		UserID:       int(recipe.UserID.Int32),
		Title:        recipe.Title,
		Description:  recipe.Description.String,
		Tags:         decodeJSONFromSlice(recipe.Tags),
		Ingredients:  decodeJSONFromSlice(recipe.Ingredients),
		Instructions: decodeJSONFromSlice(recipe.Instructions),
		TotalTime:    int(recipe.TotalTime.Int32),
		Difficulty:   recipe.Difficulty.String,
		Servings:     int(recipe.Servings.Int32),
		PhotoURL:     recipe.PhotoUrl.String,
	}
}

// decodeJSONFromSlice takes a slice of strings which may contain JSON-encoded arrays.
// It decodes each JSON array string (if detected) and appends the resulting values to the result.
// Example: [ "[\"breakfast\", \"easy\"]" ] becomes []string{"breakfast", "easy"}.
func decodeJSONFromSlice(slice []string) []string {
	var result []string
	for _, s := range slice {
		// Trim any trailing commas and spaces.
		sTrim := strings.TrimSpace(strings.TrimSuffix(s, ","))
		if strings.HasPrefix(sTrim, "[") && strings.HasSuffix(sTrim, "]") {
			var decoded []string
			if err := json.Unmarshal([]byte(sTrim), &decoded); err == nil {
				result = append(result, decoded...)
				continue
			}
			// If decoding fails, fall back to the original string.
		}
		result = append(result, s)
	}
	return result
}

// Simplified response for a user's recipes
type RecipeResponse struct {
	ID          int    `json:"id"`          // ID of the recipe
	Title       string `json:"title"`       // Title of the recipe
	Description string `json:"description"` // Description of the recipe
	PhotoURL    string `json:"photo_url"`   // URL of the recipe photo
	Tags        string `json:"tags"`        // Tags associated with the recipe (comma-separated)
}

// MyRecipeResponseFromRow converts a database row to RecipeResponse for a user's recipes.
func MyRecipeResponseFromRow(row database.GetRecipesByUserRow) RecipeResponse {
	return RecipeResponse{
		ID:          int(row.RecipeID),
		Title:       row.RecipeName,
		Description: row.RecipeDescription.String,
		PhotoURL:    row.PhotoUrl.String,
		Tags:        strings.Join(row.Tags, ", "),
	}
}

// RecipeResponseFromRow converts a database row to RecipeResponse for general recipes.
func RecipeResponseFromRow(row database.GetRecipesRow) RecipeResponse {
	return RecipeResponse{
		ID:          int(row.RecipeID),
		Title:       row.RecipeName,
		Description: row.RecipeDescription.String,
		PhotoURL:    row.PhotoUrl.String,
		Tags:        strings.Join(row.Tags, ", "),
	}
}

// FavouriteRecipeResponseFromRow converts a list of favorite recipes to a RecipesResponse.
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

// RecipesResponse is a wrapper for multiple recipes.
type RecipesResponse struct {
	Recipes []RecipeResponse `json:"recipes"` // List of recipes
}
