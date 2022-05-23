package template

import (
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

	Action: func(c *cli.Context) error {
		options := templateOptions{
			Name: MustName(c, "demo"),
		}

		return runTemplate(c.Context, "", TemplateASPNET, options)
	},
}
