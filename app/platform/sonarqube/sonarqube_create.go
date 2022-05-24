package sonarqube

import (
	"fmt"
	"runtime"

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
			image := "sonarqube:9-community"

			if runtime.GOARCH == "arm64" {
				image = "mwizner/sonarqube:9.4.0-community"
			}

			target := 9000
			port := app.MustPortOrRandom(c, target)

			username := "admin"
			password := "admin"

			options := docker.RunOptions{
				Labels: map[string]string{
					common.KindKey: SonarQube,
				},

				MaxNoProcs: 8192,
				MaxNoFiles: 131072,

				Env: map[string]string{
					"SONAR_ES_BOOTSTRAP_CHECKS_DISABLE": "true",
					"SONAR_SEARCH_JAVAADDITIONALOPTS":   "-Dbootstrap.system_call_filter=false",
				},

				Ports: map[int]int{
					port: target,
				},

				// /opt/sonarqube/data
				// /opt/sonarqube/logs
				// /opt/sonarqube/extensions
			}

			if err := docker.Run(ctx, image, options); err != nil {
				return err
			}

			cli.Table([]string{"Name", "Value"}, [][]string{
				{"Host", fmt.Sprintf("localhost:%d", port)},
				{"Username", username},
				{"Password", password},
				{"URL", fmt.Sprintf("http://localhost:%d", port)},
			})

			return nil
		},
	}
}
