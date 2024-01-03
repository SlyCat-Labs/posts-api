package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jdashel/posts-api/internal/domain/interfaces"
	"github.com/jdashel/posts-api/internal/domain/models"
)

// PostsHandler handles requests related to posts
type PostsHandlers struct {
	useCases interfaces.PostsUseCase
}

// NewHandler creates a new PostsHandlers instance
func NewPostsHandlers(useCases interfaces.PostsUseCase) PostsHandlers {
	return PostsHandlers{useCases: useCases}
}

// CreatePost creates a new post
func (h *PostsHandlers) CreatePost(c *gin.Context) {
	// Extract token from request
	token := c.Request.Header.Get("Authorization") // Assuming token is in the Authorization header
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Bind post data from request body
	var post models.Post // Assuming a models package with a Post struct
	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create the post using the use case
	createdPost, err := h.useCases.CreatePost(c.Request.Context(), token, &post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdPost)
}

// GetAllPosts retrieves posts with pagination
func (h *PostsHandlers) GetAllPosts(c *gin.Context) {
	// Extract token from request
	token := c.Request.Header.Get("Authorization") // Assuming token is in the Authorization header
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Get pagination parameters from query string
	pageNumber, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		// Handle invalid pageNumber
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}
	pageSize, err := strconv.Atoi(c.DefaultQuery("size", "10"))
	if err != nil {
		// Handle invalid pageSize
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page size"})
		return
	}

	// Fetch posts with pagination from the use case
	posts, err := h.useCases.GetAllPosts(c.Request.Context(), token, pageNumber, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}

// GetPostById retrieves a post by its ID
func (h *PostsHandlers) GetPostById(c *gin.Context) {
	// Extract token from request
	token := c.Request.Header.Get("Authorization") // Assuming token is in the Authorization header
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id := c.Param("id")
	post, err := h.useCases.GetPostById(c.Request.Context(), token, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if post == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// UpdatePost updates an existing post
func (h *PostsHandlers) UpdatePost(c *gin.Context) {
	// Extract token from request
	token := c.Request.Header.Get("Authorization") // Assuming token is in the Authorization header
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id := c.Param("id")
	var updatedPost models.Post
	if err := c.BindJSON(&updatedPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post, err := h.useCases.UpdatePost(c.Request.Context(), id, token, &updatedPost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
}

// DeletePost deletes a post by its ID
func (h *PostsHandlers) DeletePost(c *gin.Context) {
	// Extract token from request
	token := c.Request.Header.Get("Authorization") // Assuming token is in the Authorization header
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	id := c.Param("id")
	if err := h.useCases.DeletePost(c.Request.Context(), token, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
