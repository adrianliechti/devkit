package immudb

import (
	"github.com/adrianliechti/devkit/app/common"
	"github.com/adrianliechti/devkit/app/database"
	"github.com/adrianliechti/devkit/pkg/cli"
)

const (
	ImmuDB = "immudb"
)

var Command = &cli.Command{
	Name:  ImmuDB,
	Usage: "local ImmuDB server",

	Category: database.Category,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		common.ListCommand(ImmuDB),

		CreateCommand(),
		common.DeleteCommand(ImmuDB),

		common.LogsCommand(ImmuDB),
	},
}
