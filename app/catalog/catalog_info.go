package catalog

import (
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/container"
)

type InfoHandler = func(container *container.Container) map[string]string

func InfoCommand(kind string, h InfoHandler) *cli.Command {
	return &cli.Command{
		Name:  "info",
		Usage: "display instance info",

		Action: func(c *cli.Context) error {
			ctx := c.Context
			container := MustContainer(ctx, kind)

			info := h(&container)

			printMapTable(info)
			return nil
		},
	}
}
