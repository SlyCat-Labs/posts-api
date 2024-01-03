package usecases

import (
	"context"
	"errors"
	"fmt"

	"github.com/jdashel/posts-api/internal/domain/interfaces"
	"github.com/jdashel/posts-api/internal/domain/models"
)

type UsersUseCase struct {
	repository   interfaces.UsersRepository
	hashService  interfaces.HashService
	tokenService interfaces.TokenService
	uuidService  interfaces.UUIDService
}

// Users usecases constructor
func NewUsersUseCase(
	repository interfaces.UsersRepository,
	hashService interfaces.HashService,
	tokenService interfaces.TokenService,
	uuidService interfaces.UUIDService,
) *UsersUseCase {
	return &UsersUseCase{repository, hashService, tokenService, uuidService}
}

// Signup creates a new user account
func (uc *UsersUseCase) Signup(ctx context.Context, email, password string) (string, error) {
	// Hash password using hashService
	hashedPassword, err := uc.hashService.HashPassword(password)
	if err != nil {
		return "", err
	}
	password = hashedPassword

	// Generate a UUID for the post
	id, err := uc.uuidService.GenerateID(ctx)
	if err != nil {
		return "", err
	}

	user := &models.User{
		ID:       id,
		Email:    email,
		Password: password,
	}

	// Create user account in repository
	_, err = uc.repository.Create(ctx, user)
	if err != nil {
		return "", err
	}

	// Generate authentication token using tokenService
	token, err := uc.tokenService.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

// Signin authenticates a user
func (uc *UsersUseCase) Signin(ctx context.Context, email, password string) (string, error) {
	// Read user account in repository
	user, err := uc.repository.Find(ctx, email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	fmt.Println(user)

	// Validate password with hashManager
	match := uc.hashService.ComparePassword(password, user.Password)
	if !match {
		return "", errors.New("invalid email or password") // Avoid disclosing password error details
	}

	// Generate authentication token using tokenService
	token, err := uc.tokenService.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

// GetProfile retrieves a user's profile
func (uc *UsersUseCase) GetProfile(ctx context.Context, token string) (*models.User, error) {
	// Validate token using token manager
	claims, err := uc.tokenService.ParseToken(token)
	if err != nil {
		return nil, err
	}

	// Extract user ID from claims
	userID := claims["user_id"].(string)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID in token: %w", err)
	}

	// Retrieve user from repository
	user, err := uc.repository.Read(ctx, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// UpdateProfile updates a user's profile information
func (uc *UsersUseCase) UpdateProfile(ctx context.Context, token string, user *models.User) error {
	// Validate token and extract user ID
	claims, err := uc.tokenService.ParseToken(token)
	if err != nil {
		return err
	}
	userID := claims["user_id"].(string)
	if err != nil {
		return fmt.Errorf("invalid user ID in token: %w", err)
	}

	// Ensure user is updating their own profile
	if userID != user.ID {
		return errors.New("unauthorized to update another user's profile")
	}

	// Update user information in repository
	return uc.repository.Update(ctx, userID, user)
}

// DeleteProfile deletes a user's account
func (uc *UsersUseCase) DeleteProfile(ctx context.Context, token string, user *models.User) error {
	// Validate token and extract user ID
	claims, err := uc.tokenService.ParseToken(token)
	if err != nil {
		return err
	}
	userID := claims["user_id"].(string)
	if err != nil {
		return fmt.Errorf("invalid user ID in token: %w", err)
	}

	// Delete user from repository
	return uc.repository.Delete(ctx, userID)
}
