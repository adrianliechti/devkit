package minio

import (
	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/app/common"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/docker"
)

func ConsoleCommand() *cli.Command {
	return &cli.Command{
		Name:  "console",
		Usage: "open web console",

		Flags: []cli.Flag{
			app.PortFlag,
		},

		Action: func(c *cli.Context) error {
			ctx := c.Context
			container := common.MustContainer(ctx, MinIO)

			port := app.MustPortOrRandom(c, 9090)
			target := 9001

			return docker.PortForward(c.Context, container, port, target)
		},
	}
}
