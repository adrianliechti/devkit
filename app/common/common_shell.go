package common

import (
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/docker"
)

func ShellCommand(kind, shell string, arg ...string) *cli.Command {
	return &cli.Command{
		Name:  "shell",
		Usage: "run shell in instance (" + shell + ")",

		Action: func(c *cli.Context) error {
			ctx := c.Context
			container := MustContainer(ctx, kind)

			return docker.ExecInteractive(ctx, container, docker.ExecOptions{}, shell, arg...)
		},
	}
}
