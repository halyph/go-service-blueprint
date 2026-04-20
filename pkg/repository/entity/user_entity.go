// Package entity contains database entity types with ORM mappings.
package entity

import (
	"time"

	"github.com/uptrace/bun"
)

// UserEntity represents the database model for users
type UserEntity struct {
	bun.BaseModel `bun:"table:users"`

	ID        int64     `bun:"id,pk,autoincrement"`
	Username  string    `bun:"username,notnull,unique"`
	Email     string    `bun:"email,notnull,unique"`
	FirstName string    `bun:"first_name"`
	LastName  string    `bun:"last_name"`
	IsActive  bool      `bun:"is_active,notnull,default:true"`
	CreatedAt time.Time `bun:"created_at,notnull,default:current_timestamp"`
}
