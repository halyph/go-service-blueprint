package dto

// UserDTO represents a user data transfer object for API responses
type UserDTO struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Active   bool   `json:"active"`
}
