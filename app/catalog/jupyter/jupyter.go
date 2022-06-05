package jupyter

import (
	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/app/catalog"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/container/jupyter"
)

const (
	Jupyter = "jupyter"
)

var Command = &cli.Command{
	Name:  Jupyter,
	Usage: "local Jupyter notebook",

	Category: app.PlatformCategory,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		catalog.ListCommand(Jupyter),

		catalog.CreateCommand(Jupyter, jupyter.New, jupyter.Info),
		catalog.DeleteCommand(Jupyter),

		catalog.InfoCommand(Jupyter, jupyter.Info),
		catalog.LogsCommand(Jupyter),

		catalog.ShellCommand(Jupyter, jupyter.DefaultShell),
	},
}
