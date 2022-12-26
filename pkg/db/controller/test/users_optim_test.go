package test

import (
	"context"
	"database/sql"

	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/njayp/parthenon/pkg/db/controller/users/uuidindex"
)

func TestUsersOptimization(t *testing.T) {
	ctx := context.Background()
	_, dbc := setupDBnClients(ctx, t)
	usersOptimization(ctx, t, dbc)
}

func usersOptimization(ctx context.Context, t *testing.T, dbc *uuidindex.UUIDIndexDBC) {
	t.Run("games optimization", func(t *testing.T) {
		t.Parallel()

		// add user to prime queries
		userid := uuid.New().String()
		userTribe := uuidindex.TRIBES[0]
		dbc.AddUser(ctx, userTribe, userid)

		t.Run("use EXPLAIN on search query. Ensure it uses the index", func(t *testing.T) {
			query := fmt.Sprintf("EXPLAIN "+uuidindex.GET_QUERY, userTribe, userid)
			row := dbc.Client.QueryRow(query)

			key := new(sql.NullString)
			possibleKeys := new(sql.NullString)
			keyLen := new(sql.NullString)
			searchtype := new(sql.NullString)
			selectType := new(sql.NullString)
			err := row.Scan(new(sql.NullString), selectType, new(sql.NullString), new(sql.NullString), searchtype, possibleKeys, key, keyLen, new(sql.NullString), new(sql.NullString), new(sql.NullString), new(sql.NullString))
			if err != nil {
				t.Error(err)
			}
			if key.String != "PRIMARY" || selectType.String != "SIMPLE" || searchtype.String != "const" {
				t.Error(fmt.Errorf("\nquery type: %s\nsearch type: %s\npossible: %s\nkey: %s\nkey len:%s", selectType.String, searchtype.String, possibleKeys.String, key.String, keyLen.String))
			}
		})
	})
}
