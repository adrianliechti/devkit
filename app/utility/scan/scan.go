package scan

import (
	"context"
	"os"

	"github.com/adrianliechti/devkit/app/utility"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/docker"
	"github.com/adrianliechti/devkit/pkg/engine"
)

var Command = &cli.Command{
	Name:  "scan",
	Usage: "vulnerability scanning (using trivy)",

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

	image := "aquasec/trivy:0.53.0"

	args := []string{
		"filesystem",
		"--scanners", "vuln,misconfig,secret",
		"/src",
	}

	options := docker.RunOptions{
		Dir: "/src",

		Volumes: []engine.ContainerMount{
			{
				Path:     "/src",
				HostPath: wd,
			},
			{
				Path:   "/root/.cache/",
				Volume: "trivy-cache",
			},
		},
	}

	return docker.RunInteractive(ctx, image, options, args...)
}
