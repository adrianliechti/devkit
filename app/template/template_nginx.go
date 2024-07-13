package template

import (
	"context"

	"github.com/adrianliechti/devkit/pkg/cli"
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
		options := templateOptions{
			Name: MustName(ctx, cmd, "demo"),
		}

		return runTemplate(ctx, "", TemplateNginx, options)
	},
}
