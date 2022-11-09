package test

import (
	"context"
	"database/sql"

	"fmt"
	"strings"
	"testing"

	"github.com/njayp/parthenon/pkg/bff/dbcontroller/games/spatialindex"
	"github.com/njayp/parthenon/pkg/bff/dbcontroller/test"
)

func TestDBC(t *testing.T) {
	ctx := context.Background()
	cli, rm, err := test.SetupDB(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer rm(ctx)

	dbc, err := spatialindex.NewSpatialIndexDBC(cli)
	if err != nil {
		t.Fatal(err)
	}

	rowsContains := func(rows *sql.Rows, expected string) error {
		var text string
		line := new(string)
		for rows.Next() {
			err := rows.Scan(line)
			if err != nil {
				return fmt.Errorf("Scan Error: %s", err.Error())
			}
			text += *line
		}
		if !strings.Contains(text, expected) {
			return fmt.Errorf("%s does not contain %s", text, expected)
		}

		return nil
	}

	queryContains := func(query, expected string) error {
		rows, err := dbc.GetClient().Query(query)
		if err != nil {
			return fmt.Errorf("Query Error: %s", err.Error())
		}

		return rowsContains(rows, expected)
	}

	gameName := "chess"
	t.Run("add game", func(t *testing.T) {
		err := dbc.EnsureGame(gameName)
		if err != nil {
			t.Error(err)
		}
		err = queryContains(fmt.Sprintf("SHOW TABLES LIKE '%s';", gameName), gameName)
		if err != nil {
			t.Error(err)
		}
	})

	userid1 := "testuser1"
	t.Run("add user to game", func(t *testing.T) {

		err := dbc.SetUserLocation(userid1, "33.4581414", "-111.9071715")
		if err != nil {
			t.Error(err)
		}
		err = dbc.AddUser(gameName, userid1)
		if err != nil {
			t.Error(err)
		}

		// check that user1 is in table
		err = queryContains(fmt.Sprintf("SELECT userid FROM %s;", gameName), userid1)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("add second user to game. Search and find user.", func(t *testing.T) {
		userid2 := "testuser2"
		err := dbc.SetUserLocation(userid2, "33.46", "-111.9")
		if err != nil {
			t.Error(err)
		}

		results, err := dbc.ProcessSearchRadius(gameName, userid2)
		if err != nil {
			t.Error(err)
		}

		if len(results) < 1 || !strings.Contains(results[0].UserID(), userid1) {
			t.Errorf("could not find user1: %+v", results)
		}

		// add user2
		err = dbc.AddUser(gameName, userid2)
		if err != nil {
			t.Error(err)
		}

		// check that user2 is in table
		err = queryContains(fmt.Sprintf("SELECT userid FROM %s;", gameName), userid2)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("add third user to game. Search and not find user.", func(t *testing.T) {
		userid3 := "testuser3"
		err := dbc.SetUserLocation(userid3, "37", "-122")
		if err != nil {
			t.Error(err)
		}

		// add user3
		err = dbc.AddUser(gameName, userid3)
		if err != nil {
			t.Error(err)
		}

		// check that user3 is in table
		err = queryContains(fmt.Sprintf("SELECT userid FROM %s;", gameName), userid3)
		if err != nil {
			t.Error(err)
		}

		// search with user1. user2 should be in results. user3 should not
		results, err := dbc.ProcessSearchRadius(gameName, userid1)
		if err != nil {
			t.Error(err)
		}

		for _, result := range results {
			if strings.Contains(result.UserID(), userid3) {
				t.Error("user3 was found but should not have been")
			}
		}
	})
}
