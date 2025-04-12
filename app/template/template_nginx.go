package template

import (
	"context"

	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/go-cli"
)

var nginxCommand = &cli.Command{
	Name:  "nginx",
	Usage: "create Nginx web app",

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

		return runTemplate(ctx, client, "", TemplateNginx, options)
	},
}
