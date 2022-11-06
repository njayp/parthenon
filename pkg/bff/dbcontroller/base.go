package dbcontroller

import (
	"database/sql"
	"fmt"
)

type BaseDBController struct {
	client *sql.DB
}

func (b *BaseDBController) SetClient(client *sql.DB) {
	b.client = client
}

func (b *BaseDBController) BaseEnsureTable(tableName, props string) error {
	_, err := b.client.Exec(fmt.Sprintf("CREATE TABLE [IF NOT EXISTS] %s(%s);", tableName, props))
	return err
}

//UNIQUE(id)
