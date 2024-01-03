package services

import (
	"context"

	"github.com/google/uuid"
)

// UUIDManager generates UUIDs
type UUIDService struct{}

func NewUUIDService() *UUIDService {
	return &UUIDService{}
}

// GenerateID generates a new UUID
func (m *UUIDService) GenerateID(ctx context.Context) (string, error) {
	return uuid.New().String(), nil
}
