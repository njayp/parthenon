package test

import (
	"context"
	"fmt"
	"io"
	"os"
	"testing"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/njayp/parthenon/pkg/db/cli"
	"github.com/njayp/parthenon/pkg/db/cli/mysqlCli"
	"github.com/njayp/parthenon/pkg/db/controller/games/spatialindex"
	"github.com/njayp/parthenon/pkg/db/controller/users/uuidindex"
)

type RmImageFunc = func(context.Context) error

// DockerRunMYSQL will run a MYSQL container on port 3306,
// then set a goroutine that will rm the container when ctx is cancelled
func DockerRunMYSQL(ctx context.Context) (RmImageFunc, error) {
	cli, err := client.NewClientWithOpts()
	if err != nil {
		fmt.Println("Unable to create docker client")
		panic(err)
	}

	imageName := "mysql:latest"
	reader, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	defer reader.Close()
	io.Copy(os.Stdout, reader)

	hostBinding := nat.PortBinding{
		HostIP:   "0.0.0.0",
		HostPort: "3306",
	}
	containerPort, err := nat.NewPort("tcp", "3306")
	if err != nil {
		panic("Unable to get the port")
	}

	portBinding := nat.PortMap{containerPort: []nat.PortBinding{hostBinding}}
	cont, err := cli.ContainerCreate(
		context.Background(),
		&container.Config{
			Image: imageName,
			Env:   []string{"MYSQL_ROOT_PASSWORD=password"},
		},
		&container.HostConfig{
			PortBindings: portBinding,
		}, nil, nil, "")
	if err != nil {
		panic(err)
	}

	return func(ctx context.Context) error {
		return cli.ContainerRemove(ctx, cont.ID, types.ContainerRemoveOptions{
			RemoveVolumes: true,
			Force:         true,
		})
	}, cli.ContainerStart(ctx, cont.ID, types.ContainerStartOptions{})
}

// setupDB runs MYSQL on a docker container and waits for it to start up.
// Returns mysqlCli, rmContainer, err
func setupDB(ctx context.Context) (cli.DBCli, RmImageFunc, error) {
	rm, err := DockerRunMYSQL(ctx)
	if err != nil {
		return nil, nil, err
	}

	cli, err := mysqlCli.NewMYSQLDBCli("localhost")
	if err != nil {
		rm(ctx)
		return nil, nil, err
	}

	// PingUntilConnect ctx, 60s timeout
	tCtx, cancel := context.WithTimeout(ctx, time.Second*60)
	defer cancel()
	err = cli.PingUntilConnect(tCtx)
	if err != nil {
		rm(ctx)
		return nil, nil, err
	}

	return cli, rm, nil
}

func setupDBnClients(ctx context.Context, t *testing.T) (*spatialindex.SpatialIndexDBC, *uuidindex.UUIDIndexDBC) {
	cli, rm, err := setupDB(ctx)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		rm(ctx)
	})

	sDBC, err := spatialindex.NewSpatialIndexDBC(ctx, cli)
	if err != nil {
		rm(ctx)
		t.Fatal(err)
	}

	uDBC, err := uuidindex.NewUUIDIndexDBC(ctx, cli)
	if err != nil {
		rm(ctx)
		t.Fatal(err)
	}

	return sDBC, uDBC
}
