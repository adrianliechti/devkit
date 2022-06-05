package mariadb

import (
	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/app/catalog"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/container/mariadb"
)

const (
	MariaDB = "mariadb"
)

var Command = &cli.Command{
	Name:  MariaDB,
	Usage: "local MariaDB server",

	Category: app.DatabaseCategory,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		catalog.ListCommand(MariaDB),

		catalog.CreateCommand(MariaDB, mariadb.New, mariadb.Info),
		catalog.DeleteCommand(MariaDB),

		catalog.InfoCommand(MariaDB, mariadb.Info),
		catalog.LogsCommand(MariaDB),

		catalog.ClientCommand(MariaDB, mariadb.ClientCmd),
		catalog.ShellCommand(MariaDB, mariadb.DefaultShell),
	},
}
