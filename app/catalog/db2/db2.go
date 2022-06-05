package db2

import (
	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/app/catalog"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/container/db2"
)

const (
	DB2 = "db2"
)

var Command = &cli.Command{
	Name:  DB2,
	Usage: "local DB2 server",

	Category: app.DatabaseCategory,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		catalog.ListCommand(DB2),

		catalog.CreateCommand(DB2, db2.New, db2.Info),
		catalog.DeleteCommand(DB2),

		catalog.InfoCommand(DB2, db2.Info),
		catalog.LogsCommand(DB2),

		catalog.ClientCommand(DB2, db2.ClientCmd),
		catalog.ShellCommand(DB2, db2.DefaultShell),
	},
}
