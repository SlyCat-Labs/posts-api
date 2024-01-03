package interfaces

type HashService interface {
	// HashPassword hashes a password
	HashPassword(password string) (string, error)

	// ComparePassword compares a plain password with a hashed password
	ComparePassword(password string, hashedPassword string) bool
}
