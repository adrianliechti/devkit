package template

import (
	"context"

	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/pkg/cli"
)

var aspnetCommand = &cli.Command{
	Name:  "aspnet",
	Usage: "create ASP.NET Core app",

	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "name",
			Usage: "application name",
		},
	},

	Action: func(ctx context.Context, cmd *cli.Command) error {
		client := app.MustClient(ctx, cmd)

		options := templateOptions{
			Name: MustName(ctx, cmd, "demo"),
		}

		return runTemplate(ctx, client, "", TemplateASPNET, options)
	},
}
