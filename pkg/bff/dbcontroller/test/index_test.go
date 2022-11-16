package test

import (
	"context"
	"testing"
)

// TODO currently we have one games dbc and one users dbc, so we test those directly.
// In the future, run a test matrix where we run the same tests for all games dbcs etc
func TestIndex(t *testing.T) {
	ctx := context.Background()
	sDBC, uDBC := setupDBnClients(ctx, t)

	gamesOptimization(ctx, t, sDBC)
	gamesFunctionality(ctx, t, sDBC)
	usersFunctionality(ctx, t, uDBC)
	usersOptimization(ctx, t, uDBC)
}
