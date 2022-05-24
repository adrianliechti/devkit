package elasticsearch

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
			image := "docker.elastic.co/elasticsearch/elasticsearch:8.2.0"

			target := 9200
			port := app.MustPortOrRandom(c, target)

			username := "elastic"
			password := password.MustGenerate(10, 4, 0, false, false)

			options := docker.RunOptions{
				Labels: map[string]string{
					common.KindKey: Elasticsearch,
				},

				Env: map[string]string{
					"node.name": "es",

					"cluster.name":   "default",
					"discovery.type": "single-node",

					"xpack.security.enabled": "true",

					"ELASTIC_PASSWORD": password,
				},

				Ports: map[int]int{
					port: target,
				},

				// Volumes: map[string]string{
				// 	name: "/usr/share/elasticsearch/data",
				// },
			}

			if err := docker.Run(ctx, image, options); err != nil {
				return err
			}

			cli.Table([]string{"Name", "Value"}, [][]string{
				{"Host", fmt.Sprintf("localhost:%d", port)},
				{"Username", username},
				{"Password", password},
				{"URL", fmt.Sprintf("http://%s:%s@localhost:%d", username, password, port)},
			})

			return nil
		},
	}
}
