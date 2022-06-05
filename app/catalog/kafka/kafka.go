package kafka

import (
	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/app/catalog"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/container/kafka"
)

const (
	Kafka = "kafka"
)

var Command = &cli.Command{
	Name:  Kafka,
	Usage: "local Kafka broker",

	Category: app.MessagingCategory,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		catalog.ListCommand(Kafka),

		catalog.CreateCommand(Kafka, kafka.New, kafka.Info),
		catalog.DeleteCommand(Kafka),

		catalog.InfoCommand(Kafka, kafka.Info),
		catalog.LogsCommand(Kafka),

		catalog.ShellCommand(Kafka, kafka.DefaultShell),
	},
}
