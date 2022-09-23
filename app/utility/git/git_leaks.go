package git

import (
	"context"
	"os"

	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/docker"
	"github.com/adrianliechti/devkit/pkg/engine"
)

var leaksCommand = &cli.Command{
	Name:  "leaks",
	Usage: "find leaks in repository",

	Action: func(c *cli.Context) error {
		return leaks(c.Context)
	},
}

func leaks(ctx context.Context) error {
	path, err := os.Getwd()

	if err != nil {
		return err
	}

	options := docker.RunOptions{
		User: "root",

		Volumes: []engine.ContainerMount{
			{
				Path:     "/src",
				HostPath: path,
			},
		},
	}

	args := []string{
		"detect",
		"--source=/src",
		"--no-banner",
		"-v",
		//"--config=/config",
	}

	return docker.RunInteractive(ctx, "zricethezav/gitleaks:v8.13.0", options, args...)
}
