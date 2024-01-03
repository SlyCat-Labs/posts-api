package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jdashel/posts-api/internal/domain/models"
)

type UsersRepository struct {
	db *sql.DB
}

// UsersRepository constructor
func NewUsersRepository(db *sql.DB) *UsersRepository {
	return &UsersRepository{db: db}
}

// Create a new user
func (repo *UsersRepository) Create(ctx context.Context, user *models.User) (*models.User, error) {
	stmt := `INSERT INTO users (id, email, password) VALUES ($1, $2, $3) RETURNING *`
	row := repo.db.QueryRowContext(ctx, stmt, user.ID, user.Email, user.Password)

	var newUser models.User
	var deletedAt sql.NullTime
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &deletedAt)
	if err != nil {
		return nil, err
	}

	return &newUser, nil
}

// Read a user by id
func (repo *UsersRepository) Read(ctx context.Context, id string) (*models.User, error) {
	stmt := `SELECT * FROM users WHERE id = $1`
	row := repo.db.QueryRowContext(ctx, stmt, id)

	var user models.User
	var deletedAt sql.NullTime
	err := row.Scan(&user.ID, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt, &deletedAt)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	} else if err != nil {
		return nil, err
	}

	return &user, nil
}

// Finda user by email
func (repo *UsersRepository) Find(ctx context.Context, email string) (*models.User, error) {
	stmt := `SELECT * FROM users WHERE email = $1`
	row := repo.db.QueryRowContext(ctx, stmt, email)

	var user models.User
	var deletedAt sql.NullTime
	err := row.Scan(&user.ID, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt, &deletedAt)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	} else if err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateProfile updates a user's profile information
func (repo *UsersRepository) Update(ctx context.Context, id string, user *models.User) error {
	stmt := `UPDATE users SET email = $1 WHERE id = $2`
	_, err := repo.db.ExecContext(ctx, stmt, user.Email, id)
	return err
}

// DeleteProfile deletes a user's account
func (repo *UsersRepository) Delete(ctx context.Context, id string) error {
	stmt := `DELETE FROM users WHERE id = $1`
	_, err := repo.db.ExecContext(ctx, stmt, id)
	return err
}
