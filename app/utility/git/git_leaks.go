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

	Action: func(ctx context.Context, cmd *cli.Command) error {
		return leaks(ctx)
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

	return docker.RunInteractive(ctx, "zricethezav/gitleaks:v8.18.4", options, args...)
}
