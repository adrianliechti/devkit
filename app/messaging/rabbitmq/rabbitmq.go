package rabbitmq

import (
	"github.com/adrianliechti/devkit/app/common"
	"github.com/adrianliechti/devkit/app/messaging"
	"github.com/adrianliechti/devkit/pkg/cli"
)

const (
	RabbitMQ = "rabbitmq"
)

var Command = &cli.Command{
	Name:  RabbitMQ,
	Usage: "local RabbitMQ broker",

	Category: messaging.Category,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		common.ListCommand(RabbitMQ),

		CreateCommand(),
		common.DeleteCommand(RabbitMQ),

		common.LogsCommand(RabbitMQ),
		common.ShellCommand(RabbitMQ, "/bin/bash"),
		ConsoleCommand(),
	},
}
