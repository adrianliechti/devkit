package template

import (
	"github.com/adrianliechti/devkit/pkg/cli"
)

var golangCommand = &cli.Command{
	Name:  "golang",
	Usage: "create Go web app",

	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "name",
			Usage: "module name",
		},
	},

	Action: func(c *cli.Context) error {
		options := templateOptions{
			Name: MustName(c, "demo"),
		}

		return runTemplate(c.Context, "", TemplateGolang, options)
	},
}
