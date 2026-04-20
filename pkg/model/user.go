package model

// User represents a domain user entity
type User struct {
	ID        int64
	Username  string
	Email     string
	FirstName string
	LastName  string
	IsActive  bool
}
