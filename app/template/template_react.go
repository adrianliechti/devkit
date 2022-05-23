package template

import (
	"github.com/adrianliechti/devkit/pkg/cli"
)

var reactCommand = &cli.Command{
	Name:  "react",
	Usage: "create React web app",

	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "name",
			Usage:    "package name",
			Required: true,
		},
	},

	Action: func(c *cli.Context) error {
		options := templateOptions{
			Name: c.String("name"),
		}

		return runTemplate(c.Context, "", TemplateReact, options)
	},
}
