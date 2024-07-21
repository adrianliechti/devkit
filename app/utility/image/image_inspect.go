package image

import (
	"context"
	"os"

	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/engine"
)

var inspectCommand = &cli.Command{
	Name:  "inspect",
	Usage: "inspect image using whaler",

	Flags: []cli.Flag{
		ImageFlag,
	},

	Action: func(ctx context.Context, cmd *cli.Command) error {
		client := app.MustClient(ctx, cmd)
		image := MustImage(ctx, cmd)

		return runWhaler(ctx, client, image)
	},
}

func runWhaler(ctx context.Context, client engine.Client, image string) error {
	container := engine.Container{
		Image: "pegleg/whaler",

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
