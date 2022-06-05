package jenkins

import (
	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/app/catalog"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/container/jenkins"
)

const (
	Jenkins = "jenkins"
)

var Command = &cli.Command{
	Name:  Jenkins,
	Usage: "local Jenkins server",

	Category: app.PlatformCategory,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		catalog.ListCommand(Jenkins),

		catalog.CreateCommand(Jenkins, jenkins.New),
		catalog.DeleteCommand(Jenkins),

		catalog.InfoCommand(Jenkins, jenkins.Info),
		catalog.LogsCommand(Jenkins),

		catalog.ShellCommand(Jenkins, jenkins.DefaultShell),
	},
}
