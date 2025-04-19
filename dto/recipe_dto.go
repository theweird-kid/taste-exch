package dto

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

type RecipeResponse struct {
	ID           int      `json:id`             // ID of the recipe
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
