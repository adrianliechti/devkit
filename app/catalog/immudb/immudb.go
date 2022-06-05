package immudb

import (
	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/app/catalog"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/container/immudb"
)

const (
	ImmuDB = "immudb"
)

var Command = &cli.Command{
	Name:  ImmuDB,
	Usage: "local ImmuDB server",

	Category: app.DatabaseCategory,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		catalog.ListCommand(ImmuDB),

		catalog.CreateCommand(ImmuDB, immudb.New),
		catalog.DeleteCommand(ImmuDB),

		catalog.InfoCommand(ImmuDB, immudb.Info),
		catalog.LogsCommand(ImmuDB),

		catalog.ConsoleCommand(ImmuDB, immudb.ConsolePort),
		catalog.ShellCommand(ImmuDB, immudb.DefaultShell),
	},
}
