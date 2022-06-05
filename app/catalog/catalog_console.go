package catalog

import (
	"fmt"
	"time"

	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/container"
	"github.com/adrianliechti/devkit/pkg/docker"
)

type ConsoleHandler = func() container.ContainerPort

func ConsoleCommand(kind string, infoHandler InfoHandler, consoleHandler ConsoleHandler) *cli.Command {
	return &cli.Command{
		Name:  "console",
		Usage: "open web console",

		Action: func(c *cli.Context) error {
			ctx := c.Context
			container := MustContainer(ctx, kind)

			info := infoHandler(&container)
			mapping := consoleHandler()

			port := app.MustPortOrRandom(c, "", mapping.Port)

			time.AfterFunc(1*time.Second, func() {
				cli.OpenURL(fmt.Sprintf("http://localhost:%d", port))
			})

			printMapTable(info)

			return docker.PortForward(c.Context, container.Name, port, mapping.Port)
		},
	}
}
