package grafana

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
			image := "grafana/grafana"

			port := app.MustPortOrRandom(c, "", 3000)

			username := "admin"
			password := password.MustGenerate(10, 4, 0, false, false)

			options := docker.RunOptions{
				Labels: map[string]string{
					common.KindKey: Grafana,
				},

				Env: map[string]string{
					"GF_SECURITY_ADMIN_USER":     username,
					"GF_SECURITY_ADMIN_PASSWORD": password,
				},

				Ports: map[int]int{
					port: 3000,
				},

				// Volumes: map[string]string{
				// 	name: "/var/lib/grafana",
				// },
			}

			if err := docker.Run(ctx, image, options); err != nil {
				return err
			}

			cli.Table([]string{"Name", "Value"}, [][]string{
				{"URL", fmt.Sprintf("http://localhost:%d", port)},
				{"Username", username},
				{"Password", password},
			})

			return nil
		},
	}
}
