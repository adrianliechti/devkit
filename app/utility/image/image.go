package image

import (
	"context"
	"errors"

	"github.com/adrianliechti/devkit/app/utility"
	"github.com/adrianliechti/go-cli"
)

var Command = &cli.Command{
	Name:  "image",
	Usage: "docker/oci image tools",

	HideHelpCommand: true,

	Category: utility.Category,

	Commands: []*cli.Command{
		lintCommand,
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

func Image(ctx context.Context, cmd *cli.Command) string {
	image := cmd.String(ImageFlag.Name)
	return image
}

func MustImage(ctx context.Context, cmd *cli.Command) string {
	image := Image(ctx, cmd)

	if len(image) == 0 {
		cli.Fatal(errors.New("image missing"))
	}

	return image
}
