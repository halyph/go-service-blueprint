//go:build tools

package tools

import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/jmattheis/goverter/cmd/goverter"
	_ "github.com/vektra/mockery/v2"
)
