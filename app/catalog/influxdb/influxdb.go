package influxdb

import (
	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/app/catalog"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/container/influxdb"
)

const (
	InfluxDB = "influxdb"
)

var Command = &cli.Command{
	Name:  InfluxDB,
	Usage: "local InfluxDB server",

	Category: app.DatabaseCategory,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		catalog.ListCommand(InfluxDB),

		catalog.CreateCommand(InfluxDB, influxdb.New, influxdb.Info),
		catalog.DeleteCommand(InfluxDB),

		catalog.InfoCommand(InfluxDB, influxdb.Info),
		catalog.LogsCommand(InfluxDB),

		catalog.ConsoleCommand(InfluxDB, influxdb.Info, influxdb.ConsolePort),
		catalog.ShellCommand(InfluxDB, influxdb.DefaultShell),
	},
}
