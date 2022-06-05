package rabbitmq

import (
	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/app/catalog"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/container/rabbitmq"
)

const (
	RabbitMQ = "rabbitmq"
)

var Command = &cli.Command{
	Name:  RabbitMQ,
	Usage: "local RabbitMQ broker",

	Category: app.MessagingCategory,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		catalog.ListCommand(RabbitMQ),

		catalog.CreateCommand(RabbitMQ, rabbitmq.New, rabbitmq.Info),
		catalog.DeleteCommand(RabbitMQ),

		catalog.InfoCommand(RabbitMQ, rabbitmq.Info),
		catalog.LogsCommand(RabbitMQ),

		catalog.ConsoleCommand(RabbitMQ, rabbitmq.Info, rabbitmq.ConsolePort),
		catalog.ShellCommand(RabbitMQ, rabbitmq.DefaultShell),
	},
}
