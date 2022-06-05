package etcd

import (
	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/app/catalog"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/container/etcd"
)

const (
	ETCD = "etcd"
)

var Command = &cli.Command{
	Name:  ETCD,
	Usage: "local etcd server",

	Category: app.DatabaseCategory,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		catalog.ListCommand(ETCD),

		catalog.CreateCommand(ETCD, etcd.New, etcd.Info),
		catalog.DeleteCommand(ETCD),

		catalog.InfoCommand(ETCD, etcd.Info),
		catalog.LogsCommand(ETCD),

		catalog.ShellCommand(ETCD, etcd.DefaultShell),
	},
}
