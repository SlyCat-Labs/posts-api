package interfaces

import (
	"context"

	"github.com/jdashel/posts-api/internal/domain/models"
)

// UsersRepository defines the interface for interacting with user data
type UsersRepository interface {
	Create(ctx context.Context, user *models.User) (*models.User, error)
	Read(ctx context.Context, id string) (*models.User, error)
	Find(ctx context.Context, email string) (*models.User, error)
	Update(ctx context.Context, id string, user *models.User) error
	Delete(ctx context.Context, id string) error
}

// UsersUseCase represents the use cases for users
type UsersUseCase interface {
	Signup(ctx context.Context, user *models.User) (string, error)
	Signin(ctx context.Context, email string, password string) (string, error)
	GetProfile(ctx context.Context, token string) (*models.User, error)
	UpdateProfile(ctx context.Context, token string) error
	DeleteProfile(ctx context.Context, token string) error
}
