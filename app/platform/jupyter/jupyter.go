package jupyter

import (
	"github.com/adrianliechti/devkit/app/common"
	"github.com/adrianliechti/devkit/app/platform"
	"github.com/adrianliechti/devkit/pkg/cli"
)

const (
	Jupyter = "jupyter"
)

var Command = &cli.Command{
	Name:  Jupyter,
	Usage: "local Jupyter notebook",

	Category: platform.Category,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		common.ListCommand(Jupyter),

		CreateCommand(),
		common.DeleteCommand(Jupyter),

		common.LogsCommand(Jupyter),
		common.ShellCommand(Jupyter, "/bin/bash"),
	},
}
