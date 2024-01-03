package interfaces

import "github.com/golang-jwt/jwt"

type TokenService interface {
	// GenerateToken generates a new authentication token for a user ID
	GenerateToken(userID string) (string, error)

	// ParseToken parses a token and returns its claims
	ParseToken(tokenString string) (jwt.MapClaims, error)
}
