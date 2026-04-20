package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/halyph/go-service-blueprint/pkg/service/factorial"
)

var (
	version = "dev"
	commit  = "unknown"
)

// noopStorage is a simple storage implementation that doesn't cache anything
type noopStorage struct{}

func (n *noopStorage) Factorial(int64) (int64, error) {
	return 0, nil // Return 0 to indicate not cached
}

func main() {
	versionFlag := flag.Bool("version", false, "Print version information")
	numberFlag := flag.Int64("n", 5, "Calculate factorial of n")
	flag.Parse()

	if *versionFlag {
		fmt.Printf("go-service-blueprint CLI\n")
		fmt.Printf("Version: %s\n", version)
		fmt.Printf("Commit: %s\n", commit)
		fmt.Printf("Go: %s\n", runtime.Version())
		fmt.Printf("OS/Arch: %s/%s\n", runtime.GOOS, runtime.GOARCH)
		os.Exit(0)
	}

	if *numberFlag < 0 {
		fmt.Fprintf(os.Stderr, "Error: n must be non-negative\n")
		os.Exit(1)
	}

	service := factorial.New(&noopStorage{})
	result, err := service.Calculate(*numberFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("factorial(%d) = %d\n", *numberFlag, result)
}
