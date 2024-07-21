package image

import (
	"context"

	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/engine"
)

var browseCommand = &cli.Command{
	Name:  "browse",
	Usage: "browse image using dive (interactive)",

	Flags: []cli.Flag{
		ImageFlag,
	},

	Action: func(ctx context.Context, cmd *cli.Command) error {
		client := app.MustClient(ctx, cmd)
		image := MustImage(ctx, cmd)

		return runDive(ctx, client, image)
	},
}

func runDive(ctx context.Context, client engine.Client, image string) error {
	container := engine.Container{
		Image: "wagoodman/dive:v0.12",

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
