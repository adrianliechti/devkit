package image

import (
	"context"

	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/pkg/engine"
	"github.com/adrianliechti/go-cli"
)

var scanCommand = &cli.Command{
	Name:  "scan",
	Usage: "scan image vulnerabilies using trivy",

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

		return runTrivy(ctx, client, image)
	},
}

func runTrivy(ctx context.Context, client engine.Client, image string) error {
	container := engine.Container{
		Image: "aquasec/trivy:0.59.0",

		Args: []string{
			"--quiet",
			"image",
			image,
		},

		Mounts: []engine.ContainerMount{
			{
				Path:   "/root/.cache/",
				Volume: "trivy-cache",
			},
		},
	}

	return client.Run(ctx, container, engine.RunOptions{})
}
