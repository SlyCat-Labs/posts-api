package interfaces

import (
	"context"
)

// IDManager generates unique IDs
type UUIDService interface {
	GenerateID(ctx context.Context) (string, error)
}
