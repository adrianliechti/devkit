package template

import (
	"context"

	"github.com/adrianliechti/devkit/pkg/cli"
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
		options := templateOptions{
			Name: MustName(ctx, cmd, "demo"),
		}

		return runTemplate(ctx, "", TemplateAngular, options)
	},
}
