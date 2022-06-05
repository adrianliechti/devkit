package minio

import (
	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/app/catalog"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/container/minio"
)

const (
	MinIO = "minio"
)

var Command = &cli.Command{
	Name:  MinIO,
	Usage: "local MinIO server",

	Category: app.StorageCategory,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		catalog.ListCommand(MinIO),

		catalog.CreateCommand(MinIO, minio.New),
		catalog.DeleteCommand(MinIO),

		catalog.InfoCommand(MinIO, minio.Info),
		catalog.LogsCommand(MinIO),

		catalog.ConsoleCommand(MinIO, minio.ConsolePort),
		catalog.ShellCommand(MinIO, minio.DefaultShell),
	},
}
