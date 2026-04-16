package postgres

import (
	"auth/internal/domain"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(ctx context.Context, user *domain.User) error {
	id := uuid.New().String()
	query := `INSERT INTO users (id, email, password, created_at, updated_at) 
	          VALUES ($1, $2, $3, $4, $5)`
	
	now := time.Now()
	_, err := r.db.Exec(ctx, query, id, user.Email, user.Password, now, now)
	if err != nil {
		return fmt.Errorf("error inserting user: %w", err)
	}

	user.ID = id
	user.CreatedAt = now
	user.UpdatedAt = now
	return nil
}

func (r *UserRepo) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `SELECT id, email, password, created_at, updated_at FROM users WHERE email = $1`
	
	var user domain.User
	err := r.db.QueryRow(ctx, query, email).Scan(
		&user.ID, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		return nil, domain.ErrUserNotFound
	}

	return &user, nil
}

func (r *UserRepo) GetByID(ctx context.Context, id string) (*domain.User, error) {
	query := `SELECT id, email, password, created_at, updated_at FROM users WHERE id = $1`
	
	var user domain.User
	err := r.db.QueryRow(ctx, query, id).Scan(
		&user.ID, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		return nil, domain.ErrUserNotFound
	}

	return &user, nil
}

func (r *UserRepo) UpdatePassword(ctx context.Context, id string, newPassword string) error {
	query := `UPDATE users SET password = $1, updated_at = $2 WHERE id = $3`
	_, err := r.db.Exec(ctx, query, newPassword, time.Now(), id)
	return err
}
