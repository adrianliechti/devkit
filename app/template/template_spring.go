package template

import (
	"github.com/adrianliechti/devkit/pkg/cli"
)

var springCommand = &cli.Command{
	Name:  "spring",
	Usage: "create Java Spring web app",

	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "group",
			Usage:    "application group",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "name",
			Usage:    "application name",
			Required: true,
		},
	},

	Action: func(c *cli.Context) error {
		options := templateOptions{
			Group: c.String("group"),
			Name:  c.String("name"),
		}

		return runTemplate(c.Context, "", TemplateSpring, options)
	},
}
