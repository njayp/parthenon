package dbcli

import "database/sql"

// DBCli and Client Factory
//
//go:generate mockgen -destination=mocks/db_mock.go . DB
type DBCli interface {
	EnsureDBandCli(dbName string) (*sql.DB, error)
}
