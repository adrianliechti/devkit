package jenkins

import (
	"github.com/adrianliechti/devkit/app/common"
	"github.com/adrianliechti/devkit/app/messaging"
	"github.com/adrianliechti/devkit/pkg/cli"
)

const (
	Jenkins = "jenkins"
)

var Command = &cli.Command{
	Name:  Jenkins,
	Usage: "local Jenkins broker",

	Category: messaging.Category,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		common.ListCommand(Jenkins),

		CreateCommand(),
		common.DeleteCommand(Jenkins),

		common.LogsCommand(Jenkins),
		common.ShellCommand(Jenkins, "/bin/bash"),
	},
}
