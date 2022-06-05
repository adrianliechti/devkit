package catalog

import (
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/docker"
)

type ClientHandler = func() (shell string, args []string)

func ClientCommand(kind string, h ClientHandler) *cli.Command {
	return &cli.Command{
		Name:  "cli",
		Usage: "run client in instance",

		Action: func(c *cli.Context) error {
			ctx := c.Context

			name := MustContainer(ctx, kind)

			command, args := h()

			return docker.ExecInteractive(ctx, name, docker.ExecOptions{}, command, args...)
		},
	}
}
