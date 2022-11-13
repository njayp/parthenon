package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/njayp/parthenon/pkg/bff/dbcontroller/users/uuidindex"
)

func TestUsersFunctionality(t *testing.T) {
	ctx := context.Background()
	_, dbc := setupDBnClients(ctx, t)
	usersFunctionality(ctx, t, dbc)
}

func usersFunctionality(ctx context.Context, t *testing.T, dbc *uuidindex.UUIDIndexDBC) {
	t.Run("functionality", func(t *testing.T) {
		t.Parallel()

		queryContains := func(query, expected string) error {
			rows, err := dbc.Client.QueryContext(ctx, query)
			if err != nil {
				return fmt.Errorf("Query Error: %s", err.Error())
			}

			return rowsContains(rows, expected)
		}

		t.Run("ensure tables", func(t *testing.T) {
			t.Parallel()

			for _, tribe := range uuidindex.TRIBES {
				err := queryContains(fmt.Sprintf("SHOW TABLES LIKE '%s';", tribe), tribe)
				if err != nil {
					t.Error(err)
				}
			}
		})
	})
}
