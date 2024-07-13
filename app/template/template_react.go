package template

import (
	"context"

	"github.com/adrianliechti/devkit/pkg/cli"
)

var reactCommand = &cli.Command{
	Name:  "react",
	Usage: "create React web app",

	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "name",
			Usage: "package name",
		},
	},

	Action: func(ctx context.Context, cmd *cli.Command) error {
		options := templateOptions{
			Name: MustName(ctx, cmd, "demo"),
		}

		return runTemplate(ctx, "", TemplateReact, options)
	},
}
