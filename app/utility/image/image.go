package image

import (
	"errors"

	"github.com/adrianliechti/devkit/app/utility"
	"github.com/adrianliechti/devkit/pkg/cli"
)

var Command = &cli.Command{
	Name:  "image",
	Usage: "docker/oci image tools",

	HideHelpCommand: true,

	Category: utility.Category,

	Subcommands: []*cli.Command{
		lintCommand,
		packCommand,

		scanCommand,
		inspectCommand,

		bomCommand,

		browseCommand,
	},
}

var ImageFlag = &cli.StringFlag{
	Name:     "image",
	Usage:    "docker image",
	Required: true,
}

func Image(c *cli.Context) string {
	image := c.String(ImageFlag.Name)
	return image
}

func MustImage(c *cli.Context) string {
	image := Image(c)

	if len(image) == 0 {
		cli.Fatal(errors.New("image missing"))
	}

	return image
}
