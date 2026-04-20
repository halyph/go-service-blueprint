// Package converter provides type conversions between domain models and DTOs.
package converter

import (
	"github.com/halyph/go-service-blueprint/pkg/handler/dto"
	"github.com/halyph/go-service-blueprint/pkg/model"
)

// UserConverter converts between User and UserDTO
// goverter:converter
type UserConverter interface {
	// ConvertUser converts User to UserDTO
	// goverter:map . FullName | FormatFullName
	// goverter:map IsActive Active
	ConvertUser(source model.User) dto.UserDTO

	// ConvertUserList converts slice of Users to DTOs
	ConvertUserList(source []model.User) []dto.UserDTO
}

// FormatFullName formats full name for DTO from User
func FormatFullName(user model.User) string {
	return FormatName(user.FirstName, user.LastName)
}

// FormatName combines first and last name
func FormatName(firstName, lastName string) string {
	if firstName == "" && lastName == "" {
		return ""
	}
	if firstName == "" {
		return lastName
	}
	if lastName == "" {
		return firstName
	}
	return firstName + " " + lastName
}
