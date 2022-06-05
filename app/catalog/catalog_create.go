package catalog

import (
	"fmt"

	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/container"
)

type CreateHandler = func() container.Container

func CreateCommand(kind string, h CreateHandler) *cli.Command {
	spec := h()

	spec.Labels = map[string]string{
		KindKey: kind,
	}

	flags := []cli.Flag{}

	for _, p := range spec.Ports {
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
	}

	return &cli.Command{
		Name:  "create",
		Usage: "create instance",

		Flags: flags,

		Action: func(c *cli.Context) error {
			ports := []container.ContainerPort{}

			for _, p := range spec.Ports {
				name := p.Name

				if name == "" {
					name = "port"
				}

				proto := p.Protocol

				if proto == "" {
					proto = container.ProtocolTCP
				}

				hostIP := "127.0.0.1"
				hostPort := app.MustPortOrRandom(c, name, p.Port)

				port := container.ContainerPort{
					Port: p.Port,

					HostIP:   hostIP,
					HostPort: int32(hostPort),
				}

				ports = append(ports, port)
			}

			cli.Info("created")

			// if err := docker.Run(ctx, image, options); err != nil {
			// 	return err
			// }

			// cli.Table([]string{"Name", "Value"}, [][]string{
			// 	{"Host", fmt.Sprintf("localhost:%d", port)},
			// 	{"database", database},
			// 	{"Username", username},
			// 	{"Password", password},
			// })

			return nil
		},
	}
}
