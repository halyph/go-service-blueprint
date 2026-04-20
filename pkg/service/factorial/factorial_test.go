package factorial

import (
	"errors"
	"testing"

	"github.com/halyph/go-service-blueprint/pkg/service/factorial/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_Calculate(t *testing.T) {
	t.Run("returns error for negative input", func(t *testing.T) {
		mockStorage := new(mocks.MockStorage)
		svc := New(mockStorage)

		result, err := svc.Calculate(-1)

		assert.Equal(t, int64(0), result)
		assert.Equal(t, ErrNegativeInput, err)
	})

	t.Run("returns 1 for input 0", func(t *testing.T) {
		mockStorage := new(mocks.MockStorage)
		svc := New(mockStorage)

		result, err := svc.Calculate(0)

		assert.Equal(t, int64(1), result)
		assert.NoError(t, err)
	})

	t.Run("returns 1 for input 1", func(t *testing.T) {
		mockStorage := new(mocks.MockStorage)
		svc := New(mockStorage)

		result, err := svc.Calculate(1)

		assert.Equal(t, int64(1), result)
		assert.NoError(t, err)
	})

	t.Run("returns cached value from storage", func(t *testing.T) {
		mockStorage := new(mocks.MockStorage)
		mockStorage.On("Factorial", int64(5)).Return(int64(120), nil)
		svc := New(mockStorage)

		result, err := svc.Calculate(5)

		assert.Equal(t, int64(120), result)
		assert.NoError(t, err)
		mockStorage.AssertExpectations(t)
	})

	t.Run("calculates factorial when not in storage", func(t *testing.T) {
		mockStorage := new(mocks.MockStorage)
		mockStorage.On("Factorial", int64(5)).Return(int64(0), nil)
		svc := New(mockStorage)

		result, err := svc.Calculate(5)

		assert.Equal(t, int64(120), result)
		assert.NoError(t, err)
		mockStorage.AssertExpectations(t)
	})

	t.Run("returns error when storage fails", func(t *testing.T) {
		mockStorage := new(mocks.MockStorage)
		storageErr := errors.New("storage error")
		mockStorage.On("Factorial", mock.Anything).Return(int64(0), storageErr)
		svc := New(mockStorage)

		result, err := svc.Calculate(5)

		assert.Equal(t, int64(0), result)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "checking storage")
		mockStorage.AssertExpectations(t)
	})
}
