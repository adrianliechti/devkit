package influxdb

import (
	"github.com/adrianliechti/devkit/app/common"
	"github.com/adrianliechti/devkit/app/database"
	"github.com/adrianliechti/devkit/pkg/cli"
)

const (
	InfluxDB = "influxdb"
)

var Command = &cli.Command{
	Name:  InfluxDB,
	Usage: "local InfluxDB server",

	Category: database.Category,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		common.ListCommand(InfluxDB),

		CreateCommand(),
		common.DeleteCommand(InfluxDB),

		common.LogsCommand(InfluxDB),
		common.ShellCommand(InfluxDB, "/bin/bash"),
	},
}
