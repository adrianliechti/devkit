package immudb

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
			app.PortFlag,
		},

		Action: func(c *cli.Context) error {
			ctx := c.Context
			image := "codenotary/immudb"

			port := app.MustPortOrRandom(c, 3322)
			consolePort := app.MustRandomPort(c, port+1)

			username := "immudb"
			password := password.MustGenerate(10, 4, 0, false, false)

			options := docker.RunOptions{
				Labels: map[string]string{
					common.KindKey: ImmuDB,
				},

				Env: map[string]string{
					"IMMUDB_ADDRESS": "0.0.0.0",

					"IMMUDB_ADMIN_PASSWORD": password,
				},

				Ports: map[int]int{
					port:        3322,
					consolePort: 8080,
				},

				// Volumes: map[string]string{
				// 	name: "/var/lib/immudb",
				// },
			}

			if err := docker.Run(ctx, image, options); err != nil {
				return err
			}

			cli.Table([]string{"Name", "Value"}, [][]string{
				{"Host", fmt.Sprintf("localhost:%d", port)},
				{"Username", username},
				{"Password", password},
				{"URL", fmt.Sprintf("http://localhost:%d", port)},
				{"Console", fmt.Sprintf("http://localhost:%d", consolePort)},
			})

			return nil
		},
	}
}
