package elasticsearch

import (
	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/app/common"
	"github.com/adrianliechti/devkit/pkg/cli"
)

const (
	Elasticsearch = "elasticsearch"
)

var Command = &cli.Command{
	Name:  Elasticsearch,
	Usage: "local Elasticsearch server",

	Category: app.CategoryDatabase,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		common.ListCommand(Elasticsearch),

		CreateCommand(),
		common.DeleteCommand(Elasticsearch),

		common.LogsCommand(Elasticsearch),
		common.ShellCommand(Elasticsearch, "/bin/bash"),
	},
}
