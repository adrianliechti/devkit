package grafana

import (
	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/app/catalog"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/container/grafana"
)

const (
	Grafana = "grafana"
)

var Command = &cli.Command{
	Name:  Grafana,
	Usage: "local Grafana server",

	Category: app.PlatformCategory,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		catalog.ListCommand(Grafana),

		catalog.CreateCommand(Grafana, grafana.New, grafana.Info),
		catalog.DeleteCommand(Grafana),

		catalog.InfoCommand(Grafana, grafana.Info),
		catalog.LogsCommand(Grafana),

		catalog.ShellCommand(Grafana, grafana.DefaultShell),
	},
}
