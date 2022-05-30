package influxdb

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
			image := "influxdb:2.2"

			port := app.MustPortOrRandom(c, "", 8086)

			organization := "default"
			bucket := "default"

			username := "admin"
			token := password.MustGenerate(10, 4, 0, false, false)
			password := password.MustGenerate(10, 4, 0, false, false)

			options := docker.RunOptions{
				Labels: map[string]string{
					common.KindKey: InfluxDB,
				},

				Env: map[string]string{
					"DOCKER_INFLUXDB_INIT_MODE": "setup",

					"DOCKER_INFLUXDB_INIT_ORG":    organization,
					"DOCKER_INFLUXDB_INIT_BUCKET": bucket,

					"DOCKER_INFLUXDB_INIT_USERNAME": username,
					"DOCKER_INFLUXDB_INIT_PASSWORD": password,

					"DOCKER_INFLUXDB_INIT_ADMIN_TOKEN": token,
				},

				Ports: map[int]int{
					port: 8086,
				},

				// Volumes: map[string]string{
				// 	name: "/var/lib/influxdb2",
				// },
			}

			if err := docker.Run(ctx, image, options); err != nil {
				return err
			}

			cli.Table([]string{"Name", "Value"}, [][]string{
				{"URL", fmt.Sprintf("http://localhost:%d", port)},
				{"organization", organization},
				{"bucket", bucket},
				{"Username", username},
				{"Password", password},
				{"Token", token},
			})

			return nil
		},
	}
}
