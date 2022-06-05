package image

import (
	"context"

	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/docker"
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
	args := []string{
		// "--debug",
		image,
	}

	options := docker.RunOptions{
		Env: map[string]string{
			"DOCKER_CONTENT_TRUST": "1",
		},

		Volumes: []docker.ContainerMount{
			{
				Path:     "/var/run/docker.sock",
				HostPath: "/var/run/docker.sock",
			},
		},
	}

	return docker.RunInteractive(ctx, "goodwithtech/dockle", options, args...)
}
