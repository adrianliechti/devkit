package mongodb

import (
	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/app/catalog"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/container/mongodb"
)

const (
	MongoDB = "mongodb"
)

var Command = &cli.Command{
	Name:  MongoDB,
	Usage: "local MongoDB server",

	Category: app.DatabaseCategory,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		catalog.ListCommand(MongoDB),

		catalog.CreateCommand(MongoDB, mongodb.New, mongodb.Info),
		catalog.DeleteCommand(MongoDB),

		catalog.InfoCommand(MongoDB, mongodb.Info),
		catalog.LogsCommand(MongoDB),

		catalog.ClientCommand(MongoDB, mongodb.ClientCmd),
		catalog.ShellCommand(MongoDB, mongodb.DefaultShell),
	},
}
