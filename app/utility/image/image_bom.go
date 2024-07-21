package image

import (
	"context"

	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/engine"
)

var bomCommand = &cli.Command{
	Name:  "bom",
	Usage: "show image bill of material using syft",

	Flags: []cli.Flag{
		ImageFlag,
	},

	Action: func(ctx context.Context, cmd *cli.Command) error {
		client := app.MustClient(ctx, cmd)
		image := MustImage(ctx, cmd)

		return runSyft(ctx, client, image)
	},
}

func runSyft(ctx context.Context, client engine.Client, image string) error {
	container := engine.Container{
		Image: "anchore/syft:v1.9.0",

		Args: []string{
			"-o", "syft-table",
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
