package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/halyph/go-service-blueprint/pkg/model"
	"github.com/halyph/go-service-blueprint/pkg/model/converter"
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
			got := converter.FormatName(tt.firstName, tt.lastName)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFormatFullName(t *testing.T) {
	user := model.User{
		FirstName: "John",
		LastName:  "Doe",
	}

	got := converter.FormatFullName(user)
	assert.Equal(t, "John Doe", got)
}
