package mariadb

import (
	"github.com/adrianliechti/devkit/app/common"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/docker"
)

func ClientCommand() *cli.Command {
	return &cli.Command{
		Name:  "cli",
		Usage: "run mysql in instance",

		Action: func(c *cli.Context) error {
			ctx := c.Context
			container := common.MustContainer(ctx, MariaDB)

			options := docker.ExecOptions{}

			return docker.ExecInteractive(ctx, container, options,
				"/bin/bash", "-c",
				"mysql --user=root --password=${MARIADB_ROOT_PASSWORD} ${MARIADB_DATABASE}",
			)
		},
	}
}
