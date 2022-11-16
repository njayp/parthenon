package test

import (
	"context"
	"database/sql"

	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/njayp/parthenon/pkg/bff/dbcontroller/games"
	"github.com/njayp/parthenon/pkg/bff/dbcontroller/games/spatialindex"
)

const (
	// ST_Distance_Sphere does not use index
	RQ = `SELECT userid,
		ST_Distance_Sphere(location, %s) AS distance_m
		FROM %s
		HAVING distance_m <= 25000;`

	// again, no index
	RQ2 = `SELECT userid
		FROM %s
		WHERE ST_Distance_Sphere(location, %s) <= 25000;`
)

func TestGamesOptimization(t *testing.T) {
	ctx := context.Background()
	dbc, _ := setupDBnClients(ctx, t)
	gamesOptimization(ctx, t, dbc)
}

func gamesOptimization(ctx context.Context, t *testing.T, dbc *spatialindex.SpatialIndexDBC) {
	t.Run("games optimization", func(t *testing.T) {
		t.Parallel()

		// add game to run queries
		gameName := "optimization"
		err := dbc.EnsureGameTable(ctx, gameName)
		if err != nil {
			t.Fatal(err)
		}

		// add user to prime queries
		userid := uuid.New().String()
		dbc.SetUserLocation(ctx, userid, "0", "0")
		dbc.AddUser(ctx, gameName, userid)

		t.Run("use EXPLAIN on search query. Ensure it uses the index", func(t *testing.T) {
			query := fmt.Sprintf("EXPLAIN "+spatialindex.RADIUS_QUERY, games.UserLocationVarName(userid), gameName, games.UserLocationVarName(userid))
			row := dbc.Client.QueryRow(query)

			key := new(sql.NullString)
			possibleKeys := new(sql.NullString)
			keyLen := new(sql.NullString)
			searchtype := new(sql.NullString)
			selectType := new(sql.NullString)
			err = row.Scan(new(sql.NullString), selectType, new(sql.NullString), new(sql.NullString), searchtype, possibleKeys, key, keyLen, new(sql.NullString), new(sql.NullString), new(sql.NullString), new(sql.NullString))
			if err != nil {
				t.Error(err)
			}
			if key.String != "location" || selectType.String != "SIMPLE" || searchtype.String != "range" {
				t.Error(fmt.Errorf("\nquery type: %s\nsearch type: %s\npossible: %s\nkey: %s\nkey len:%s", selectType.String, searchtype.String, possibleKeys.String, key.String, keyLen.String))
			}
		})
	})
}
