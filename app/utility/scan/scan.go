package scan

import (
	"context"
	"os"

	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/app/utility"
	"github.com/adrianliechti/devkit/pkg/engine"
	"github.com/adrianliechti/go-cli"
)

var Command = &cli.Command{
	Name:  "scan",
	Usage: "vulnerability scanning (using trivy)",

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
		Image: "aquasec/trivy:0.59.0",

		Dir: "/src",

		Args: []string{
			"--quiet",
			"filesystem",
			"--scanners", "vuln,misconfig,secret",
			"/src",
		},

		Mounts: []engine.ContainerMount{
			{
				Path:     "/src",
				HostPath: path,
			},
			{
				Path:   "/root/.cache/",
				Volume: "trivy-cache",
			},
		},
	}

	return client.Run(ctx, container, engine.RunOptions{})
}
