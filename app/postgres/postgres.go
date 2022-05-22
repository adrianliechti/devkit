package postgres

import (
	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/app/common"
	"github.com/adrianliechti/devkit/pkg/cli"
)

const (
	PostgreSQL = "postgres"
)

var Command = &cli.Command{
	Name:  PostgreSQL,
	Usage: "local PostgreSQL server",

	Category: app.CategoryDatabase,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		common.ListCommand(PostgreSQL),

		CreateCommand(),
		common.DeleteCommand(PostgreSQL),

		common.LogsCommand(PostgreSQL),
		common.ShellCommand(PostgreSQL, "/bin/bash"),
		ClientCommand(),
	},
}
