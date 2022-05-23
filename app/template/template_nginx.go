package template

import (
	"github.com/adrianliechti/devkit/pkg/cli"
)

var nginxCommand = &cli.Command{
	Name:  "nginx",
	Usage: "create Nginx web app",

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

		return runTemplate(c.Context, "", TemplateNginx, options)
	},
}
