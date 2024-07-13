package image

import (
	"context"

	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/docker"
	"github.com/adrianliechti/devkit/pkg/engine"
)

var browseCommand = &cli.Command{
	Name:  "browse",
	Usage: "browse image using dive (interactive)",

	Flags: []cli.Flag{
		ImageFlag,
	},

	Action: func(ctx context.Context, cmd *cli.Command) error {
		image := MustImage(ctx, cmd)
		return runDive(ctx, image)
	},
}

func runDive(ctx context.Context, image string) error {
	tool := "wagoodman/dive:v0.12"

	options := docker.RunOptions{
		Volumes: []engine.ContainerMount{
			{
				Path:     "/var/run/docker.sock",
				HostPath: "/var/run/docker.sock",
			},
		},
	}

	return docker.RunInteractive(ctx, tool, options, image)
}
