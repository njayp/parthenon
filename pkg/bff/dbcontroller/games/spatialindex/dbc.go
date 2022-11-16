package spatialindex

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/njayp/parthenon/pkg/bff/dbcli"
	"github.com/njayp/parthenon/pkg/bff/dbcontroller"
	"github.com/njayp/parthenon/pkg/bff/dbcontroller/games"
)

const (
	// this uses index
	RADIUS_QUERY = `SELECT BIN_TO_UUID(userid), ST_Distance_Sphere(location, %s)
	FROM %s
	WHERE ST_Contains(ST_Buffer(%s, 25000), location);`
	TABLE_PROPS = `pk INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
	location POINT NOT NULL SRID 4326,
	userid BINARY(16),
	SPATIAL INDEX(location)`
	INSERT_QUERY = `INSERT INTO %s(location, userid)
	VALUES (%s,UUID_TO_BIN('%s'));`
	SET_USER_LOCATION_QUERY = `SET %s = ST_GeomFromText('POINT(%s %s)', 4326);`
)

type SpatialIndexDBC struct {
	dbcontroller.BaseDBController
}

func NewSpatialIndexDBC(ctx context.Context, db dbcli.DBCli) (*SpatialIndexDBC, error) {
	dbc := &SpatialIndexDBC{}
	err := dbc.EnsureDBandCli(ctx, db, "games")
	return dbc, err
}

func (c *SpatialIndexDBC) EnsureGameTable(ctx context.Context, tableName string) error {
	return c.BaseEnsureTable(
		ctx,
		tableName,
		TABLE_PROPS,
	)
}

func (c *SpatialIndexDBC) SetUserLocation(ctx context.Context, userid, latitude, longitude string) error {
	uservar := games.UserLocationVarName(userid)
	query := fmt.Sprintf(SET_USER_LOCATION_QUERY,
		uservar,
		latitude, longitude,
	)
	_, err := c.Client.ExecContext(ctx, query)
	return err
}

func (c *SpatialIndexDBC) AddUser(ctx context.Context, tableName, userid string) error {
	uservar := games.UserLocationVarName(userid)
	query := fmt.Sprintf(
		INSERT_QUERY,
		tableName,
		uservar,
		userid,
	)
	_, err := c.Client.ExecContext(ctx, query)
	return err
}

func (c *SpatialIndexDBC) SearchRadius(ctx context.Context, tableName, userid string) (*sql.Rows, error) {
	uservar := games.UserLocationVarName(userid)
	query := fmt.Sprintf(RADIUS_QUERY,
		uservar,
		tableName,
		uservar,
	)
	return c.Client.QueryContext(ctx, query)
}

type SearchResult struct {
	UserID   string
	Distance float64
}

func (c *SpatialIndexDBC) ProcessSearchRadius(ctx context.Context, tableName, userid string) ([]SearchResult, error) {
	rows, err := c.SearchRadius(ctx, tableName, userid)
	if err != nil {
		return nil, err
	}
	results := make([]SearchResult, 0)
	for rows.Next() {
		distance := new(float64)
		userid := new(string)
		err = rows.Scan(userid, distance)
		if err != nil {
			return nil, err
		}
		results = append(results, SearchResult{
			UserID:   *userid,
			Distance: *distance,
		})
	}
	return results, nil
}
