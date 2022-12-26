package controller

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/njayp/parthenon/pkg/db/cli"
)

const (
	ENSURE_TABLE = "CREATE TABLE IF NOT EXISTS %s(%s);"
)

type BaseDBController struct {
	Client *sql.DB
}

func (b *BaseDBController) BaseEnsureTable(ctx context.Context, tableName, props string) error {
	_, err := b.Client.ExecContext(ctx, fmt.Sprintf(ENSURE_TABLE, tableName, props))
	return err
}

func (b *BaseDBController) EnsureDBandCli(ctx context.Context, db cli.DBCli, dbName string) error {
	var err error
	b.Client, err = db.EnsureDBandCli(ctx, dbName)
	return err
}
