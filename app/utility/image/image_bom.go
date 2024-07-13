package image

import (
	"context"

	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/docker"
	"github.com/adrianliechti/devkit/pkg/engine"
)

var bomCommand = &cli.Command{
	Name:  "bom",
	Usage: "show image bill of material using syft",

	Flags: []cli.Flag{
		ImageFlag,
	},

	Action: func(c *cli.Context) error {
		image := MustImage(c)
		return runSyft(c.Context, image)
	},
}

func runSyft(ctx context.Context, image string) error {
	tool := "anchore/syft:v1.9.0"

	args := []string{
		"-o", "syft-table",
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
