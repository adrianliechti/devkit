package postgres

import (
	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/app/catalog"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/container/postgres"
)

const (
	PostgreSQL = "postgres"
)

var Command = &cli.Command{
	Name:  PostgreSQL,
	Usage: "local PostgreSQL server",

	Category: app.DatabaseCategory,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		catalog.ListCommand(PostgreSQL),

		catalog.CreateCommand(PostgreSQL, postgres.New),
		catalog.DeleteCommand(PostgreSQL),

		catalog.InfoCommand(PostgreSQL, postgres.Info),
		catalog.LogsCommand(PostgreSQL),

		catalog.ClientCommand(PostgreSQL, postgres.ClientCmd),
		catalog.ShellCommand(PostgreSQL, postgres.DefaultShell),
	},
}
