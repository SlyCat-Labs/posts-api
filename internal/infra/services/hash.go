package services

import (
	"golang.org/x/crypto/bcrypt"
)

// HashManager handles password hashing
type HashService struct{}

func NewHashService() *HashService {
	return &HashService{}
}

// HashPassword hashes a password using bcrypt
func (hm *HashService) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// ComparePassword compares a plain password with a hashed password
func (hm *HashService) ComparePassword(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
