package nats

import (
	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/app/catalog"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/container/nats"
)

const (
	NATS = "nats"
)

var Command = &cli.Command{
	Name:  NATS,
	Usage: "local NATS server",

	Category: app.MessagingCategory,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		catalog.ListCommand(NATS),

		catalog.CreateCommand(NATS, nats.New),
		catalog.DeleteCommand(NATS),

		catalog.InfoCommand(NATS, nats.Info),
		catalog.LogsCommand(NATS),
	},
}
