package catalog

import (
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/docker"
)

func ShellCommand(kind string, shell string) *cli.Command {
	return &cli.Command{
		Name:  "shell",
		Usage: "run shell in instance",

		Action: func(c *cli.Context) error {
			ctx := c.Context
			container := MustContainer(ctx, kind)

			return docker.ExecInteractive(ctx, container.Name, docker.ExecOptions{}, shell)
		},
	}
}
