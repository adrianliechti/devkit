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
	"github.com/adrianliechti/devkit/pkg/to"
)

var Command = &cli.Command{
	Name:  "code",
	Usage: "start temporary Code IDE",

	Category: utility.Category,

	Flags: []cli.Flag{
		app.PortFlag(""),
	},

	Action: func(c *cli.Context) error {
		port := app.MustPortOrRandom(c, "", 3000)
		return startCode(c.Context, port)
	},
}

func startCode(ctx context.Context, port int) error {
	image := "adrianliechti/loop-code"

	path, err := os.Getwd()

	if err != nil {
		return err
	}

	docker.Pull(ctx, image, docker.PullOptions{
		Platform: "linux/amd64",
	})

	cli.Table([]string{"Name", "Value"}, [][]string{
		{"URL", fmt.Sprintf("http://localhost:%d", port)},
	})

	cli.Info()
	cli.Info("Forward " + path + " to " + "/workspace")
	cli.Info()

	options := docker.RunOptions{
		Platform: "linux/amd64",

		Ports: []docker.ContainerPort{
			{
				Port:     3000,
				Protocol: docker.ProtocolTCP,

				HostPort: to.IntPtr(port),
			},
		},

		Volumes: []docker.ContainerMount{
			{
				Path:     "/workspace",
				HostPath: path,
			},
		},

		Stdout: io.Discard,
	}

	return docker.RunInteractive(ctx, image, options)
}
