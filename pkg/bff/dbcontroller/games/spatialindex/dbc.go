package spatialindex

import (
	"database/sql"
	"fmt"

	"github.com/njayp/parthenon/pkg/bff/dbcli"
	"github.com/njayp/parthenon/pkg/bff/dbcontroller"
	"github.com/njayp/parthenon/pkg/bff/dbcontroller/games"
)

const (
	// ST_Distance_Sphere does not use index
	RADIUS_QUERY = `SELECT userid,
	ST_Distance_Sphere(location, %s) AS distance_m
	FROM %s
	HAVING distance_m <= 25000;`
	// this uses index
	RADIUS_QUERY2 = `SELECT userid, ST_Distance_Sphere(location, %s)
	FROM %s
	WHERE ST_Contains(ST_Buffer(%s, 25000), location);`
	// again, no index
	RADIUS_QUERY3 = `SELECT userid
	FROM %s
	WHERE ST_Distance_Sphere(location, %s) <= 25000;`
	TABLE_PROPS = `pk INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
	location POINT NOT NULL SRID 4326,
	userid BINARY(16),
	SPATIAL INDEX(location)`
	INSERT_QUERY = `INSERT INTO %s(location, userid)
	VALUES (%s,'%s');`
	SET_USER_LOCATION_QUERY = `SET %s = ST_GeomFromText('POINT(%s %s)', 4326);`
)

type spatialIndexDBC struct {
	dbcontroller.BaseDBController
}

func NewSpatialIndexDBC(db dbcli.DBCli) (*spatialIndexDBC, error) {
	client, err := db.EnsureDBandCli("games")
	if err != nil {
		return nil, err
	}

	dbc := &spatialIndexDBC{}
	dbc.SetClient(client)
	return dbc, nil
}

func (c *spatialIndexDBC) EnsureGame(tableName string) error {
	return c.BaseEnsureTable(
		tableName,
		TABLE_PROPS,
	)
}

// TODO make thread-safe
func (c *spatialIndexDBC) SetUserLocation(userid, latitude, longitude string) error {
	query := fmt.Sprintf(SET_USER_LOCATION_QUERY,
		games.UserLocationVarName(userid),
		latitude, longitude,
	)
	_, err := c.GetClient().Exec(query)
	return err
}

func (c *spatialIndexDBC) AddUser(tableName, userid string) error {
	query := fmt.Sprintf(
		INSERT_QUERY,
		tableName,
		games.UserLocationVarName(userid),
		userid,
	)
	_, err := c.GetClient().Exec(query)
	return err
}

func (c *spatialIndexDBC) SearchRadius(tableName, userid string) (*sql.Rows, error) {
	query := fmt.Sprintf(RADIUS_QUERY2,
		games.UserLocationVarName(userid),
		tableName,
		games.UserLocationVarName(userid),
	)
	return c.GetClient().Query(query)
}

type SearchResult struct {
	userid   string
	distance float64
}

func (s SearchResult) Distance() float64 {
	return s.distance
}

func (s SearchResult) UserID() string {
	return s.userid
}

func (c *spatialIndexDBC) ProcessSearchRadius(tableName, userid string) ([]SearchResult, error) {
	rows, err := c.SearchRadius(tableName, userid)
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
			userid:   *userid,
			distance: *distance,
		})
	}
	return results, nil
}
