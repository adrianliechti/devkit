package mariadb

import (
	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/app/common"
	"github.com/adrianliechti/devkit/pkg/cli"
)

const (
	MariaDB = "mariadb"
)

var Command = &cli.Command{
	Name:  MariaDB,
	Usage: "local MariaDB server",

	Category: app.CategoryDatabase,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		common.ListCommand(MariaDB),

		CreateCommand(),
		common.DeleteCommand(MariaDB),

		common.LogsCommand(MariaDB),
		common.ShellCommand(MariaDB, "/bin/bash"),
		ClientCommand(),
	},
}
