package cluster

import (
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/kind"
)

func DeleteCommand() *cli.Command {
	return &cli.Command{
		Name:  "delete",
		Usage: "delete instance",

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "name",
				Usage: "cluster name",
			},
		},

		Action: func(c *cli.Context) error {
			name := c.String("name")

			if name == "" {
				name = MustCluster(c.Context)
			}

			return kind.Delete(c.Context, name)
		},
	}
}
