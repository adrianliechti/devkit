package vault

import (
	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/app/common"
	"github.com/adrianliechti/devkit/pkg/cli"
)

const (
	Vault = "vault"
)

var Command = &cli.Command{
	Name:  Vault,
	Usage: "local Vault server",

	Category: app.CategoryStorage,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		common.ListCommand(Vault),

		createCommand(),
		common.DeleteCommand(Vault),

		common.LogsCommand(Vault),
		common.ShellCommand(Vault, "/bin/ash"),
	},
}
