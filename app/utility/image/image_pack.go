package image

import (
	"context"
	"os"

	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/docker"
)

var packCommand = &cli.Command{
	Name:  "pack",
	Usage: "create image using buildpacks",

	Flags: []cli.Flag{
		ImageFlag,
		&cli.StringFlag{
			Name:  "builder",
			Usage: "builder image",

			Value: "gcr.io/buildpacks/builder",
		},
	},

	Action: func(c *cli.Context) error {
		image := MustImage(c)
		builder := c.String("builder")

		return runPack(c.Context, image, builder)
	},
}

func runPack(ctx context.Context, image, builder string) error {
	wd, err := os.Getwd()

	if err != nil {
		return err
	}

	args := []string{
		"build",
		image,
		"--path", "/src",
		"--builder", builder,
	}

	options := docker.RunOptions{
		User: "0:0",
		Volumes: []docker.ContainerMount{
			{
				Path:     "/src",
				HostPath: wd,
			},
			{
				Path:     "/var/run/docker.sock",
				HostPath: "/var/run/docker.sock",
			},
		},
	}

	return docker.RunInteractive(ctx, "buildpacksio/pack", options, args...)
}
