package spatialindex

import (
	"database/sql"
	"fmt"

	"github.com/njayp/parthenon/pkg/bff/dbcli"
	"github.com/njayp/parthenon/pkg/bff/dbcontroller"
	"github.com/njayp/parthenon/pkg/bff/dbcontroller/games"
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
		`pk INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
		location POINT NOT NULL SRID 4326,
		userid BINARY(16),
	  	SPATIAL INDEX(location)`,
	)
}

// TODO make thread-safe
func (c *spatialIndexDBC) SetUserLocation(userid, latitude, longitude string) error {
	query := fmt.Sprintf(`SET %s = ST_GeomFromText('POINT(%s %s)', 4326);`,
		games.UserLocationVarName(userid),
		latitude, longitude,
	)
	_, err := c.GetClient().Exec(query)
	return err
}

func (c *spatialIndexDBC) AddUser(tableName, userid string) error {
	query := fmt.Sprintf(
		`INSERT INTO %s(location, userid)
		VALUES (%s,'%s');`,
		tableName,
		games.UserLocationVarName(userid),
		userid,
	)
	_, err := c.GetClient().Exec(query)
	return err
}

func (c *spatialIndexDBC) SearchRadius(tableName, userid string) (*sql.Rows, error) {
	query := fmt.Sprintf(`SELECT userid,
		ST_Distance_Sphere(location, %s) AS distance_m
		FROM %s
		HAVING distance_m <= 25000;`,
		games.UserLocationVarName(userid),
		tableName,
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
