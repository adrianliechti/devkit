package code

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/app/utility"
	"github.com/adrianliechti/devkit/pkg/engine"
	"github.com/adrianliechti/go-cli"
)

var Command = &cli.Command{
	Name:  "code",
	Usage: "start temporary Code IDE",

	Category: utility.Category,

	Flags: []cli.Flag{
		app.PortFlag(""),
	},

	Action: func(ctx context.Context, cmd *cli.Command) error {
		client := app.MustClient(ctx, cmd)

		items := []string{
			"default",
			"dotnet",
			"golang",
			"java",
			"node",
			"python",
		}

		i, _, err := cli.Select("select stack", items)

		if err != nil {
			return err
		}

		stack := strings.ToLower(items[i])
		port := app.MustPortOrRandom(ctx, cmd, "", 3000)

		return startCode(ctx, client, stack, port)
	},
}

func startCode(ctx context.Context, client engine.Client, stack string, port int) error {
	image := "ghcr.io/adrianliechti/loop-code"

	if stack != "" && stack != "default" {
		image = image + "-" + strings.ToLower(stack)
	}

	path, err := os.Getwd()

	if err != nil {
		return err
	}

	cli.MustRun("Pulling Image...", func() error {
		return client.Pull(ctx, image, "", engine.PullOptions{})
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

	spec := engine.Container{
		Image: image,

		Ports: []engine.ContainerPort{
			{
				Port:  3000,
				Proto: engine.ProtocolTCP,

				HostPort: port,
			},
		},

		Mounts: []engine.ContainerMount{
			{
				Path:     "/workspace",
				HostPath: path,
			},
		},
	}

	return client.Run(ctx, spec, engine.RunOptions{
		Stdout: io.Discard,
		Stderr: io.Discard,
	})
}
