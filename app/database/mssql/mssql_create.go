package mssql

import (
	"fmt"
	"runtime"

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
			image := "mcr.microsoft.com/mssql/server:2019-latest"

			if runtime.GOARCH == "arm64" {
				image = "mcr.microsoft.com/azure-sql-edge"
			}

			port := app.MustPortOrRandom(c, "", 1433)

			username := "sa"
			password := password.MustGenerate(10, 4, 0, false, false)

			options := docker.RunOptions{
				Labels: map[string]string{
					common.KindKey: MSSQL,
				},

				Env: map[string]string{
					"ACCEPT_EULA": "Y",
					"MSSQL_PID":   "Developer",
					"SA_PASSWORD": password,
				},

				// Env: map[string]string{
				// 	"ACCEPT_EULA":       "1",
				// 	"MSSQL_PID":         "Developer",
				// 	"MSSQL_SA_PASSWORD": password,
				// },

				Ports: map[int]int{
					port: 1433,
				},

				// Volumes: map[string]string{
				// 	name: "/var/opt/mssql",
				// },
			}

			if err := docker.Run(ctx, image, options); err != nil {
				return err
			}

			cli.Table([]string{"Name", "Value"}, [][]string{
				{"Host", fmt.Sprintf("localhost:%d", port)},
				{"Username", username},
				{"Password", password},
			})

			return nil
		},
	}
}
