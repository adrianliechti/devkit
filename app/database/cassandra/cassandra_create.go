package cassandra

import (
	"fmt"

	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/app/common"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/docker"
)

func CreateCommand() *cli.Command {
	return &cli.Command{
		Name:  "create",
		Usage: "create instance",

		Flags: []cli.Flag{
			app.PortFlag,
		},

		Action: func(c *cli.Context) error {
			ctx := c.Context
			image := "cassandra:4"

			target := 9042
			port := app.MustPortOrRandom(c, target)

			options := docker.RunOptions{
				Labels: map[string]string{
					common.KindKey: Cassandra,
				},

				Env: map[string]string{},

				Ports: map[int]int{
					port: target,
				},

				// Volumes: map[string]string{
				// 	name: /var/lib/cassandra",
				// },
			}

			if err := docker.Run(ctx, image, options); err != nil {
				return err
			}

			cli.Table([]string{"Name", "Value"}, [][]string{
				{"Host", fmt.Sprintf("localhost:%d", port)},
			})

			return nil
		},
	}
}
