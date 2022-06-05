package mailtrap

import (
	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/app/catalog"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/container/mailtrap"
)

const (
	MailTrap = "mailtrap"
)

var Command = &cli.Command{
	Name:  MailTrap,
	Usage: "local MailTrap server",

	Category: app.PlatformCategory,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		catalog.ListCommand(MailTrap),

		catalog.CreateCommand(MailTrap, mailtrap.New, mailtrap.Info),
		catalog.DeleteCommand(MailTrap),

		catalog.InfoCommand(MailTrap, mailtrap.Info),
		catalog.LogsCommand(MailTrap),

		catalog.ConsoleCommand(MailTrap, mailtrap.Info, mailtrap.ConsolePort),
		catalog.ShellCommand(MailTrap, mailtrap.DefaultShell),
	},
}
