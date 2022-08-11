package image

import (
	"context"

	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/docker"
	"github.com/adrianliechti/devkit/pkg/engine"
)

var lintCommand = &cli.Command{
	Name:  "lint",
	Usage: "lint Dockerfile using dockle",

	Flags: []cli.Flag{
		ImageFlag,
	},

	Action: func(c *cli.Context) error {
		image := MustImage(c)
		return runDockle(c.Context, image)
	},
}

func runDockle(ctx context.Context, image string) error {
	tool := "goodwithtech/dockle:v0.4.6"

	args := []string{
		// "--debug",
		image,
	}

	options := docker.RunOptions{
		Env: map[string]string{
			"DOCKER_CONTENT_TRUST": "1",
		},

		Volumes: []engine.ContainerMount{
			{
				Path:     "/var/run/docker.sock",
				HostPath: "/var/run/docker.sock",
			},
		},
	}

	return docker.RunInteractive(ctx, tool, options, args...)
}
