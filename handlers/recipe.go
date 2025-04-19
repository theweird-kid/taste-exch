package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/theweird-kid/taste-exch/dto"
)

func (h *Handler) GetRecipes(c *gin.Context) {

}

func (h *Handler) GetRecipe(c *gin.Context) {
	recID, _ := strconv.Atoi(c.Param("id"))

	recipe, err := h.Queries.GetRecipeById(c, int32(recID))
	if err == sql.ErrNoRows {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid recipe id",
		})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server problem",
		})
		return
	} else {
		resp := dto.RecipeResponse{
			ID:           int(recipe.ID),
			UserID:       int(recipe.UserID.Int32),
			Title:        recipe.Title,
			Description:  recipe.Description.String,
			Tags:         recipe.Tags,
			Ingredients:  recipe.Ingredients,
			Instructions: recipe.Instructions,
			TotalTime:    int(recipe.TotalTime.Int32),
			Difficulty:   recipe.Difficulty.String,
			Servings:     int(recipe.Servings.Int32),
			PhotoURL:     recipe.PhotoUrl.String,
		}

		c.JSON(http.StatusOK, resp)
	}
}

func (h *Handler) GetFavouriteRecipes(c *gin.Context) {

}

func (h *Handler) GetMyRecipes(c *gin.Context) {

}
