package application

import (
	"auth/internal/domain"
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo          domain.UserRepository
	tokenProvider domain.TokenProvider
}

func NewAuthService(repo domain.UserRepository, tp domain.TokenProvider) *AuthService {
	return &AuthService{
		repo:          repo,
		tokenProvider: tp,
	}
}

// Register se encarga de dar de alta a un nuevo usuario.
func (s *AuthService) Register(ctx context.Context, email, password string) error {
	// 1. Verificar si ya existe (Regla de negocio)
	existing, _ := s.repo.GetByEmail(ctx, email)
	if existing != nil {
		return domain.ErrUserAlreadyExists
	}

	// 2. Hashear la password (Seguridad ante todo, loco)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error hashing password: %w", err)
	}

	// 3. Crear la entidad
	user := &domain.User{
		Email:    email,
		Password: string(hashedPassword),
	}

	// 4. Persistir
	return s.repo.Create(ctx, user)
}

// Login verifica las credenciales y devuelve un token JWT si todo está OK.
func (s *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return "", domain.ErrInvalidCredentials
	}

	// Comparar la password con el hash
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", domain.ErrInvalidCredentials
	}

	// Generar el token (Arquitectura Hexagonal: usamos el puerto!)
	token, err := s.tokenProvider.Generate(user)
	if err != nil {
		return "", fmt.Errorf("error generando token: %w", err)
	}

	return token, nil
}

func (s *AuthService) ChangePassword(ctx context.Context, userID, oldPassword, newPassword string) error {
	user, err := s.repo.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	// 1. Verificar clave vieja
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword))
	if err != nil {
		return domain.ErrInvalidCredentials
	}

	// 2. Hashear la nueva clave
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// 3. Persistir
	return s.repo.UpdatePassword(ctx, userID, string(hashedPassword))
}
