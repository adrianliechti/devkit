package nats

import (
	"github.com/adrianliechti/devkit/app/common"
	"github.com/adrianliechti/devkit/app/messaging"
	"github.com/adrianliechti/devkit/pkg/cli"
)

const (
	NATS = "nats"
)

var Command = &cli.Command{
	Name:  NATS,
	Usage: "local NATS server",

	Category: messaging.Category,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		common.ListCommand(NATS),

		CreateCommand(),
		common.DeleteCommand(NATS),

		common.LogsCommand(NATS),
	},
}
