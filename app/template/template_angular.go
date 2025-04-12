package template

import (
	"context"

	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/go-cli"
)

var angularCommand = &cli.Command{
	Name:  "angular",
	Usage: "create Angular app",

	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "name",
			Usage: "package name",
		},
	},

	Action: func(ctx context.Context, cmd *cli.Command) error {
		client := app.MustClient(ctx, cmd)

		options := templateOptions{
			Name: MustName(ctx, cmd, "demo"),
		}

		return runTemplate(ctx, client, "", TemplateAngular, options)
	},
}
