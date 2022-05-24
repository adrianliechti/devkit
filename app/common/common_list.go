package common

import (
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/docker"
)

func ListCommand(kind string) *cli.Command {
	return &cli.Command{
		Name:  "list",
		Usage: "list instances",

		Action: func(c *cli.Context) error {
			list, err := docker.List(c.Context, docker.ListOptions{
				All: true,

				Filter: []string{
					"label=" + KindKey + "=" + kind,
				},
			})

			if err != nil {
				return err
			}

			for _, c := range list {
				name := c.Names[0]
				cli.Info(name)
			}

			return nil
		},
	}
}
