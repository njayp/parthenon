package test

import (
	"context"
	"time"

	"github.com/njayp/parthenon/pkg/bff/dbcli"
	"github.com/njayp/parthenon/pkg/bff/dbcli/mysqlCli"
	runimage "github.com/njayp/parthenon/pkg/bff/itest/docker"
)

// SetupDB runs MYSQL on a docker container and waits for it to start up.
// Returns mysqlCli, rmContainer, err
func SetupDB(ctx context.Context) (dbcli.DBCli, func(context.Context) error, error) {
	rm, err := runimage.DockerRunMYSQLWithCtx(ctx)
	if err != nil {
		return nil, nil, err
	}

	cli, err := mysqlCli.NewMYSQLDBCli("localhost")
	if err != nil {
		return nil, nil, err
	}

	// Ping ctx, 30s timeout
	tCtx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()
	err = cli.PingUntilConnect(tCtx)
	if err != nil {
		rm(ctx)
		return nil, nil, err
	}

	return cli, rm, nil
}
