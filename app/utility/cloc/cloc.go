package cloc

import (
	"context"
	"os"

	"github.com/adrianliechti/devkit/app/utility"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/docker"
	"github.com/adrianliechti/devkit/pkg/engine"
)

var Command = &cli.Command{
	Name:  "cloc",
	Usage: "count lines of code",

	Category: utility.Category,

	Action: func(ctx context.Context, cmd *cli.Command) error {
		return runCloc(ctx)
	},
}

func runCloc(ctx context.Context) error {
	wd, err := os.Getwd()

	if err != nil {
		return err
	}

	args := []string{
		"--quiet",
		"--hide-rate",
		"/src",
	}

	options := docker.RunOptions{
		Volumes: []engine.ContainerMount{
			{
				Path:     "/src",
				HostPath: wd,
			},
		},
	}

	return docker.RunInteractive(ctx, "aldanial/cloc", options, args...)
}
