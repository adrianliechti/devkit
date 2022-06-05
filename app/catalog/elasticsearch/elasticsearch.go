package elasticsearch

import (
	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/app/catalog"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/container/elasticsearch"
)

const (
	Elasticsearch = "elasticsearch"
)

var Command = &cli.Command{
	Name:  Elasticsearch,
	Usage: "local Elasticsearch server",

	Category: app.DatabaseCategory,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		catalog.ListCommand(Elasticsearch),

		catalog.CreateCommand(Elasticsearch, elasticsearch.New),
		catalog.DeleteCommand(Elasticsearch),

		catalog.InfoCommand(Elasticsearch, elasticsearch.Info),
		catalog.LogsCommand(Elasticsearch),

		catalog.ShellCommand(Elasticsearch, elasticsearch.DefaultShell),
	},
}
