package template

import (
	"context"

	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/go-cli"
)

var pythonCommand = &cli.Command{
	Name:  "python",
	Usage: "create Python web app",

	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "name",
			Usage: "app name",
		},
	},

	Action: func(ctx context.Context, cmd *cli.Command) error {
		client := app.MustClient(ctx, cmd)

		options := templateOptions{
			Name: MustName(ctx, cmd, "demo"),
		}

		return runTemplate(ctx, client, "", TemplatePython, options)
	},
}
