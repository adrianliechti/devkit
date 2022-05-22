package influxdb

import (
	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/app/common"
	"github.com/adrianliechti/devkit/pkg/cli"
)

const (
	InfluxDB = "influxdb"
)

var Command = &cli.Command{
	Name:  InfluxDB,
	Usage: "local InfluxDB server",

	Category: app.CategoryDatabase,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		common.ListCommand(InfluxDB),

		CreateCommand(),
		common.DeleteCommand(InfluxDB),

		common.LogsCommand(InfluxDB),
		common.ShellCommand(InfluxDB, "/bin/bash"),
	},
}
