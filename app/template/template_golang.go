package template

import (
	"context"

	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/pkg/cli"
)

var golangCommand = &cli.Command{
	Name:  "golang",
	Usage: "create Go web app",

	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "name",
			Usage: "module name",
		},
	},

	Action: func(ctx context.Context, cmd *cli.Command) error {
		client := app.MustClient(ctx, cmd)

		options := templateOptions{
			Name: MustName(ctx, cmd, "demo"),
		}

		return runTemplate(ctx, client, "", TemplateGolang, options)
	},
}
