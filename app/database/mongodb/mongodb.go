package mongodb

import (
	"github.com/adrianliechti/devkit/app/common"
	"github.com/adrianliechti/devkit/app/database"
	"github.com/adrianliechti/devkit/pkg/cli"
)

const (
	MongoDB = "mongodb"
)

var Command = &cli.Command{
	Name:  MongoDB,
	Usage: "local MongoDB server",

	Category: database.Category,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		common.ListCommand(MongoDB),

		CreateCommand(),
		common.DeleteCommand(MongoDB),

		common.LogsCommand(MongoDB),
		common.ShellCommand(MongoDB, "/bin/bash"),
		ClientCommand(),
	},
}
