package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/theweird-kid/taste-exch/dto"
)

func (h *Handler) GetRecipes(c *gin.Context) {
	pageNo, _ := strconv.Atoi(c.Param("page"))

	recipes, err := h.Queries.GetRecipes(c, int32(10*pageNo))
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNoContent, gin.H{
			"message": "no more pages",
		})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server problem",
		})
		return
	}

	// Convert database rows to DTO
	var resp dto.RecipesResponse
	resp.Recipes = make([]dto.RecipeResponse, 0, len(recipes))
	for _, recipe := range recipes {
		resp.Recipes = append(resp.Recipes, dto.RecipeResponseFromRow(recipe))
	}

	// Return the response
	c.JSON(http.StatusOK, resp)
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

	recipes, err := h.Queries.GetFavouriteRecipes(c, sql.NullInt32{Int32: int32(uid), Valid: true})
	if err == sql.ErrNoRows {
		// No recipes found
		c.JSON(http.StatusOK, dto.RecipesResponse{Recipes: []dto.RecipeResponse{}})
		return
	} else if err != nil {
		// Internal server error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Problem"})
		return
	}

	resp := dto.FavouriteRecipeResponseFromRow(recipes)
	c.JSON(http.StatusOK, resp)
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

	recipes, err := h.Queries.GetRecipesByUser(c, sql.NullInt32{Int32: int32(uid), Valid: true})
	if err == sql.ErrNoRows {
		// No recipes found
		c.JSON(http.StatusOK, dto.RecipesResponse{Recipes: []dto.RecipeResponse{}})
		return
	} else if err != nil {
		// Internal server error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Problem"})
		return
	}

	// Convert database rows to DTO
	var resp dto.RecipesResponse
	resp.Recipes = make([]dto.RecipeResponse, 0, len(recipes))
	for _, recipe := range recipes {
		resp.Recipes = append(resp.Recipes, dto.MyRecipeResponseFromRow(recipe))
	}

	// Return the response
	c.JSON(http.StatusOK, resp)
}
