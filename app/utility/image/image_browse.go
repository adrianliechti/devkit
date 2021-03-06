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

	Action: func(c *cli.Context) error {
		image := MustImage(c)
		return runDive(c.Context, image)
	},
}

func runDive(ctx context.Context, image string) error {
	options := docker.RunOptions{
		Volumes: []engine.ContainerMount{
			{
				Path:     "/var/run/docker.sock",
				HostPath: "/var/run/docker.sock",
			},
		},
	}

	return docker.RunInteractive(ctx, "wagoodman/dive", options, image)
}
