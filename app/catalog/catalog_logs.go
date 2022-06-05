package catalog

import (
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/docker"
)

func LogsCommand(kind string) *cli.Command {
	return &cli.Command{
		Name:  "logs",
		Usage: "show instance logs",

		Action: func(c *cli.Context) error {
			ctx := c.Context
			container := MustContainer(ctx, kind)

			return docker.Logs(ctx, container, docker.LogsOptions{
				Follow: true,
			})
		},
	}
}
