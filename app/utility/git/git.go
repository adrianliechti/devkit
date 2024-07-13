package git

import (
	"github.com/adrianliechti/devkit/app/utility"
	"github.com/adrianliechti/devkit/pkg/cli"
)

var Command = &cli.Command{
	Name:  "git",
	Usage: "git repository tools",

	Category: utility.Category,

	HideHelpCommand: true,

	Commands: []*cli.Command{
		blobsCommand,
		leaksCommand,
		purgeCommand,
	},
}
