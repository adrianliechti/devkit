package catalog

import (
	"fmt"

	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/container"
	"github.com/adrianliechti/devkit/pkg/docker"
)

type CreateHandler = func() container.Container

func CreateCommand(kind string, create CreateHandler, info InfoHandler) *cli.Command {
	flags := []cli.Flag{}
	portFlags := []*cli.IntFlag{}

	for _, p := range create().Ports {
		port := p.Port
		name := p.Name

		if name == "" {
			name = "port"
		}

		proto := p.Protocol

		if proto == "" {
			proto = container.ProtocolTCP
		}

		f := app.PortFlag(name)
		f.DefaultText = fmt.Sprintf("%d (%s) or random", port, proto)

		flags = append(flags, f)
		portFlags = append(portFlags, f)
	}

	return &cli.Command{
		Name:  "create",
		Usage: "create instance",

		Flags: flags,

		Action: func(c *cli.Context) error {
			ctx := c.Context
			spec := create()

			spec.Labels = map[string]string{
				KindKey: kind,
			}

			for _, p := range spec.Ports {
				flag := app.PortFlagName(p.Name)

				for _, f := range portFlags {
					if f.Name != flag {
						continue
					}

					hostIP := "127.0.0.1"
					hostPort := app.MustPortOrRandom(c, f.Name, p.Port)

					p.HostIP = hostIP
					p.HostPort = &hostPort
				}
			}

			options := convertRunOptions(spec)

			if err := docker.Run(ctx, spec.Image, options, spec.Args...); err != nil {
				return err
			}

			info := info(&spec)
			printMapTable(info)

			return nil
		},
	}
}
