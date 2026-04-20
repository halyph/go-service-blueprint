package factorial

import (
	"errors"
	"fmt"
)

// ErrNegativeInput is returned when the input is negative
var ErrNegativeInput = errors.New("factorial is not defined for negative numbers")

// Service provides factorial calculation functionality
type Service struct {
	storage Storage
}

// New creates a new factorial Service
func New(s Storage) *Service {
	return &Service{
		storage: s,
	}
}

// Calculate returns the factorial of n.
// Returns an error if n is negative.
func (s *Service) Calculate(n int64) (int64, error) {
	if n < 0 {
		return 0, ErrNegativeInput
	}

	if n == 0 || n == 1 {
		return 1, nil
	}

	alreadyCalculated, err := s.storage.Factorial(n)
	if err != nil {
		return 0, fmt.Errorf("checking storage: %w", err)
	}

	if alreadyCalculated > 0 {
		return alreadyCalculated, nil
	}

	result := n
	for i := n - 1; i > 1; i-- {
		result *= i
	}
	return result, nil
}
