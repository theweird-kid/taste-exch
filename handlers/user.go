package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/theweird-kid/taste-exch/dto"
	"github.com/theweird-kid/taste-exch/internal/database"
	"github.com/theweird-kid/taste-exch/utils"
)

type Handler struct {
	Queries *database.Queries
}

func (h *Handler) GetUser(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Param("id"))

	usrRow, err := h.Queries.GetUserByID(c, int32(uid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp := dto.UserResponse{
		ID:         int(usrRow.ID),
		Name:       usrRow.Name,
		Email:      usrRow.Email,
		ProfileURL: usrRow.ProfileUrl.String,
		CreatedAt:  usrRow.CreatedAt,
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) SignIn(c *gin.Context) {
	var req dto.SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	log.Println(req)

	// Check User
	user, err := h.Queries.GetUserByEmail(c, req.Email)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch user",
		})
		return
	}

	// Check Hashed Password
	if err := utils.VerifyPassword(user.Password, req.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// Generate JWT
	token, err := utils.CreateToken(int(user.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (h *Handler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	existingUser, err := h.Queries.GetUserByEmail(c, req.Email)
	if err == nil && existingUser.ID != 0 {
		// User already exists
		c.JSON(http.StatusConflict, gin.H{
			"error": "User with this email already exists",
		})
		return
	} else if err != nil && err != sql.ErrNoRows {
		// Some other error occurred
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to check if user exists",
		})
		return
	}

	hashedPass, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	usr, err := h.Queries.CreateUser(c, database.CreateUserParams{
		Name:       req.Name,
		Email:      req.Email,
		Password:   hashedPass,
		ProfileUrl: sql.NullString{},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": usr.ID,
	})
}

func (h *Handler) UpdateUser(c *gin.Context) {

}

func (h *Handler) LikeRecipe(c *gin.Context) {
	// Extract user_id from the context
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

	var req dto.LikeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Status == "like" {
		err := h.Queries.LikeRecipe(c, database.LikeRecipeParams{
			UserID:   sql.NullInt32{Int32: int32(uid)},
			RecipeID: sql.NullInt32{Int32: int32(req.RecipeID)},
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like recipe"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Recipe liked successfully"})

	} else if req.Status == "unlike" {
		err := h.Queries.UnlikeRecipe(c, database.UnlikeRecipeParams{
			UserID:   sql.NullInt32{Int32: int32(uid)},
			RecipeID: sql.NullInt32{Int32: int32(req.RecipeID)},
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unlike recipe"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Recipe unliked successfully"})

	} else {
		// Invalid status
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status. Must be 'like' or 'unlike'"})
	}
}

func (h *Handler) FavouriteRecipe(c *gin.Context) {
	// Extract user_id from the context
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

	var req dto.FavouriteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Status == "favourite" {
		err := h.Queries.FavouriteRecipe(c, database.FavouriteRecipeParams{
			UserID:   sql.NullInt32{Int32: int32(uid)},
			RecipeID: sql.NullInt32{Int32: int32(req.RecipeID)},
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to favourite recipe"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Recipe favourite successfully"})

	} else if req.Status == "unfavourite" {
		err := h.Queries.UnfavouriteRecipe(c, database.UnfavouriteRecipeParams{
			UserID:   sql.NullInt32{Int32: int32(uid)},
			RecipeID: sql.NullInt32{Int32: int32(req.RecipeID)},
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unfavourite recipe"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Recipe unfavourite successfully"})

	} else {
		// Invalid status
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status. Must be 'favourite' or 'unfavourite'"})
	}
}

func (h *Handler) NewRecipe(c *gin.Context) {
	// Extract user_id from the context
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

	// Bind the request body to the DTO
	var req dto.NewRecipeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set the user_id in the request
	req.UserID = uid

	recipeParams := database.CreateRecipeParams{
		UserID:       sql.NullInt32{Int32: int32(req.UserID), Valid: true},
		Title:        req.Title,
		Description:  sql.NullString{String: req.Description, Valid: req.Description != ""},
		Tags:         req.Tags,
		Ingredients:  req.Ingredients,
		Instructions: req.Instructions,
		TotalTime:    sql.NullInt32{Int32: int32(req.TotalTime), Valid: req.TotalTime > 0},
		Difficulty:   sql.NullString{String: req.Difficulty, Valid: req.Difficulty != ""},
		Servings:     sql.NullInt32{Int32: int32(req.Servings), Valid: req.Servings > 0},
		PhotoUrl:     sql.NullString{String: req.PhotoURL, Valid: req.PhotoURL != ""},
	}

	// Call the database query to create the recipe
	recipeID, err := h.Queries.CreateRecipe(c, recipeParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create recipe"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"recipe_id": recipeID})
}
