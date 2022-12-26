package test

import (
	"context"

	"fmt"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/njayp/parthenon/pkg/db/controller/games/spatialindex"
)

func TestGamesFunctionality(t *testing.T) {
	ctx := context.Background()
	dbc, _ := setupDBnClients(ctx, t)
	gamesFunctionality(ctx, t, dbc)
}

func gamesFunctionality(ctx context.Context, t *testing.T, dbc *spatialindex.SpatialIndexDBC) {
	t.Run("games functionality", func(t *testing.T) {
		t.Parallel()

		queryContains := func(query, expected string) error {
			rows, err := dbc.Client.QueryContext(ctx, query)
			if err != nil {
				return fmt.Errorf("Query Error: %s", err.Error())
			}

			return rowsContains(rows, expected)
		}

		gameName := "chess"
		t.Run("add game", func(t *testing.T) {
			err := dbc.EnsureGameTable(ctx, gameName)
			if err != nil {
				t.Error(err)
			}
			err = queryContains(fmt.Sprintf("SHOW TABLES LIKE '%s';", gameName), gameName)
			if err != nil {
				t.Error(err)
			}
		})

		userid1 := uuid.New().String()
		userid2 := uuid.New().String()

		t.Run("add user to game", func(t *testing.T) {
			err := dbc.SetUserLocation(ctx, userid1, "33.4581414", "-111.9071715")
			if err != nil {
				t.Error(err)
			}
			err = dbc.AddUser(ctx, gameName, userid1)
			if err != nil {
				t.Error(err)
			}
		})

		t.Run("add second user to game. Search and find user.", func(t *testing.T) {
			err := dbc.SetUserLocation(ctx, userid2, "33.46", "-111.9")
			if err != nil {
				t.Error(err)
			}

			// search with user2, find user1
			results, err := dbc.ProcessSearchRadius(ctx, gameName, userid2)
			if err != nil {
				t.Error(err)
			}
			if len(results) < 1 || !strings.Contains(results[0].UserID, userid1) {
				t.Errorf("could not find user1: %+v", results)
			}

			// add user2
			err = dbc.AddUser(ctx, gameName, userid2)
			if err != nil {
				t.Error(err)
			}
		})

		t.Run("add third user to game. Search and not find user.", func(t *testing.T) {
			userid3 := uuid.New().String()
			err := dbc.SetUserLocation(ctx, userid3, "37", "-122")
			if err != nil {
				t.Error(err)
			}

			// add user3
			err = dbc.AddUser(ctx, gameName, userid3)
			if err != nil {
				t.Error(err)
			}

			// search with user1. user2 should be in results. user3 should not
			results, err := dbc.ProcessSearchRadius(ctx, gameName, userid1)
			if err != nil {
				t.Error(err)
			}

			// throw err if user3 in results
			for _, result := range results {
				if strings.Contains(result.UserID, userid3) {
					t.Error("user3 was found but should not have been")
				}
			}

			// throw err if user2 not in results
			for _, result := range results {
				if strings.Contains(result.UserID, userid2) {
					return
				}
			}
			t.Error("user2 was not found but should have been")
		})
	})
}
