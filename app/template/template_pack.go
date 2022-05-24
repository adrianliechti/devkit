package template

import (
	"github.com/adrianliechti/devkit/pkg/cli"
)

var packCommand = &cli.Command{
	Name:  "pack",
	Usage: "create app using buildpacks",

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

		return runTemplate(c.Context, "", TemplatePack, options)
	},
}
