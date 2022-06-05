package image

import (
	"context"

	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/docker"
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
	args := []string{}

	if verbose {
		args = append(args, "-v")
	}

	args = append(args, image)

	options := docker.RunOptions{
		Volumes: []docker.ContainerMount{
			{
				Path:     "/var/run/docker.sock",
				HostPath: "/var/run/docker.sock",
			},
		},
	}

	return docker.RunInteractive(ctx, "pegleg/whaler", options, args...)
}
