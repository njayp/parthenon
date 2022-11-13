package uuidindex

import (
	"context"
	"fmt"

	"github.com/njayp/parthenon/pkg/bff/dbcli"
	"github.com/njayp/parthenon/pkg/bff/dbcontroller"
)

const (
	// TODO
	TABLE_PROPS = `pk INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
	location POINT NOT NULL SRID 4326,
	userid BINARY(16),
	SPATIAL INDEX(location)`
	INSERT_QUERY = `INSERT INTO %s(location, userid)
	VALUES (%s);`
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
		err = dbc.EnsureUserTable(ctx, tribe)
		if err != nil {
			return nil, err
		}
	}

	return dbc, err
}

func (u *UUIDIndexDBC) EnsureUserTable(ctx context.Context, tableName string) error {
	return u.BaseEnsureTable(
		ctx,
		tableName,
		TABLE_PROPS,
	)
}

func (u *UUIDIndexDBC) AddUser(ctx context.Context, tableName, userid string) error {
	// TODO
	query := fmt.Sprintf(
		INSERT_QUERY,
		tableName,
		userid,
	)
	_, err := u.Client.ExecContext(ctx, query)
	return err
}
