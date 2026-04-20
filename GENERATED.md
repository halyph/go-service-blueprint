# Generated Code

This project uses code generation tools. Generated code is **committed to the repository** but can be identified and regenerated using the patterns below.

## Finding All Generated Code

### Mockery (Test Mocks)
```bash
find . -path "*/mocks/*.go" -type f
```

### Goverter (Type Converters)
```bash
find . -path "*/generated/*.go" -type f
```

### All Generated Code
```bash
find . \( -path "*/mocks/*.go" -o -path "*/generated/*.go" \) -type f
```

## Regenerating All Code

```bash
make generate
```

This runs:
- `mockery` - Generates mocks from interfaces (configured in `.mockery.yaml`)
- `goverter` - Generates type-safe converters (configured via `//go:generate` directives)

## Generated Code Locations

**Current structure:**
```
pkg/
├── model/converter/generated/     # goverter: User → UserDTO converters
├── repository/converter/generated/ # goverter: UserEntity ↔ User converters
└── service/factorial/mocks/       # mockery: Storage interface mock
```

## Why Commit Generated Code?

1. **Faster builds** - No generation step needed for CI/builds
2. **Code review** - Changes in generated code are visible in PRs
3. **Reproducibility** - Anyone can build without installing generators
4. **Type safety** - IDE can validate generated code immediately

## Regeneration Triggers

Run `make generate` after:
- Modifying converter interfaces (goverter)
- Adding/changing interface methods (mockery)
- Updating `.mockery.yaml` configuration
