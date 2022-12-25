package dbcli

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

// shared logic for sql DBs
type BaseDBCli struct {
	driver, url string
	rootClient  *sql.DB
}

// DO NOT CALL DIRECTLY. NO DRIVERS LOADED
func NewBaseDBCli(driver, hostName string) (*BaseDBCli, error) {
	db := &BaseDBCli{
		driver: driver,
		// root (user)
		// password (password)
		// using this string will connect w/o specifying a db,
		// concat with a dbName to access that db
		url: fmt.Sprintf("root:password@tcp(%s:3306)/", hostName),
	}

	var err error
	db.rootClient, err = db.makeCli("")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (b *BaseDBCli) makeDB(ctx context.Context, dbName string) error {
	_, err := b.rootClient.ExecContext(ctx, fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s;", dbName))
	return err
}

func (b *BaseDBCli) makeCli(dbName string) (*sql.DB, error) {
	return sql.Open(b.driver, b.url+dbName)
}

// EnsureDBandClient ensures that the DB exists, then creates a client for it
func (b *BaseDBCli) EnsureDBandCli(ctx context.Context, dbName string) (*sql.DB, error) {
	err := b.makeDB(ctx, dbName)
	if err != nil {
		return nil, err
	}
	return b.makeCli(dbName)
}

func (b *BaseDBCli) Ping(ctx context.Context) error {
	return b.rootClient.PingContext(ctx)
}

func (b *BaseDBCli) PingUntilConnect(ctx context.Context) error {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("context cancelled")
		case <-ticker.C:
			if err := b.Ping(ctx); err == nil {
				return nil
			}
		}
	}
}
