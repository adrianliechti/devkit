package catalog

import (
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/container"
)

type ConsoleHandler = func() container.ContainerPort

func ConsoleCommand(kind string, h ConsoleHandler) *cli.Command {
	return &cli.Command{
		Name:  "console",
		Usage: "open web console",

		Action: func(c *cli.Context) error {
			ctx := c.Context
			container := MustContainer(ctx, kind)

			port := h()

			_ = port

			//return docker.ExecInteractive(ctx, container, docker.ExecOptions{}, shell, arg...)
			_ = container
			return nil
		},
	}
}
