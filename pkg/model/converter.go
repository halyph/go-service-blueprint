package model

//go:generate goverter gen .

// UserConverter converts between User and UserDTO
// goverter:converter
// goverter:output:file generated_converter.go
type UserConverter interface {
	// ConvertUser converts User to UserDTO
	// goverter:map . FullName | FormatFullName
	// goverter:map IsActive Active
	ConvertUser(source User) UserDTO

	// ConvertUserList converts slice of Users to DTOs
	ConvertUserList(source []User) []UserDTO
}

// FormatFullName formats full name for DTO from User
func FormatFullName(user User) string {
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
