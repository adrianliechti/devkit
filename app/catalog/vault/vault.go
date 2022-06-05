package vault

import (
	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/app/catalog"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/container/vault"
)

const (
	Vault = "vault"
)

var Command = &cli.Command{
	Name:  Vault,
	Usage: "local Vault server",

	Category: app.StorageCategory,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		catalog.ListCommand(Vault),

		catalog.CreateCommand(Vault, vault.New),
		catalog.DeleteCommand(Vault),

		catalog.InfoCommand(Vault, vault.Info),
		catalog.LogsCommand(Vault),

		catalog.ShellCommand(Vault, vault.DefaultShell),
	},
}
