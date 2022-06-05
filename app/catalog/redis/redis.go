package redis

import (
	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/app/catalog"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/container/redis"
)

const (
	Redis = "redis"
)

var Command = &cli.Command{
	Name:  Redis,
	Usage: "local Redis server",

	Category: app.DatabaseCategory,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		catalog.ListCommand(Redis),

		catalog.CreateCommand(Redis, redis.New),
		catalog.DeleteCommand(Redis),

		catalog.InfoCommand(Redis, redis.Info),
		catalog.LogsCommand(Redis),

		catalog.ClientCommand(Redis, redis.ClientCmd),
		catalog.ShellCommand(Redis, redis.DefaultShell),
	},
}
