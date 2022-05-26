package db2

import (
	"github.com/adrianliechti/devkit/app/common"
	"github.com/adrianliechti/devkit/app/database"
	"github.com/adrianliechti/devkit/pkg/cli"
)

const (
	DB2 = "db2"
)

var Command = &cli.Command{
	Name:  DB2,
	Usage: "local DB2 server",

	Category: database.Category,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		common.ListCommand(DB2),

		CreateCommand(),
		common.DeleteCommand(DB2),

		common.LogsCommand(DB2),
		ClientCommand(),
		common.ShellCommand(DB2, "/bin/bash"),
	},
}
