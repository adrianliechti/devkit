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

	Action: func(c *cli.Context) error {
		image := MustImage(c)
		return runTrivy(c.Context, image)
	},
}

func runTrivy(ctx context.Context, image string) error {
	options := docker.RunOptions{
		Env: map[string]string{},

		Volumes: []engine.ContainerMount{
			{
				Path:   "/root/.cache/",
				Volume: "trivy-cache",
			},
		},
	}

	return docker.RunInteractive(ctx, "aquasec/trivy", options, image)
}
