package usecases

import (
	"context"
	"errors"

	"github.com/jdashel/posts-api/internal/domain/interfaces"
	"github.com/jdashel/posts-api/internal/domain/models"
)

type PostsUseCases struct {
	repository   interfaces.PostsRepository
	tokenService interfaces.TokenService
	uuidService  interfaces.UUIDService
}

// Posts usecases constructor
func NewPostsUseCases(repository interfaces.PostsRepository,
	tokenService interfaces.TokenService, uuidService interfaces.UUIDService) *PostsUseCases {
	return &PostsUseCases{repository, tokenService, uuidService}
}

// CreatePost creates a new post
func (uc *PostsUseCases) CreatePost(ctx context.Context, token string, post *models.Post) (*models.Post, error) {
	// Validate user token
	claims, err := uc.tokenService.ParseToken(token)
	if err != nil {
		return nil, errors.New("invalid token")
	}
	authorID := claims["user_id"].(string)
	if err != nil {
		return nil, errors.New("invalid user ID in token")
	}

	// Generate a UUID for the post
	postID, err := uc.uuidService.GenerateID(ctx)
	if err != nil {
		return nil, err
	}
	post.ID = postID

	// Set the author ID
	post.AuthorID = authorID

	// Create the post in the repository
	createdPost, err := uc.repository.Create(ctx, post)
	if err != nil {
		return nil, err
	}

	return createdPost, nil
}

// GetPostById retrieves a post by ID
func (uc *PostsUseCases) GetPostById(ctx context.Context, token string, id string) (*models.Post, error) {
	// Validate user token
	claims, err := uc.tokenService.ParseToken(token)
	if err != nil {
		return nil, errors.New("invalid token")
	}
	authorID := claims["user_id"].(string)
	if err != nil {
		return nil, errors.New("invalid user ID in token")
	}

	return uc.repository.Read(ctx, id, authorID)
}

// GetAllPosts retrieves all posts
func (uc *PostsUseCases) GetAllPosts(ctx context.Context, token string) ([]*models.Post, error) {
	// Validate user token
	claims, err := uc.tokenService.ParseToken(token)
	if err != nil {
		return nil, errors.New("invalid token")
	}
	authorID := claims["user_id"].(string)
	if err != nil {
		return nil, errors.New("invalid user ID in token")
	}

	return uc.repository.Find(ctx, authorID)
}

// UpdatePost updates an existing post
func (uc *PostsUseCases) UpdatePost(ctx context.Context, token string, id string, post *models.Post) (*models.Post, error) {
	// Validate user token
	claims, err := uc.tokenService.ParseToken(token)
	if err != nil {
		return nil, errors.New("invalid token")
	}
	authorID := claims["user_id"].(string)
	if err != nil {
		return nil, errors.New("invalid user ID in token")
	}

	return uc.repository.Update(ctx, id, authorID, post)
}

// DeletePost deletes a post
func (uc *PostsUseCases) DeletePost(ctx context.Context, token string, id string) error {
	// Validate user token
	claims, err := uc.tokenService.ParseToken(token)
	if err != nil {
		return errors.New("invalid token")
	}
	authorID := claims["user_id"].(string)
	if err != nil {
		return errors.New("invalid user ID in token")
	}
	return uc.repository.Delete(ctx, id, authorID)
}
