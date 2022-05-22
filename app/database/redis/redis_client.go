package redis

import (
	"github.com/adrianliechti/devkit/app/common"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/docker"
)

func ClientCommand() *cli.Command {
	return &cli.Command{
		Name:  "cli",
		Usage: "run redis-cli in instance",

		Action: func(c *cli.Context) error {
			ctx := c.Context
			container := common.MustContainer(ctx, Redis)

			options := docker.ExecOptions{}

			return docker.ExecInteractive(ctx, container, options,
				"/bin/bash", "-c",
				"redis-cli",
			)
		},
	}
}
