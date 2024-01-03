package http

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jdashel/posts-api/internal/domain/usecases"
	"github.com/jdashel/posts-api/internal/infra/config"
	"github.com/jdashel/posts-api/internal/infra/database"
	"github.com/jdashel/posts-api/internal/infra/handlers"
	"github.com/jdashel/posts-api/internal/infra/repositories"
	"github.com/jdashel/posts-api/internal/infra/services"
)

// StartServer starts the HTTP server
func StartServer(ctx context.Context) error {
	config, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}

	db, err := database.Connect(ctx, config.DatabaseURL)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// Repositories injection
	postsRepository := repositories.NewPostsRepository(db)
	usersRepository := repositories.NewUsersRepository(db)

	// Services injection
	hashService := services.NewHashService()
	tokenService := services.NewTokenService(config.SecretKey)
	idService := services.NewUUIDService()

	// Usecases injection
	postsUsecases := usecases.NewPostsUseCases(postsRepository, tokenService, idService)
	usersUsecases := usecases.NewUsersUseCase(usersRepository, hashService, tokenService, idService)

	// Handlers injection
	usersHandler := handlers.NewUsersHandler(usersUsecases)
	postsHandler := handlers.NewPostsHandlers(postsUsecases)

	router := gin.Default()

	// Posts routes
	router.POST("/posts", postsHandler.CreatePost)
	router.GET("/posts", postsHandler.GetAllPosts)
	router.GET("/posts/:id", postsHandler.GetPostById)
	router.PUT("/posts/:id", postsHandler.UpdatePost)
	router.DELETE("/posts/:id", postsHandler.DeletePost)

	// Users routes
	router.POST("/signup", usersHandler.SignupHandler())
	router.POST("/signin", usersHandler.SigninHandler())
	router.GET("/profile", usersHandler.ProfileHandler())

	return router.Run("localhost:" + config.ServerPort)
}
