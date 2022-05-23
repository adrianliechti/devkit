package template

import (
	"github.com/adrianliechti/devkit/pkg/cli"
)

var aspnetCommand = &cli.Command{
	Name:  "aspnet",
	Usage: "create ASP.NET Core app",

	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "name",
			Usage:    "application name",
			Required: true,
		},
	},

	Action: func(c *cli.Context) error {
		options := templateOptions{
			Name: c.String("name"),
		}

		return runTemplate(c.Context, "", TemplateASPNET, options)
	},
}
