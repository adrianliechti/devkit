package catalog

import (
	"context"
	"fmt"
	"strings"

	"github.com/adrianliechti/devkit/app"
	"github.com/adrianliechti/devkit/catalog"
	"github.com/adrianliechti/devkit/pkg/engine"
	"github.com/adrianliechti/go-cli"
)

func Command(m catalog.Manager) *cli.Command {
	kind := kind(m)

	cmd := &cli.Command{
		Name:     kind,
		Category: strings.ToUpper(string(m.Category())),

		HideHelpCommand: true,

		Commands: []*cli.Command{
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
		cmd.Commands = append(cmd.Commands, clientCommand(p))
	}

	if p, ok := m.(catalog.ConsoleProvider); ok {
		cmd.Commands = append(cmd.Commands, consoleCommand(p))
	}

	if p, ok := m.(catalog.ShellProvider); ok {
		cmd.Commands = append(cmd.Commands, shellCommand(p))
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

		Action: func(ctx context.Context, cmd *cli.Command) error {
			client := app.MustClient(ctx, cmd)

			containers, err := client.List(ctx, engine.ListOptions{
				All: true,

				LabelSelector: map[string]string{
					KindKey: kind(m),
				},
			})

			if err != nil {
				return err
			}

			for _, container := range containers {
				cli.Info(container.Name)
			}

			return nil
		},
	}
}

func createCommand(m catalog.Manager) *cli.Command {
	ref, _ := m.New()
	kind := kind(m)

	flags := []cli.Flag{
		&cli.StringFlag{
			Name:  "name",
			Usage: "module name",
		},
	}

	portFlags := []*cli.IntFlag{}

	for _, p := range ref.Ports {
		port := p.Port
		name := p.Name

		if name == "" {
			name = "port"
		}

		proto := string(p.Proto)

		if proto == "" {
			proto = string(engine.ProtocolTCP)
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

		Action: func(ctx context.Context, cmd *cli.Command) error {
			client := app.MustClient(ctx, cmd)
			container, err := m.New()

			if err != nil {
				return err
			}

			cli.MustRun("Pulling Image...", func() error {
				client.Pull(ctx, container.Image, container.Platform, engine.PullOptions{})
				return nil
			})

			if name := cmd.String("name"); name != "" {
				container.Name = name
			}

			container.Labels = map[string]string{
				KindKey: kind,
			}

			for i, p := range container.Ports {
				flag := app.PortFlagName(p.Name)

				for _, f := range portFlags {
					if f.Name != flag {
						continue
					}

					hostIP := "127.0.0.1"
					hostPort := app.MustPortOrRandom(ctx, cmd, f.Name, p.Port)

					p.HostIP = hostIP
					p.HostPort = hostPort
				}

				container.Ports[i] = p
			}

			var containerID string

			cli.MustRun("Creating Container...", func() error {
				id, err := client.Create(ctx, container, engine.CreateOptions{})

				if err != nil {
					return err
				}

				containerID = id
				return nil
			})

			container, err = client.Inspect(ctx, containerID)

			if err != nil {
				return err
			}

			info, err := m.Info(container)

			if err != nil {
				return err
			}

			printContainerInfo(container, info)
			return nil
		},
	}
}

func deleteCommand(m catalog.Manager) *cli.Command {
	kind := kind(m)

	return &cli.Command{
		Name:  "delete",
		Usage: "delete instance",

		Action: func(ctx context.Context, cmd *cli.Command) error {
			client := app.MustClient(ctx, cmd)
			container := MustContainer(ctx, client, kind, true)

			cli.MustRun("Deleting Container...", func() error {
				client.Delete(ctx, container.ID, engine.DeleteOptions{})
				return nil
			})

			return nil
		},
	}
}

func infoCommand(m catalog.Manager) *cli.Command {
	kind := kind(m)

	return &cli.Command{
		Name:  "info",
		Usage: "show instance info",

		Action: func(ctx context.Context, cmd *cli.Command) error {
			client := app.MustClient(ctx, cmd)
			container := MustContainer(ctx, client, kind, true)

			info, err := m.Info(container)

			if err != nil {
				return err
			}

			printContainerInfo(container, info)
			return nil
		},
	}
}

func logsCommand(m catalog.Manager) *cli.Command {
	kind := kind(m)

	return &cli.Command{
		Name:  "logs",
		Usage: "show instance logs",

		Action: func(ctx context.Context, cmd *cli.Command) error {
			client := app.MustClient(ctx, cmd)
			container := MustContainer(ctx, client, kind, true)

			return client.Logs(ctx, container.ID, engine.LogsOptions{
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

		Action: func(ctx context.Context, cmd *cli.Command) error {
			client := app.MustClient(ctx, cmd)
			container := MustContainer(ctx, client, kind, true)

			image, args, err := p.Client(container)

			if err != nil {
				return err
			}

			if image == "" {
				return client.Exec(ctx, container.ID, args, engine.ExecOptions{})
			}

			cli.MustRun("Pulling Image...", func() error {
				client.Pull(ctx, image, "", engine.PullOptions{})
				return nil
			})

			spec := engine.Container{
				Image: image,

				Args: args,
			}

			return client.Run(ctx, spec, engine.RunOptions{})
		},
	}
}

func shellCommand(p catalog.ShellProvider) *cli.Command {
	kind := kind(p)

	return &cli.Command{
		Name:  "shell",
		Usage: "run shell in instance",

		Action: func(ctx context.Context, cmd *cli.Command) error {
			client := app.MustClient(ctx, cmd)
			container := MustContainer(ctx, client, kind, true)

			shell, err := p.Shell(container)

			if err != nil {
				return err
			}

			return client.Exec(ctx, container.ID, []string{shell}, engine.ExecOptions{})
		},
	}
}

func consoleCommand(p catalog.ConsoleProvider) *cli.Command {
	kind := kind(p)

	return &cli.Command{
		Name:  "console",
		Usage: "run console in browser",

		Flags: []cli.Flag{
			app.PortFlag(""),
		},

		Action: func(ctx context.Context, cmd *cli.Command) error {
			client := app.MustClient(ctx, cmd)
			container := MustContainer(ctx, client, kind, true)

			info, err := p.Info(container)

			if err != nil {
				return err
			}

			mapping, err := p.ConsolePort(container)

			if err != nil {
				return err
			}

			port := app.MustPortOrRandom(ctx, cmd, "", mapping.Port)
			ready := make(chan struct{})

			println("mapping", port)

			go func() {
				<-ready
				printContainerInfo(container, info)

				cli.OpenURL(fmt.Sprintf("http://localhost:%d", port))
			}()

			ports := map[int]int{
				port: mapping.Port,
			}

			return client.PortForward(ctx, container.Name, "", ports, ready)
		},
	}
}
