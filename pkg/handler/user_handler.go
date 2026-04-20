// Package handler provides HTTP request handlers for API endpoints.
package handler

import (
	"encoding/json"
	"net/http"

	"github.com/halyph/go-service-blueprint/pkg/model"
	"github.com/halyph/go-service-blueprint/pkg/model/converter/generated"
)

// UserHandler handles HTTP requests for users
type UserHandler struct {
	converter *generated.UserConverterImpl
}

// NewUserHandler creates a new user handler
func NewUserHandler() *UserHandler {
	return &UserHandler{
		converter: &generated.UserConverterImpl{},
	}
}

// GetUser handles GET /users/:id
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	// Example user (in real app, get from repository via service)
	user := model.User{
		ID:        1,
		Username:  "johndoe",
		Email:     "john@example.com",
		FirstName: "John",
		LastName:  "Doe",
		IsActive:  true,
	}

	userDTO := h.converter.ConvertUser(user)

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(userDTO)
}

// ListUsers handles GET /users
func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	// Example users (in real app, get from repository via service)
	users := []model.User{
		{
			ID:        1,
			Username:  "johndoe",
			Email:     "john@example.com",
			FirstName: "John",
			LastName:  "Doe",
			IsActive:  true,
		},
		{
			ID:        2,
			Username:  "janedoe",
			Email:     "jane@example.com",
			FirstName: "Jane",
			LastName:  "Doe",
			IsActive:  true,
		},
	}

	userDTOs := h.converter.ConvertUserList(users)

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"users": userDTOs,
		"total": len(userDTOs),
	})
}
