package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/njayp/parthenon/pkg/db/controller/users/uuidindex"
)

func TestUsersFunctionality(t *testing.T) {
	ctx := context.Background()
	_, dbc := setupDBnClients(ctx, t)
	usersFunctionality(ctx, t, dbc)
}

func usersFunctionality(ctx context.Context, t *testing.T, dbc *uuidindex.UUIDIndexDBC) {
	t.Run("users functionality", func(t *testing.T) {
		t.Parallel()

		queryContains := func(query, expected string) error {
			rows, err := dbc.Client.QueryContext(ctx, query)
			if err != nil {
				return fmt.Errorf("Query Error: %s", err.Error())
			}

			return rowsContains(rows, expected)
		}

		t.Run("ensure tables", func(t *testing.T) {
			for _, tribe := range uuidindex.TRIBES {
				err := queryContains(fmt.Sprintf("SHOW TABLES LIKE '%s';", tribe), tribe)
				if err != nil {
					t.Error(err)
				}
			}
		})

		userid1 := uuid.New().String()
		tribe := uuidindex.TRIBES[0]

		t.Run("add user1", func(t *testing.T) {
			err := dbc.AddUser(ctx, tribe, userid1)
			if err != nil {
				t.Error(err)
			}
		})

		t.Run("get user1", func(t *testing.T) {
			row := dbc.GetUser(ctx, tribe, userid1)
			line := new(string)
			row.Scan(line)
			if *line != userid1 {
				t.Errorf("was expecting %s, got %s", userid1, *line)
			}
		})
	})
}
