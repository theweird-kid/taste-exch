package dto

type FavouriteRequest struct {
	Status   string `json:"status" binding:"required,oneof=favourite unfavourite"`
	RecipeID int64  `json:"recipe_id" binding:"required"` // ID of the recipe
}
