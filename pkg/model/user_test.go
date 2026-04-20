package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatName(t *testing.T) {
	tests := []struct {
		name      string
		firstName string
		lastName  string
		want      string
	}{
		{
			name:      "both names present",
			firstName: "John",
			lastName:  "Doe",
			want:      "John Doe",
		},
		{
			name:      "only first name",
			firstName: "John",
			lastName:  "",
			want:      "John",
		},
		{
			name:      "only last name",
			firstName: "",
			lastName:  "Doe",
			want:      "Doe",
		},
		{
			name:      "both empty",
			firstName: "",
			lastName:  "",
			want:      "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatName(tt.firstName, tt.lastName)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFormatFullName(t *testing.T) {
	user := User{
		FirstName: "John",
		LastName:  "Doe",
	}

	got := FormatFullName(user)
	assert.Equal(t, "John Doe", got)
}
