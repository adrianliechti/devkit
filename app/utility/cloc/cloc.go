package cloc

import (
	"context"
	"os"

	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/app/utility"
	"github.com/adrianliechti/devkit/pkg/engine"
	"github.com/adrianliechti/go-cli"
)

var Command = &cli.Command{
	Name:  "cloc",
	Usage: "count lines of code",

	Category: utility.Category,

	Action: func(ctx context.Context, cmd *cli.Command) error {
		client := app.MustClient(ctx, cmd)

		path, err := os.Getwd()

		if err != nil {
			return err
		}

		return runCloc(ctx, client, path)
	},
}

func runCloc(ctx context.Context, client engine.Client, path string) error {
	container := engine.Container{
		Image: "aldanial/cloc",

		Args: []string{
			"--quiet",
			"--hide-rate",
			"/src",
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
