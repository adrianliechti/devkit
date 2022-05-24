package template

import (
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

	Action: func(c *cli.Context) error {
		options := templateOptions{
			Name: MustName(c, "demo"),
		}

		return runTemplate(c.Context, "", TemplateAngular, options)
	},
}
