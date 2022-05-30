package mailtrap

import (
	"github.com/adrianliechti/devkit/app/common"
	"github.com/adrianliechti/devkit/app/platform"
	"github.com/adrianliechti/devkit/pkg/cli"
)

const (
	MailTrap = "mailtrap"
)

var Command = &cli.Command{
	Name:  MailTrap,
	Usage: "local MailTrap server",

	Category: platform.Category,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		common.ListCommand(MailTrap),

		CreateCommand(),
		common.DeleteCommand(MailTrap),

		common.LogsCommand(MailTrap),
		common.ShellCommand(MailTrap, "/bin/bash"),
		ConsoleCommand(),
	},
}
