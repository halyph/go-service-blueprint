# go-service-blueprint

Template Go service demonstrating [go-service-common-make](https://github.com/halyph/go-service-common-make) integration.

## Quick Start

```bash
make install      # Install tools (golangci-lint, mockery, goverter)
make test         # Run linting and tests
make build.local  # Build for local OS
```

## What's Included

**Structure:**
- `cmd/server` - HTTP server application
- `cmd/cli` - CLI tool (demonstrates multi-command builds)
- `pkg/model` - User domain model with goverter converters
- `pkg/service/factorial` - Example service with mockery mocks
- `pkg/repository` - Database layer with bun ORM
- `res/migrations` - SQL migrations (testcontainers in tests)

**Tools (in `.tools/`):**
- golangci-lint - Linting
- mockery - Mock generation from interfaces
- goverter - Type-safe struct converters

**Features:**
- Multi-command builds (server + cli)
- Integration tests with testcontainers (PostgreSQL)
- Golden file tests for JSON serialization
- Multi-arch builds (linux/darwin, amd64/arm64)
- Multi-Dockerfile support

## Common Commands

```bash
make help          # Show all available targets
make generate      # Generate mocks and converters
make goldenfiles   # Update golden test files
make build.linux   # Build for linux (amd64 + arm64)
make docker        # Build Docker images
make clean         # Clean artifacts
```

## Testing

```bash
make test          # Runs linting + unit tests + integration tests
```

Integration tests automatically start PostgreSQL using testcontainers. No manual Docker setup needed.

## Configuration

Override defaults in your Makefile after the `include` statement:

```makefile
TEAM := myteam
DOCKER_REGISTRY := registry.example.com
TEST_FLAGS += -tags=integration
```

## Files

- `.golangci.yml` - Linter configuration
- `.mockery.yaml` - Mock generation config
- `Dockerfile` - Base image
- `Dockerfile.extra` - Extended image (demonstrates multi-Dockerfile builds)
