package factorial

// Storage defines the interface for factorial storage operations
// mockery will generate mocks for this interface
type Storage interface {
	Factorial(n int64) (int64, error)
}
