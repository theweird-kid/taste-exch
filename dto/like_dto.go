package dto

type LikeRequest struct {
	Status   string `json:"status" binding:"required,oneof=like unlike"` // "like" or "unlike"
	RecipeID int64  `json:"recipe_id" binding:"required"`                // ID of the recipe to like/unlike
}
