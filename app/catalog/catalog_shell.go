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

			name := MustContainer(ctx, kind)

			return docker.ExecInteractive(ctx, name, docker.ExecOptions{}, shell)
		},
	}
}
