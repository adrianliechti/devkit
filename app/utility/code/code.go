package code

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/app/utility"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/docker"
)

var Command = &cli.Command{
	Name:  "code",
	Usage: "start temporary Code IDE",

	Category: utility.Category,

	Flags: []cli.Flag{
		app.PortFlag,
	},

	Action: func(c *cli.Context) error {
		port := app.MustPortOrRandom(c, 3000)
		return startCode(c.Context, port)
	},
}

func startCode(ctx context.Context, port int) error {
	image := "adrianliechti/loop-code"

	path, err := os.Getwd()

	if err != nil {
		return err
	}

	target := 3000

	if port == 0 {
		port = target
	}

	cli.Table([]string{"Name", "Value"}, [][]string{
		{"Host", fmt.Sprintf("localhost:%d", port)},
		{"URL", fmt.Sprintf("http://localhost:%d", port)},
	})

	cli.Info()
	cli.Info("Forward " + path + " to " + "/workspace")
	cli.Info()

	options := docker.RunOptions{
		Platform: "linux/amd64",

		Ports: map[int]int{
			port: target,
		},

		Volumes: map[string]string{
			path: "/workspace",
		},

		Stdout: io.Discard,
	}

	return docker.RunInteractive(ctx, image, options)
}
