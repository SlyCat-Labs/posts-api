package interfaces

import (
	"context"

	"github.com/jdashel/posts-api/internal/domain/models"
)

// PostRepository defines the interface for interacting with posts data
type PostsRepository interface {
	Create(ctx context.Context, post *models.Post) (*models.Post, error)
	Read(ctx context.Context, id string, authorId string) (*models.Post, error)
	Find(ctx context.Context, authorId string, pageNumber int, pageSize int) ([]*models.Post, error)
	Update(ctx context.Context, id string, authorId string, post *models.Post) (*models.Post, error)
	Delete(ctx context.Context, id string, authorId string) error
}

// PostsUseCase represents the use cases for posts
type PostsUseCase interface {
	CreatePost(ctx context.Context, token string, post *models.Post) (*models.Post, error)
	GetPostById(ctx context.Context, token string, id string) (*models.Post, error)
	GetAllPosts(ctx context.Context, token string, pageNumber int, pageSize int) ([]*models.Post, error)
	UpdatePost(ctx context.Context, token string, id string, post *models.Post) (*models.Post, error)
	DeletePost(ctx context.Context, token string, id string) error
}
