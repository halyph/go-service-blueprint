//go:build integration

package repository

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserRepository_GetByID(t *testing.T) {
	db := SetupTestDB(t)
	if db == nil {
		return // Test was skipped
	}

	repo := NewUserRepository(db)
	ctx := context.Background()

	t.Run("returns user when exists", func(t *testing.T) {
		// Seed data migration already inserted users with IDs 1, 2, 3
		user, err := repo.GetByID(ctx, 1)

		require.NoError(t, err)
		assert.Equal(t, int64(1), user.ID)
		assert.Equal(t, "johndoe", user.Username)
		assert.Equal(t, "john@example.com", user.Email)
		assert.Equal(t, "John", user.FirstName)
		assert.Equal(t, "Doe", user.LastName)
		assert.True(t, user.IsActive)
	})

	t.Run("returns error when user not found", func(t *testing.T) {
		user, err := repo.GetByID(ctx, 999)

		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Contains(t, err.Error(), "user not found")
	})
}

func TestUserRepository_GetByUsername(t *testing.T) {
	db := SetupTestDB(t)
	if db == nil {
		return
	}

	repo := NewUserRepository(db)
	ctx := context.Background()

	t.Run("returns user when exists", func(t *testing.T) {
		user, err := repo.GetByUsername(ctx, "janedoe")

		require.NoError(t, err)
		assert.Equal(t, "janedoe", user.Username)
		assert.Equal(t, "jane@example.com", user.Email)
		assert.Equal(t, "Jane", user.FirstName)
		assert.Equal(t, "Doe", user.LastName)
	})

	t.Run("returns error when user not found", func(t *testing.T) {
		user, err := repo.GetByUsername(ctx, "nonexistent")

		assert.Error(t, err)
		assert.Nil(t, user)
	})
}

func TestUserRepository_ListActive(t *testing.T) {
	db := SetupTestDB(t)
	if db == nil {
		return
	}

	repo := NewUserRepository(db)
	ctx := context.Background()

	users, err := repo.ListActive(ctx)

	require.NoError(t, err)
	assert.Len(t, users, 2) // johndoe and janedoe are active, bobsmith is not

	// Should be ordered by username
	assert.Equal(t, "janedoe", users[0].Username)
	assert.Equal(t, "johndoe", users[1].Username)

	// All returned users should be active
	for _, user := range users {
		assert.True(t, user.IsActive)
	}
}
