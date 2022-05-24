package mssql

import (
	"github.com/adrianliechti/devkit/app/common"
	"github.com/adrianliechti/devkit/app/database"
	"github.com/adrianliechti/devkit/pkg/cli"
)

const (
	MSSQL = "mssql"
)

var Command = &cli.Command{
	Name:  MSSQL,
	Usage: "local MSSQL server",

	Category: database.Category,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		common.ListCommand(MSSQL),

		CreateCommand(),
		common.DeleteCommand(MSSQL),

		common.LogsCommand(MSSQL),
		common.ShellCommand(MSSQL, "/bin/bash"),
		ClientCommand(),
	},
}
