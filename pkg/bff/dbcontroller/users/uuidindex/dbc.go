package uuidindex

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/njayp/parthenon/pkg/bff/dbcli"
	"github.com/njayp/parthenon/pkg/bff/dbcontroller"
)

const (
	// TODO
	TABLE_PROPS  = `userid BINARY(16) NOT NULL PRIMARY KEY`
	INSERT_QUERY = `INSERT INTO %s(userid)
	VALUES (UUID_TO_BIN('%s'));`
	GET_QUERY = `SELECT BIN_TO_UUID(userid)
	FROM %s
	WHERE userid = UUID_TO_BIN('%s')`
)

var (
	TRIBES = []string{"men", "women"}
)

type UUIDIndexDBC struct {
	dbcontroller.BaseDBController
}

func NewUUIDIndexDBC(ctx context.Context, db dbcli.DBCli) (*UUIDIndexDBC, error) {
	dbc := &UUIDIndexDBC{}

	err := dbc.EnsureDBandCli(ctx, db, "users")
	if err != nil {
		return nil, err
	}

	for _, tribe := range TRIBES {
		err = dbc.ensureUserTable(ctx, tribe)
		if err != nil {
			return nil, err
		}
	}

	return dbc, nil
}

func (u *UUIDIndexDBC) ensureUserTable(ctx context.Context, tableName string) error {
	return u.BaseEnsureTable(
		ctx,
		tableName,
		TABLE_PROPS,
	)
}

func (u *UUIDIndexDBC) AddUser(ctx context.Context, tableName, userid string) error {
	query := fmt.Sprintf(
		INSERT_QUERY,
		tableName,
		userid,
	)
	_, err := u.Client.ExecContext(ctx, query)
	return err
}

func (u *UUIDIndexDBC) GetUser(ctx context.Context, tableName, userid string) *sql.Row {
	query := fmt.Sprintf(
		GET_QUERY,
		tableName,
		userid,
	)
	return u.Client.QueryRowContext(ctx, query)
}
