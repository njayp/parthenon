package mysqlcli

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/njayp/parthenon/pkg/bff/dbcli"
)

// NewMYSQLDBCli returns mysql client factory
func NewMYSQLDBCli(hostName string) (dbcli.DBCli, error) {
	return dbcli.NewBaseDBCli("mysql", hostName)
}
