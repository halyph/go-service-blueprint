package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/halyph/go-service-blueprint/pkg/model"
	"github.com/halyph/go-service-blueprint/pkg/repository/converter/generated"
	"github.com/halyph/go-service-blueprint/pkg/repository/entity"
	"github.com/uptrace/bun"
)

// UserRepository handles user data persistence using bun
type UserRepository struct {
	db *bun.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *bun.DB) *UserRepository {
	return &UserRepository{db: db}
}

// GetByID retrieves a user by ID
func (r *UserRepository) GetByID(ctx context.Context, id int64) (*model.User, error) {
	ent := new(entity.UserEntity)
	err := r.db.NewSelect().
		Model(ent).
		Where("id = ?", id).
		Scan(ctx)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found: %d", id)
	}
	if err != nil {
		return nil, fmt.Errorf("query user: %w", err)
	}

	conv := &generated.UserEntityConverterImpl{}
	user := conv.EntityToModel(*ent)
	return &user, nil
}

// GetByUsername retrieves a user by username
func (r *UserRepository) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	ent := new(entity.UserEntity)
	err := r.db.NewSelect().
		Model(ent).
		Where("username = ?", username).
		Scan(ctx)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found: %s", username)
	}
	if err != nil {
		return nil, fmt.Errorf("query user: %w", err)
	}

	conv := &generated.UserEntityConverterImpl{}
	user := conv.EntityToModel(*ent)
	return &user, nil
}

// ListActive returns all active users
func (r *UserRepository) ListActive(ctx context.Context) ([]model.User, error) {
	var entities []entity.UserEntity
	err := r.db.NewSelect().
		Model(&entities).
		Where("is_active = ?", true).
		Order("username").
		Scan(ctx)

	if err != nil {
		return nil, fmt.Errorf("query active users: %w", err)
	}

	conv := &generated.UserEntityConverterImpl{}
	users := conv.EntityListToModelList(entities)
	return users, nil
}

// Create inserts a new user
func (r *UserRepository) Create(ctx context.Context, user *model.User) error {
	conv := &generated.UserEntityConverterImpl{}
	ent := conv.ModelToEntity(*user)

	_, err := r.db.NewInsert().
		Model(&ent).
		Exec(ctx)

	if err != nil {
		return fmt.Errorf("create user: %w", err)
	}

	return nil
}
