package dbcli

import (
	"context"
	"database/sql"
)

// DBCli and Client Factory
//
//go:generate mockgen -destination=mocks/db_mock.go . DBCli
type DBCli interface {
	EnsureDBandCli(dbName string) (*sql.DB, error)
	Ping(ctx context.Context) error
	PingUntilConnect(ctx context.Context) error
}
