//go:build integration

package repository

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	migrate "github.com/rubenv/sql-migrate"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

// SetupTestDB creates a test database using testcontainers and runs migrations
func SetupTestDB(t *testing.T) *bun.DB {
	t.Helper()

	ctx := context.Background()

	// Start PostgreSQL container
	pgContainer, err := postgres.Run(ctx,
		"postgres:15-alpine",
		postgres.WithDatabase("blueprint_test"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(60*time.Second)),
	)
	if err != nil {
		t.Skipf("Skipping integration test: cannot start PostgreSQL container: %v", err)
		return nil
	}

	// Get connection string
	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		t.Fatalf("Failed to get connection string: %v", err)
	}

	// Connect to database
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(connStr)))

	if err := sqldb.Ping(); err != nil {
		t.Fatalf("Failed to ping database: %v", err)
	}

	// Create bun.DB
	db := bun.NewDB(sqldb, pgdialect.New())

	// Add query hook for debugging (optional)
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	// Run migrations using sql-migrate
	migrations := &migrate.FileMigrationSource{
		Dir: "../../res/migrations",
	}

	migrate.SetTable("schema_migrations")
	n, err := migrate.Exec(sqldb, "postgres", migrations, migrate.Up)
	if err != nil {
		t.Fatalf("Failed to run migrations: %v", err)
	}
	t.Logf("Applied %d migrations", n)

	// Clean up after test
	t.Cleanup(func() {
		_ = db.Close()
		if err := pgContainer.Terminate(ctx); err != nil {
			t.Logf("Failed to terminate container: %v", err)
		}
	})

	return db
}

// TruncateTables removes all data from tables (useful for test isolation)
func TruncateTables(t *testing.T, db *bun.DB, tables ...string) {
	t.Helper()
	ctx := context.Background()

	for _, table := range tables {
		_, err := db.ExecContext(ctx, fmt.Sprintf("TRUNCATE TABLE %s CASCADE", table))
		if err != nil {
			t.Fatalf("Failed to truncate table %s: %v", table, err)
		}
	}
}
