package dbcontroller

import (
	"database/sql"
	"fmt"
)

const (
	ENSURE_TABLE = "CREATE TABLE IF NOT EXISTS %s(%s);"
)

type BaseDBController struct {
	client *sql.DB
}

// not thread-safe
func (b *BaseDBController) SetClient(client *sql.DB) {
	b.client = client
}

func (b *BaseDBController) GetClient() *sql.DB {
	return b.client
}

func (b *BaseDBController) BaseEnsureTable(tableName, props string) error {
	_, err := b.client.Exec(fmt.Sprintf(ENSURE_TABLE, tableName, props))
	return err
}
