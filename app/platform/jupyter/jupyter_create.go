package jupyter

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
			image := "jupyter/datascience-notebook"

			port := app.MustPortOrRandom(c, "", 8888)

			token := password.MustGenerate(10, 4, 0, false, false)

			options := docker.RunOptions{
				Labels: map[string]string{
					common.KindKey: Jupyter,
				},

				Env: map[string]string{
					"JUPYTER_TOKEN":      token,
					"JUPYTER_ENABLE_LAB": "yes",
					"RESTARTABLE":        "yes",
				},

				Ports: map[int]int{
					port: 8888,
				},

				// Volumes: map[string]string{
				// 	name: "/home/jovyan/work",
				// },
			}

			if err := docker.Run(ctx, image, options); err != nil {
				return err
			}

			cli.Table([]string{"Name", "Value"}, [][]string{
				{"URL", fmt.Sprintf("http://localhost:%d", port)},
				{"Token", token},
			})

			return nil
		},
	}
}
