package db2

import (
	"fmt"

	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/app/common"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/docker"
	"github.com/sethvargo/go-password/password"
)

func CreateCommand() *cli.Command {
	return &cli.Command{
		Name:  "create",
		Usage: "create instance",

		Flags: []cli.Flag{
			app.PortFlag(""),
		},

		Action: func(c *cli.Context) error {
			ctx := c.Context
			image := "ibmcom/db2:11.5.7.0a"

			port := app.MustPortOrRandom(c, "", 50000)

			db := "db"
			instance := "db2inst1"
			password := password.MustGenerate(10, 4, 0, false, false)

			options := docker.RunOptions{
				Labels: map[string]string{
					common.KindKey: DB2,
				},

				Privileged: true,
				Platform:   "linux/amd64",

				Env: map[string]string{
					"LICENSE": "accept",

					"DBNAME": db,

					"DB2INST1_PASSWORD": password,
				},

				Ports: map[int]int{
					port: 50000,
				},

				// Volumes: map[string]string{
				// 	name: /database",
				// },
			}

			if err := docker.Run(ctx, image, options); err != nil {
				return err
			}

			cli.Table([]string{"Name", "Value"}, [][]string{
				{"Host", fmt.Sprintf("localhost:%d", port)},
				{"Instance", instance},
				{"Database", db},
				{"Password", password},
			})

			return nil
		},
	}
}
