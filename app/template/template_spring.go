package template

import (
	"github.com/adrianliechti/devkit/pkg/cli"
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

	Action: func(c *cli.Context) error {
		options := templateOptions{
			Group: MustGroup(c, "org.example"),
			Name:  MustName(c, "demo"),
		}

		return runTemplate(c.Context, "", TemplateSpring, options)
	},
}
