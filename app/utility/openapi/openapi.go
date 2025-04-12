package openapi

import (
	"github.com/adrianliechti/devkit/app/utility"
	"github.com/adrianliechti/go-cli"
)

var Command = &cli.Command{
	Name:  "openapi",
	Usage: "swagger/openapi tools",

	Category: utility.Category,

	HideHelpCommand: true,

	Commands: []*cli.Command{
		generateCommand,
		mergeCommand,
		mockCommand,
	},
}
