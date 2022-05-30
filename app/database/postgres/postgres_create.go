package postgres

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
			image := "postgres:14-bullseye"

			port := app.MustPortOrRandom(c, "", 5432)

			database := "postgres"
			username := "postgres"
			password := password.MustGenerate(10, 4, 0, false, false)

			options := docker.RunOptions{
				Labels: map[string]string{
					common.KindKey: PostgreSQL,
				},

				Env: map[string]string{
					"POSTGRES_DB":       database,
					"POSTGRES_USER":     username,
					"POSTGRES_PASSWORD": password,
				},

				Ports: map[int]int{
					port: 5432,
				},

				// Volumes: map[string]string{
				// 	name: "/var/lib/postgresql/data",
				// },
			}

			if err := docker.Run(ctx, image, options); err != nil {
				return err
			}

			cli.Table([]string{"Name", "Value"}, [][]string{
				{"Host", fmt.Sprintf("localhost:%d", port)},
				{"database", database},
				{"Username", username},
				{"Password", password},
			})

			return nil
		},
	}
}
