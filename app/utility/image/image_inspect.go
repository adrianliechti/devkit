package image

import (
	"context"

	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/docker"
	"github.com/adrianliechti/devkit/pkg/engine"
)

var inspectCommand = &cli.Command{
	Name:  "inspect",
	Usage: "inspect image using whaler",

	Flags: []cli.Flag{
		ImageFlag,
		&cli.BoolFlag{
			Name:  "verbose",
			Usage: "verbose output",
		},
	},

	Action: func(c *cli.Context) error {
		image := MustImage(c)
		return runWhaler(c.Context, image, c.Bool("verbose"))
	},
}

func runWhaler(ctx context.Context, image string, verbose bool) error {
	tool := "pegleg/whaler"

	args := []string{}

	if verbose {
		args = append(args, "-v")
	}

	args = append(args, image)

	options := docker.RunOptions{
		Volumes: []engine.ContainerMount{
			{
				Path:     "/var/run/docker.sock",
				HostPath: "/var/run/docker.sock",
			},
		},
	}

	return docker.RunInteractive(ctx, tool, options, args...)
}
