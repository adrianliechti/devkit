package cassandra

import (
	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/app/catalog"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/container/cassandra"
)

const (
	Cassandra = "cassandra"
)

var Command = &cli.Command{
	Name:  Cassandra,
	Usage: "local Cassandra server",

	Category: app.DatabaseCategory,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		catalog.ListCommand(Cassandra),

		catalog.CreateCommand(Cassandra, cassandra.New, cassandra.Info),
		catalog.DeleteCommand(Cassandra),

		catalog.InfoCommand(Cassandra, cassandra.Info),
		catalog.LogsCommand(Cassandra),

		catalog.ClientCommand(Cassandra, cassandra.ClientCmd),
		catalog.ShellCommand(Cassandra, cassandra.DefaultShell),
	},
}
