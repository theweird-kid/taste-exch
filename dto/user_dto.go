package dto

import "time"

// RegisterRequest represents the payload for user registration
type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// RegisterResponse represents the response for user registration
type RegisterResponse struct {
	Message string `json:"message"`
}

// SignInRequest represents the payload for user sign-in
type SignInRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// SignInResponse represents the response for user sign-in
type SignInResponse struct {
	Token string `json:"token"`
}

type UserResponse struct {
	ID          int       `json:"id"`          // User ID
	Name        string    `json:"name"`        // User's name
	Email       string    `json:"email"`       // User's email
	ProfileURL  string    `json:"profile_url"` // URL of the user's profile picture
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description"`
}
