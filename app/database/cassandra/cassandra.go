package cassandra

import (
	"github.com/adrianliechti/devkit/app/common"
	"github.com/adrianliechti/devkit/app/database"
	"github.com/adrianliechti/devkit/pkg/cli"
)

const (
	Cassandra = "cassandra"
)

var Command = &cli.Command{
	Name:  Cassandra,
	Usage: "local Cassandra server",

	Category: database.Category,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		common.ListCommand(Cassandra),

		CreateCommand(),
		common.DeleteCommand(Cassandra),

		common.LogsCommand(Cassandra),
		ClientCommand(),
		common.ShellCommand(Cassandra, "/bin/bash"),
	},
}
