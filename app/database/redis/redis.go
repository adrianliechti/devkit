package redis

import (
	"github.com/adrianliechti/devkit/app/common"
	"github.com/adrianliechti/devkit/app/database"
	"github.com/adrianliechti/devkit/pkg/cli"
)

const (
	Redis = "redis"
)

var Command = &cli.Command{
	Name:  Redis,
	Usage: "local Redis server",

	Category: database.Category,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		common.ListCommand(Redis),

		CreateCommand(),
		common.DeleteCommand(Redis),

		common.LogsCommand(Redis),
		common.ShellCommand(Redis, "/bin/bash"),
		ClientCommand(),
	},
}
