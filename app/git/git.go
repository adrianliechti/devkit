package git

import (
	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/pkg/cli"
)

var Command = &cli.Command{
	Name:  "git",
	Usage: "git repository tools",

	Category: app.CategoryUtilities,

	HideHelpCommand: true,

	Subcommands: []*cli.Command{
		leaksCommand,
		blobsCommand,
		purgeCommand,
	},
}
