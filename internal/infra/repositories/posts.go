package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jdashel/posts-api/internal/domain/models"
)

type PostsRepository struct {
	db *sql.DB
}

// PostsRepository constructor
func NewPostsRepository(db *sql.DB) *PostsRepository {
	return &PostsRepository{db: db}
}

// CreatePost creates a new post in the database
func (repo *PostsRepository) Create(ctx context.Context, post *models.Post) (*models.Post, error) {
	stmt := `INSERT INTO posts (id, title, content, author_id) VALUES ($1, $2, $3, $4) RETURNING *`
	// Use QueryRowContext to retrieve the generated ID
	row := repo.db.QueryRowContext(ctx, stmt, post.ID, post.Title, post.Content, post.AuthorID)

	var deletedAt sql.NullTime
	newPost := &models.Post{}
	err := row.Scan(&newPost.ID, &newPost.Title, &newPost.Content, &newPost.AuthorID, &newPost.CreatedAt, &newPost.UpdatedAt, &deletedAt)
	if err != nil {
		return nil, err
	}

	return post, nil
}

// GetPostById retrieves a post by ID
func (repo *PostsRepository) Read(ctx context.Context, id string, authorId string) (*models.Post, error) {
	stmt := `SELECT * FROM posts WHERE id = $1 AND author_id = $2`
	row := repo.db.QueryRowContext(ctx, stmt, id, authorId)

	var post models.Post
	var deletedAt sql.NullTime
	err := row.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.CreatedAt, &post.UpdatedAt, &deletedAt)
	if err == sql.ErrNoRows {
		return nil, errors.New("post not found")
	} else if err != nil {
		return nil, err
	}

	return &post, nil
}

// GetAllPosts retrieves posts with pagination
func (repo *PostsRepository) Find(ctx context.Context, authorId string, pageNumber int, pageSize int) ([]*models.Post, error) {
	offset := (pageNumber - 1) * pageSize

	stmt := `SELECT * FROM posts WHERE author_id = $1 ORDER BY created_at ASC OFFSET $2 LIMIT $3`
	rows, err := repo.db.QueryContext(ctx, stmt, authorId, offset, pageSize)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*models.Post
	var deletedAt sql.NullTime
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.CreatedAt, &post.UpdatedAt, &deletedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

// UpdatePost updates an existing post
func (repo *PostsRepository) Update(ctx context.Context, id string, authorId string, post *models.Post) (*models.Post, error) {
	stmt := `UPDATE posts SET title = $1, content = $2 WHERE id = $3 AND author_id = $4 RETURNING *`
	row := repo.db.QueryRowContext(ctx, stmt, post.Title, post.Content, id, authorId)

	// Scan updated post data into a new struct to avoid potential conflicts
	var updatedPost models.Post
	var deletedAt sql.NullTime
	if err := row.Scan(&updatedPost.ID, &updatedPost.Title, &updatedPost.Content, &updatedPost.AuthorID, &deletedAt); err != nil {
		return nil, err
	}

	return &updatedPost, nil
}

// DeletePost deletes a post
func (repo *PostsRepository) Delete(ctx context.Context, id string, authorId string) error {
	stmt := `DELETE FROM posts WHERE id = $1 AND author_id = $2`
	_, err := repo.db.ExecContext(ctx, stmt, id, authorId)
	return err
}
