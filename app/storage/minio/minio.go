package minio

import (
	"github.com/adrianliechti/devkit/app/common"
	"github.com/adrianliechti/devkit/app/storage"
	"github.com/adrianliechti/devkit/pkg/cli"
)

const (
	MinIO = "minio"
)

var Command = &cli.Command{
	Name:  MinIO,
	Usage: "local MinIO server",

	Category: storage.Category,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		common.ListCommand(MinIO),

		CreateCommand(),
		common.DeleteCommand(MinIO),

		common.LogsCommand(MinIO),
		common.ShellCommand(MinIO, "/bin/ash"),
		ConsoleCommand(),
	},
}
