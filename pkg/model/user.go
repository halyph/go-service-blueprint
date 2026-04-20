package model

// User represents a domain user entity
type User struct {
	ID        int64  `bun:"id,pk,autoincrement"`
	Username  string `bun:"username,notnull,unique"`
	Email     string `bun:"email,notnull,unique"`
	FirstName string `bun:"first_name"`
	LastName  string `bun:"last_name"`
	IsActive  bool   `bun:"is_active,notnull,default:true"`
}

// UserDTO represents a user data transfer object for API responses
type UserDTO struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Active   bool   `json:"active"`
}
