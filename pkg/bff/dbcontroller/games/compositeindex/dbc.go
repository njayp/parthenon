package compositeindex

import (
	"github.com/njayp/parthenon/pkg/bff/dbcli"
	"github.com/njayp/parthenon/pkg/bff/dbcontroller"
)

type CompositeIndexDBC struct {
	dbcontroller.BaseDBController
}

func NewCompositeIndexDBC(db dbcli.DBCli) (*CompositeIndexDBC, error) {
	client, err := db.EnsureDBandCli("games")
	if err != nil {
		return nil, err
	}

	dbc := &CompositeIndexDBC{}
	dbc.SetClient(client)
	return dbc, nil
}

func (c *CompositeIndexDBC) EnsureTable(tableName string) {
	c.BaseEnsureTable(
		tableName,
		`pk INT NOT NULL AUTO_INCREMENT,
		latitude   INT NOT NULL,
		longitude  	INT NOT NULL,
		PRIMARY KEY(pk),
    	INDEX location (longitude,latitude)`,
	)
}
