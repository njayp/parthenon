package test

import (
	"context"
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

	queryContains := func(query, expected string) error {
		rows, err := dbc.GetClient().Query(query)
		if err != nil {
			return fmt.Errorf("Query Error: %s", err.Error())
		}
		
		var text string
		line := new(string)
		for rows.Next() {
			rows.Scan(line)
			text += *line
		}
		if !strings.Contains(text, expected) {
			return fmt.Errorf("%s does not contain %s", text, expected)
		}

		return nil
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

	t.Run("add user to game", func(t *testing.T) {
		userid := "testuserid"
		err := dbc.SetUserLocation(userid, "1.0", "1.0")
		if err != nil {
			t.Error(err)
		}
		err = dbc.AddUser(gameName, userid)
		if err != nil {
			t.Error(err)
		}
		//queryContains(t, fmt.Sprintf("SELECT COUNT(*) FROM %s;", gameName), "1")
		err = queryContains(fmt.Sprintf("SELECT userid FROM %s;", gameName), userid)
		if err != nil {
			t.Error(err)
		}
		err = queryContains(fmt.Sprintf("SELECT location FROM %s;", gameName), userid)
		if err != nil {
			t.Error(err)
		}
	})
}
