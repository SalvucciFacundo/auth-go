package domain

import (
	"context"
	"errors"
	"time"
)

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

// User representa la entidad principal de nuestro sistema de auth.
type User struct {
	ID        string
	Email     string
	Password  string // Siempre hasheada, nunca texto plano!
	CreatedAt time.Time
	UpdatedAt time.Time
}

// UserRepository es el "Puerto" (Interface) para persistir usuarios.
// Notá que recibe un context.Context, es la buena práctica de Go para timeouts.
type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetByID(ctx context.Context, id string) (*User, error)
	UpdatePassword(ctx context.Context, id string, newPassword string) error
}

// TokenProvider es el Puerto para manejar la seguridad de tokens.
type TokenProvider interface {
	Generate(user *User) (string, error)
	Validate(token string) (string, error) // Devuelve el userID si es válido
}
