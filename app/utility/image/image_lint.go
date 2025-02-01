package image

import (
	"context"

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

		cli.MustRun("Pulling Image...", func() error {
			client.Pull(ctx, image, "", engine.PullOptions{})
			return nil
		})

		return runDockle(ctx, client, image)
	},
}

func runDockle(ctx context.Context, client engine.Client, image string) error {
	container := engine.Container{
		Image: "goodwithtech/dockle:v0.4.15",

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

	return client.Run(ctx, container, engine.RunOptions{})
}
