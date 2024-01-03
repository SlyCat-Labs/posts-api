package services

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

// TokenService handles token generation and validation
type TokenService struct {
	secretKey string
}

// NewTokenService creates a new TokenService instance
func NewTokenService(secretKey string) *TokenService {
	return &TokenService{secretKey: secretKey}
}

// GenerateToken generates a new authentication token for a user ID
func (tm *TokenService) GenerateToken(userID string) (string, error) {
	// Create JWT claims with user ID and expiration time
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Expire in 24 hours
	}

	// Generate the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(tm.secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseToken parses a token and returns its claims
func (tm *TokenService) ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(strings.TrimPrefix(tokenString, "Bearer "), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(tm.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims type: %v", token.Claims)
	}
	fmt.Println(claims)

	return claims, nil
}
