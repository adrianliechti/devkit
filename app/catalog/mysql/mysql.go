package mysql

import (
	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/app/catalog"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/container/mysql"
)

const (
	MySQL = "mysql"
)

var Command = &cli.Command{
	Name:  MySQL,
	Usage: "local MySQL server",

	Category: app.DatabaseCategory,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		catalog.ListCommand(MySQL),

		catalog.CreateCommand(MySQL, mysql.New),
		catalog.DeleteCommand(MySQL),

		catalog.InfoCommand(MySQL, mysql.Info),
		catalog.LogsCommand(MySQL),

		catalog.ClientCommand(MySQL, mysql.ClientCmd),
		catalog.ShellCommand(MySQL, mysql.DefaultShell),
	},
}
