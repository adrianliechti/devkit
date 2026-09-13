package openapi

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/pkg/engine"
	"github.com/adrianliechti/go-cli"
)

var mockCommand = &cli.Command{
	Name:  "mock",
	Usage: "mock openapi server",

	Flags: []cli.Flag{
		app.PortFlag(""),
	},

	Action: func(ctx context.Context, cmd *cli.Command) error {
		client := app.MustClient(ctx, cmd)

		port := app.MustPortOrRandom(ctx, cmd, "", 4010)
		path := cli.MustFile("Select Swagger/OpenAPI schema", []string{".json", ".yaml"})

		return runMock(ctx, client, path, port)
	},
}

func runMock(ctx context.Context, client engine.Client, path string, port int) error {
	path, err := filepath.Abs(path)

	if err != nil {
		return err
	}

	dir, file := filepath.Split(path)

	args := []string{
		"mock",
		"--host", "0.0.0.0",
		"--port", fmt.Sprintf("%d", port),

		"/src/" + file,
	}

	container := engine.Container{
		Image: "stoplight/prism:5",
		Args:  args,

		Ports: []engine.ContainerPort{
			{
				Port:     port,
				HostPort: port,
			},
		},

		Mounts: []engine.ContainerMount{
			{
				Path:     "/src",
				HostPath: dir,
			},
		},
	}

	return client.Run(ctx, container, engine.RunOptions{})
}
