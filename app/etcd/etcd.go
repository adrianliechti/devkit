package etcd

import (
	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/app/common"
	"github.com/adrianliechti/devkit/pkg/cli"
)

const (
	ETCD = "etcd"
)

var Command = &cli.Command{
	Name:  ETCD,
	Usage: "local etcd server",

	Category: app.CategoryDatabase,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		common.ListCommand(ETCD),

		CreateCommand(),
		common.DeleteCommand(ETCD),

		common.LogsCommand(ETCD),
		common.ShellCommand(ETCD, "/bin/ash"),
	},
}
