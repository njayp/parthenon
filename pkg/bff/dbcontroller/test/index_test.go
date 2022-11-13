package test

import (
	"context"
	"testing"
)

func TestIndex(t *testing.T) {
	ctx := context.Background()
	dbc, _ := setupDBnClients(ctx, t)

	gamesOptimization(ctx, t, dbc)
	gamesFunctionality(ctx, t, dbc)
}
