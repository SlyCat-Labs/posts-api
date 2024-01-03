package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jdashel/posts-api/internal/domain/usecases"
)

// SignupRequest represents the data required for a user signup request
type SignRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type UsersHandler struct {
	usecases *usecases.UsersUseCase
}

func NewUsersHandler(usecases *usecases.UsersUseCase) *UsersHandler {
	return &UsersHandler{usecases}
}

// SignupHandler handles user signup requests
func (uh *UsersHandler) SignupHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get signup request data
		var signupData SignRequest
		if err := c.BindJSON(&signupData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		token, err := uh.usecases.Signup(c.Request.Context(), signupData.Email, signupData.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"token": token})
	}
}

// SigninHandler handles user sign-in requests
func (uh *UsersHandler) SigninHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get sign-in request data
		var signinData SignRequest
		if err := c.BindJSON(&signinData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		token, err := uh.usecases.Signin(c.Request.Context(), signinData.Email, signinData.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

// ProfileHandler handles user profile requests
func (uh *UsersHandler) ProfileHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve authentication token from request (assuming it's in a header)
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Call GetProfile usecase with token
		profile, err := uh.usecases.GetProfile(c.Request.Context(), token)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Respond with profile data
		c.JSON(http.StatusOK, profile)
	}
}
