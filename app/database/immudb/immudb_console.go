package immudb

import (
	"fmt"
	"time"

	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/app/common"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/docker"
)

func ConsoleCommand() *cli.Command {
	return &cli.Command{
		Name:  "console",
		Usage: "open web console",

		Flags: []cli.Flag{
			app.PortFlag(""),
		},

		Action: func(c *cli.Context) error {
			ctx := c.Context
			container := common.MustContainer(ctx, ImmuDB)

			port := app.MustPortOrRandom(c, "", 8080)

			time.AfterFunc(1*time.Second, func() {
				cli.OpenURL(fmt.Sprintf("http://localhost:%d", port))
			})

			return docker.PortForward(c.Context, container, port, 8080)
		},
	}
}
