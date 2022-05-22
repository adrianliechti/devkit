package mongodb

import (
	"github.com/adrianliechti/devkit/app/common"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/docker"
)

func ClientCommand() *cli.Command {
	return &cli.Command{
		Name:  "cli",
		Usage: "run mongo in instance",

		Action: func(c *cli.Context) error {
			ctx := c.Context
			container := common.MustContainer(ctx, MongoDB)

			options := docker.ExecOptions{}

			return docker.ExecInteractive(ctx, container, options,
				"/bin/bash", "-c",
				"mongo --quiet --norc --username ${MONGO_INITDB_ROOT_USERNAME} --password ${MONGO_INITDB_ROOT_PASSWORD}",
			)
		},
	}
}
