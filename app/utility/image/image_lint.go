package image

import (
	"context"
	"os"

	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/engine"
)

var lintCommand = &cli.Command{
	Name:  "lint",
	Usage: "lint Dockerfile using dockle",

	Flags: []cli.Flag{
		ImageFlag,
	},

	Action: func(ctx context.Context, cmd *cli.Command) error {
		client := app.MustClient(ctx, cmd)
		image := MustImage(ctx, cmd)

		return runDockle(ctx, client, image)
	},
}

func runDockle(ctx context.Context, client engine.Client, image string) error {
	container := engine.Container{
		Image: "goodwithtech/dockle:v0.4.14",

		Env: map[string]string{
			"DOCKER_CONTENT_TRUST": "1",
		},

		Args: []string{
			image,
		},

		Mounts: []engine.ContainerMount{
			{
				Path:     "/var/run/docker.sock",
				HostPath: "/var/run/docker.sock",
			},
		},
	}

	return client.Run(ctx, container, engine.RunOptions{
		TTY:         true,
		Interactive: true,

		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	})
}
