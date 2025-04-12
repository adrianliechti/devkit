package template

import (
	"context"

	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/go-cli"
)

var springCommand = &cli.Command{
	Name:  "spring",
	Usage: "create Java Spring web app",

	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "group",
			Usage: "application group",
		},
		&cli.StringFlag{
			Name:  "name",
			Usage: "application name",
		},
	},

	Action: func(ctx context.Context, cmd *cli.Command) error {
		client := app.MustClient(ctx, cmd)

		options := templateOptions{
			Group: MustGroup(ctx, cmd, "org.example"),
			Name:  MustName(ctx, cmd, "demo"),
		}

		return runTemplate(ctx, client, "", TemplateSpring, options)
	},
}
