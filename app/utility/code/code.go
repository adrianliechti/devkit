package code

import (
	"github.com/adrianliechti/devkit/app/common"
	"github.com/adrianliechti/devkit/app/utility"
	"github.com/adrianliechti/devkit/pkg/cli"
)

const (
	Code = "code"
)

var Command = &cli.Command{
	Name:  Code,
	Usage: "local Code server",

	Category: utility.Category,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		common.ListCommand(Code),

		CreateCommand(),
		common.DeleteCommand(Code),

		common.LogsCommand(Code),
		common.ShellCommand(Code, "/bin/bash"),
	},
}
