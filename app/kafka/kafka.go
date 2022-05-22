package kafka

import (
	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/app/common"
	"github.com/adrianliechti/devkit/pkg/cli"
)

const (
	Kafka = "kafka"
)

var Command = &cli.Command{
	Name:  Kafka,
	Usage: "local Kafka broker",

	Category: app.CategoryMessaging,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		common.ListCommand(Kafka),

		CreateCommand(),
		common.DeleteCommand(Kafka),

		common.LogsCommand(Kafka),
		common.ShellCommand(Kafka, "/bin/bash"),
	},
}
