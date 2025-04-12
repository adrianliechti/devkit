package image

import (
	"context"

	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/pkg/engine"
	"github.com/adrianliechti/go-cli"
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

		cli.MustRun("Pulling Image...", func() error {
			client.Pull(ctx, image, "", engine.PullOptions{})
			return nil
		})

		return runSyft(ctx, client, image)
	},
}

func runSyft(ctx context.Context, client engine.Client, image string) error {
	container := engine.Container{
		Image: "anchore/syft",

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
