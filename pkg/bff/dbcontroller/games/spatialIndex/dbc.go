package compositeindex

import (
	"database/sql"
	"fmt"

	"github.com/njayp/parthenon/pkg/bff/dbcli"
	"github.com/njayp/parthenon/pkg/bff/dbcontroller"
	"github.com/njayp/parthenon/pkg/bff/dbcontroller/games"
)

type SpatialIndexDBC struct {
	dbcontroller.BaseDBController
}

func NewSpatialIndexDBC(db dbcli.DBCli) (*SpatialIndexDBC, error) {
	client, err := db.EnsureDBandCli("games")
	if err != nil {
		return nil, err
	}

	dbc := &SpatialIndexDBC{}
	dbc.SetClient(client)
	return dbc, nil
}

func (c *SpatialIndexDBC) EnsureGame(tableName string) error {
	return c.BaseEnsureTable(
		tableName,
		`'pk' INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
		'position' POINT NOT NULL SRID 4326,
		'userid' BINARY(16),
	  	SPATIAL INDEX('position')`,
	)
}

// TODO make thread-safe
func (c *SpatialIndexDBC) SetUserLocation(userid, latitude, longituede string) error {
	query := fmt.Sprintf(`SET %s = ST_GeomFromText('POINT(%s, %s)', 4326);`,
		games.UserLocationVarName(userid),
		latitude,
		longituede,
	)
	_, err := c.GetClient().Exec(query)
	return err
}

func (c *SpatialIndexDBC) AddUser(tableName, userid string) error {
	query := fmt.Sprintf(
		`INSERT INTO %s("position, userid")
		VALUES (%s,%s);`,
		tableName,
		games.UserLocationVarName(userid),
		userid,
	)
	_, err := c.GetClient().Exec(query)
	return err
}

func (c *SpatialIndexDBC) Search(tableName, userid string) (*sql.Rows, error) {
	query := fmt.Sprintf(`SELECT name,
		ST_Distance_Sphere('position', %s) AS 'distance_m'
		FROM %s
		HAVING distance_m <= 25000;`,
		games.UserLocationVarName(userid),
		tableName,
	)
	return c.GetClient().Query(query)
}
