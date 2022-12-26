package cli

import (
	"context"
	"database/sql"
)

// DB Client Factory
//
//go:generate mockgen -destination=mocks/db_mock_test.go . DBCli
type DBCli interface {
	EnsureDBandCli(ctx context.Context, dbName string) (*sql.DB, error)
	Ping(ctx context.Context) error
	PingUntilConnect(ctx context.Context) error
}
