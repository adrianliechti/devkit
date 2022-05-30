package grafana

import (
	"github.com/adrianliechti/devkit/app/common"
	"github.com/adrianliechti/devkit/app/platform"
	"github.com/adrianliechti/devkit/pkg/cli"
)

const (
	Grafana = "grafana"
)

var Command = &cli.Command{
	Name:  Grafana,
	Usage: "local Grafana server",

	Category: platform.Category,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		common.ListCommand(Grafana),

		CreateCommand(),
		common.DeleteCommand(Grafana),

		common.LogsCommand(Grafana),
		common.ShellCommand(Grafana, "/bin/bash"),
	},
}
