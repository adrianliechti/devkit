package catalog

import (
	"fmt"
	"strings"
	"time"

	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/pkg/catalog"
	"github.com/adrianliechti/devkit/pkg/cli"
	"github.com/adrianliechti/devkit/pkg/container"
	"github.com/adrianliechti/devkit/pkg/docker"
)

func Command(m catalog.Manager) *cli.Command {
	kind := kind(m)

	cmd := &cli.Command{
		Name:     kind,
		Category: strings.ToUpper(string(m.Category())),

		HideHelpCommand: true,

		Subcommands: []*cli.Command{
			listCommand(m),

			createCommand(m),
			deleteCommand(m),

			infoCommand(m),
			logsCommand(m),
		},
	}

	if d, ok := m.(catalog.Decorator); ok {
		cmd.Usage = "local " + d.DisplayName()
		cmd.Description = d.Description()
	}

	if p, ok := m.(catalog.ClientProvider); ok {
		cmd.Subcommands = append(cmd.Subcommands, clientCommand(p))
	}

	if p, ok := m.(catalog.ConsoleProvider); ok {
		cmd.Subcommands = append(cmd.Subcommands, consoleCommand(p))
	}

	if p, ok := m.(catalog.ShellProvider); ok {
		cmd.Subcommands = append(cmd.Subcommands, shellCommand(p))
	}

	return cmd
}

func kind(m catalog.Manager) string {
	return strings.ToLower(m.Name())
}

func listCommand(m catalog.Manager) *cli.Command {
	return &cli.Command{
		Name:  "list",
		Usage: "list instances",

		Action: func(c *cli.Context) error {
			list, err := docker.List(c.Context, docker.ListOptions{
				All: true,

				Filter: []string{
					fmt.Sprintf("label=%s=%s", KindKey, kind(m)),
				},
			})

			if err != nil {
				return err
			}

			for _, c := range list {
				name := c.Names[0]
				cli.Info(name)
			}

			return nil
		},
	}
}

func createCommand(m catalog.Manager) *cli.Command {
	ref, _ := m.New()
	kind := kind(m)

	flags := []cli.Flag{}
	portFlags := []*cli.IntFlag{}

	for _, p := range ref.Ports {
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
			spec, err := m.New()

			if err != nil {
				return err
			}

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

			info, err := m.Info(spec)

			if err != nil {
				return err
			}

			printMapTable(info)
			return nil
		},
	}
}

func deleteCommand(m catalog.Manager) *cli.Command {
	kind := kind(m)

	return &cli.Command{
		Name:  "delete",
		Usage: "delete instance",

		Action: func(c *cli.Context) error {
			container := MustContainer(c.Context, kind)

			return docker.Delete(c.Context, container.Name, docker.DeleteOptions{
				Force:   true,
				Volumes: true,
			})
		},
	}
}

func infoCommand(m catalog.Manager) *cli.Command {
	kind := kind(m)

	return &cli.Command{
		Name:  "info",
		Usage: "display instance info",

		Action: func(c *cli.Context) error {
			ctx := c.Context
			container := MustContainer(ctx, kind)

			info, err := m.Info(container)

			if err != nil {
				return err
			}

			printMapTable(info)
			return nil
		},
	}
}

func logsCommand(m catalog.Manager) *cli.Command {
	kind := kind(m)

	return &cli.Command{
		Name:  "logs",
		Usage: "show instance logs",

		Action: func(c *cli.Context) error {
			ctx := c.Context
			container := MustContainer(ctx, kind)

			return docker.Logs(ctx, container.Name, docker.LogsOptions{
				Follow: true,
			})
		},
	}
}

func clientCommand(p catalog.ClientProvider) *cli.Command {
	kind := kind(p)

	return &cli.Command{
		Name:  "cli",
		Usage: "run client in instance",

		Action: func(c *cli.Context) error {
			ctx := c.Context
			container := MustContainer(ctx, kind)

			shell, args, err := p.Client(container)

			if err != nil {
				return err
			}

			return docker.ExecInteractive(ctx, container.Name, docker.ExecOptions{}, shell, args...)
		},
	}
}

func shellCommand(p catalog.ShellProvider) *cli.Command {
	kind := kind(p)

	return &cli.Command{
		Name:  "shell",
		Usage: "run shell in instance",

		Action: func(c *cli.Context) error {
			ctx := c.Context
			container := MustContainer(ctx, kind)

			shell, err := p.Shell(container)

			if err != nil {
				return err
			}

			return docker.ExecInteractive(ctx, container.Name, docker.ExecOptions{}, shell)
		},
	}
}

func consoleCommand(p catalog.ConsoleProvider) *cli.Command {
	kind := kind(p)

	return &cli.Command{
		Name:  "console",
		Usage: "open web console",

		Flags: []cli.Flag{
			app.PortFlag(""),
		},

		Action: func(c *cli.Context) error {
			ctx := c.Context
			container := MustContainer(ctx, kind)

			info, err := p.Info(container)

			if err != nil {
				return err
			}

			mapping, err := p.ConsolePort(container)

			if err != nil {
				return err
			}

			port := app.MustPortOrRandom(c, "", mapping.Port)

			time.AfterFunc(1*time.Second, func() {
				cli.OpenURL(fmt.Sprintf("http://localhost:%d", port))
			})

			printMapTable(info)

			return docker.PortForward(c.Context, container.Name, port, mapping.Port)
		},
	}
}
