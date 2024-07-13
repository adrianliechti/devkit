package image

import (
	"context"

	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/docker"
	"github.com/adrianliechti/devkit/pkg/engine"
)

var scanCommand = &cli.Command{
	Name:  "scan",
	Usage: "scan image vulnerabilies using trivy",

	Flags: []cli.Flag{
		ImageFlag,
	},

	Action: func(ctx context.Context, cmd *cli.Command) error {
		image := MustImage(ctx, cmd)
		return runTrivy(ctx, image)
	},
}

func runTrivy(ctx context.Context, image string) error {
	tool := "aquasec/trivy:0.53.0"

	options := docker.RunOptions{
		Env: map[string]string{},

		Volumes: []engine.ContainerMount{
			{
				Path:   "/root/.cache/",
				Volume: "trivy-cache",
			},
		},
	}

	return docker.RunInteractive(ctx, tool, options, "image", image)
}
