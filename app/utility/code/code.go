package code

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/app/utility"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/docker"
	"github.com/adrianliechti/devkit/pkg/engine"
)

var Command = &cli.Command{
	Name:  "code",
	Usage: "start temporary Code IDE",

	Category: utility.Category,

	Flags: []cli.Flag{
		app.PortFlag(""),
	},

	Action: func(c *cli.Context) error {
		client := app.MustClient(c)

		port := app.MustPortOrRandom(c, "", 3000)

		return startCode(c.Context, client, port)
	},
}

func startCode(ctx context.Context, client engine.Client, port int) error {
	image := "adrianliechti/loop-code"

	path, err := os.Getwd()

	if err != nil {
		return err
	}

	client.Pull(ctx, image, engine.PullOptions{
		Platform: "linux/amd64",
	})

	time.AfterFunc(2*time.Second, func() {
		cli.OpenURL(fmt.Sprintf("http://localhost:%d", port))
	})

	cli.Table([]string{"Name", "Value"}, [][]string{
		{"URL", fmt.Sprintf("http://localhost:%d", port)},
	})

	cli.Info()
	cli.Info("Forward " + path + " to " + "/workspace")
	cli.Info()

	options := docker.RunOptions{
		Platform: "linux/amd64",

		Ports: []engine.ContainerPort{
			{
				Port:  3000,
				Proto: engine.ProtocolTCP,

				HostPort: &port,
			},
		},

		Volumes: []engine.ContainerMount{
			{
				Path:     "/workspace",
				HostPath: path,
			},
		},

		Stdout: io.Discard,
	}

	return docker.RunInteractive(ctx, image, options)
}
