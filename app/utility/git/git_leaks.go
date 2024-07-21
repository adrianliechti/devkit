package git

import (
	"context"
	"os"

	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/engine"
)

var leaksCommand = &cli.Command{
	Name:  "leaks",
	Usage: "find leaks in repository",

	Action: func(ctx context.Context, cmd *cli.Command) error {
		client := app.MustClient(ctx, cmd)

		wd, err := os.Getwd()

		if err != nil {
			return err
		}

		return leaks(ctx, client, wd)
	},
}

func leaks(ctx context.Context, client engine.Client, path string) error {
	container := engine.Container{
		Image: "zricethezav/gitleaks:v8.18.4",

		RunAsUser: "root",

		Args: []string{
			"detect",
			"--source=/src",
			"--no-banner",
			"-v",
			//"--config=/config",
		},

		Mounts: []engine.ContainerMount{
			{
				Path:     "/src",
				HostPath: path,
			},
		},
	}

	return client.Run(ctx, container, engine.RunOptions{
		TTY:         true,
		Interactive: true,

		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	})
}
