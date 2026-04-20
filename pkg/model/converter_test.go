package model

import (
	"encoding/json"
	"flag"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var update = flag.Bool("update", false, "update golden files")

func TestUserConverter_ConvertUser_Golden(t *testing.T) {
	converter := &UserConverterImpl{}

	tests := []struct {
		name string
		user User
	}{
		{
			name: "full_user",
			user: User{
				ID:        1,
				Username:  "johndoe",
				Email:     "john@example.com",
				FirstName: "John",
				LastName:  "Doe",
				IsActive:  true,
			},
		},
		{
			name: "user_without_names",
			user: User{
				ID:        2,
				Username:  "noname",
				Email:     "noname@example.com",
				FirstName: "",
				LastName:  "",
				IsActive:  false,
			},
		},
		{
			name: "user_with_first_name_only",
			user: User{
				ID:        3,
				Username:  "alice",
				Email:     "alice@example.com",
				FirstName: "Alice",
				LastName:  "",
				IsActive:  true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Convert user to DTO
			dto := converter.ConvertUser(tt.user)

			// Marshal to JSON for golden file comparison
			actual, err := json.MarshalIndent(dto, "", "  ")
			require.NoError(t, err)

			// Golden file path
			goldenFile := filepath.Join("testdata", tt.name+".golden.json")

			if *update {
				// Update golden file
				err := os.WriteFile(goldenFile, actual, 0644)
				require.NoError(t, err)
			}

			// Read expected output from golden file
			expected, err := os.ReadFile(goldenFile)
			require.NoError(t, err, "Golden file not found. Run with -update to create it.")

			// Compare
			assert.JSONEq(t, string(expected), string(actual))
		})
	}
}

func TestUserConverter_ConvertUserList_Golden(t *testing.T) {
	converter := &UserConverterImpl{}

	users := []User{
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

	dtos := converter.ConvertUserList(users)

	// Marshal to JSON
	actual, err := json.MarshalIndent(dtos, "", "  ")
	require.NoError(t, err)

	goldenFile := filepath.Join("testdata", "user_list.golden.json")

	if *update {
		err := os.WriteFile(goldenFile, actual, 0644)
		require.NoError(t, err)
	}

	expected, err := os.ReadFile(goldenFile)
	require.NoError(t, err, "Golden file not found. Run with -update to create it.")

	assert.JSONEq(t, string(expected), string(actual))
}
