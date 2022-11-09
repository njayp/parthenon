package test

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/njayp/parthenon/pkg/bff/dbcli/mysqlCli"
	runimage "github.com/njayp/parthenon/pkg/bff/itest/docker"
)

func TestPing(t *testing.T) {
	fatalErr := func(err error) {
		if err != nil {
			t.Fatal(err.Error())
		}
	}

	ctx := context.Background()
	rm, err := runimage.DockerRunMYSQLWithCtx(ctx)
	fatalErr(err)
	defer rm(ctx)

	cli, err := mysqlCli.NewMYSQLDBCli("localhost")
	fatalErr(err)

	t.Run("ping until starts up", func(t *testing.T) {
		tCtx, cancel := context.WithTimeout(ctx, time.Second*30)
		defer cancel()
		err = cli.PingUntilConnect(tCtx)
		fatalErr(err)
	})

	t.Run("make DB, connect, and verify", func(t *testing.T) {
		dbname := "testo"
		testoDB, err := cli.EnsureDBandCli(dbname)
		fatalErr(err)
		err = testoDB.Ping()
		fatalErr(err)
		rows, err := testoDB.Query("SELECT DATABASE()")
		fatalErr(err)
		var text string
		row := new(string)
		for rows.Next() {
			err := rows.Scan(row)
			fatalErr(err)
			text += *row
		}
		if !strings.Contains(text, dbname) {
			t.Errorf("%s does not contain %s", text, dbname)
		}
	})
}
