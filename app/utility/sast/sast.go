package sast

import (
	"context"
	"os"

	"github.com/adrianliechti/devkit/app/utility"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/docker"
	"github.com/adrianliechti/devkit/pkg/engine"
)

var Command = &cli.Command{
	Name:  "sast",
	Usage: "static analysis for many languages (using semgrep)",

	Category: utility.Category,

	Action: func(ctx context.Context, cmd *cli.Command) error {
		return runSAST(ctx)
	},
}

func runSAST(ctx context.Context) error {
	wd, err := os.Getwd()

	if err != nil {
		return err
	}

	image := "semgrep/semgrep:1.79.0"

	args := []string{
		"semgrep",
		"scan",
		"--metrics=on",
		"--config", "auto",
		"--oss-only",
		"--quiet",
	}

	options := docker.RunOptions{
		Dir: "/src",

		Volumes: []engine.ContainerMount{
			{
				Path:     "/src",
				HostPath: wd,
			},
		},
	}

	return docker.RunInteractive(ctx, image, options, args...)
}
