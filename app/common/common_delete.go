package common

import (
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/docker"
)

func DeleteCommand(kind string) *cli.Command {
	return &cli.Command{
		Name:  "delete",
		Usage: "delete instance",

		Action: func(c *cli.Context) error {
			container := MustContainer(c.Context, kind)

			return docker.Delete(c.Context, container, docker.DeleteOptions{
				Force:   true,
				Volumes: true,
			})
		},
	}
}
