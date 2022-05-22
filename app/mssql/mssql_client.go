package mssql

import (
	"github.com/adrianliechti/devkit/app/common"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/docker"
)

func ClientCommand() *cli.Command {
	return &cli.Command{
		Name:  "cli",
		Usage: "run sqlcmd in instance",

		Action: func(c *cli.Context) error {
			ctx := c.Context
			container := common.MustContainer(ctx, MSSQL)

			options := docker.ExecOptions{}

			return docker.ExecInteractive(ctx, container, options,
				"/bin/bash", "-c",
				"/opt/mssql-tools/bin/sqlcmd -S localhost -U sa -P ${SA_PASSWORD}",
			)
		},
	}
}
