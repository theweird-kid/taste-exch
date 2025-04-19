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
		resp := dto.RecipeResponseToDto(recipe)

		c.JSON(http.StatusOK, resp)
	}
}

func (h *Handler) GetFavouriteRecipes(c *gin.Context) {

}

func (h *Handler) GetMyRecipes(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	// Type Assert userID
	uid, ok := userID.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user ID"})
		return
	}

	recipes, err := h.Queries.GetRecipesByUser(c, sql.NullInt32{Int32: int32(uid)})
	if err == sql.ErrNoRows {
		c.JSON(http.StatusOK, gin.H{"message": "No recipes"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Problem"})
		return
	} else {
		var resp dto.RecipesResponse
		resp.Recipes = make([]dto.RecipeResponse, 0)
		// for _, recipe := range recipes {
		// 	resp.Recipes = append(resp.Recipes, *dto.RecipeResponseToDto(recipe))
		// }

		c.JSON(http.StatusOK, resp)
	}
}
