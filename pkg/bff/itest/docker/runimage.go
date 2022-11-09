package runimage

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

// DockerRunMYSQLWithCtx will run a MYSQL container on port 3306,
// then set a goroutine that will rm the container when ctx is cancelled
func DockerRunMYSQLWithCtx(ctx context.Context) (func(ctx context.Context) error, error) {
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
