package mysql

import (
	"github.com/adrianliechti/devkit/app/common"
	"github.com/adrianliechti/devkit/app/database"
	"github.com/adrianliechti/devkit/pkg/cli"
)

const (
	MySQL = "mysql"
)

var Command = &cli.Command{
	Name:  MySQL,
	Usage: "local MySQL server",

	Category: database.Category,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		common.ListCommand(MySQL),

		CreateCommand(),
		common.DeleteCommand(MySQL),

		common.LogsCommand(MySQL),
		common.ShellCommand(MySQL, "/bin/bash"),
		ClientCommand(),
	},
}
