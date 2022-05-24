package template

import (
	"github.com/adrianliechti/devkit/pkg/cli"
)

var pythonCommand = &cli.Command{
	Name:  "python",
	Usage: "create Python web app",

	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "name",
			Usage: "app name",
		},
	},

	Action: func(c *cli.Context) error {
		options := templateOptions{
			Name: MustName(c, "demo"),
		}

		return runTemplate(c.Context, "", TemplatePython, options)
	},
}
