package sast

import (
	"context"
	"os"

	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/app/utility"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/engine"
)

var Command = &cli.Command{
	Name:  "sast",
	Usage: "static analysis for many languages (using semgrep)",

	Category: utility.Category,

	Action: func(ctx context.Context, cmd *cli.Command) error {
		client := app.MustClient(ctx, cmd)

		path, err := os.Getwd()

		if err != nil {
			return err
		}

		return runSAST(ctx, client, path)
	},
}

func runSAST(ctx context.Context, client engine.Client, path string) error {
	container := engine.Container{
		Image: "semgrep/semgrep:1.79.0",

		Dir: "/src",

		Args: []string{
			"semgrep",
			"scan",
			"--metrics=on",
			"--config", "auto",
			"--oss-only",
			"--quiet",
		},

		Mounts: []engine.ContainerMount{
			{
				Path:     "/src",
				HostPath: path,
			},
		},
	}

	return client.Run(ctx, container, engine.RunOptions{})
}
