package mssql

import (
	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/app/catalog"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/container/mssql"
)

const (
	MSSQL = "mssql"
)

var Command = &cli.Command{
	Name:  MSSQL,
	Usage: "local MSSQL server",

	Category: app.DatabaseCategory,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		catalog.ListCommand(MSSQL),

		catalog.CreateCommand(MSSQL, mssql.New),
		catalog.DeleteCommand(MSSQL),

		catalog.InfoCommand(MSSQL, mssql.Info),
		catalog.LogsCommand(MSSQL),

		catalog.ClientCommand(MSSQL, mssql.ClientCmd),
		catalog.ShellCommand(MSSQL, mssql.DefaultShell),
	},
}
