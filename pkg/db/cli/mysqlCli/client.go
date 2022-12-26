package mysqlCli

import (
	_ "github.com/go-sql-driver/mysql"
	dbcli "github.com/njayp/parthenon/pkg/db/cli"
)

// NewMYSQLDBCli returns mysql client factory
func NewMYSQLDBCli(hostName string) (*dbcli.BaseDBCli, error) {
	return dbcli.NewBaseDBCli("mysql", hostName)
}
