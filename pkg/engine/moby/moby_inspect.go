package moby

import (
	"context"
	"net"
	"strconv"
	"strings"

	"github.com/adrianliechti/devkit/pkg/engine"
	"github.com/docker/docker/api/types"
)

func (m *Moby) Inspect(ctx context.Context, container string) (engine.Container, error) {
	data, err := m.client.ContainerInspect(ctx, container)

	if err != nil {
		return engine.Container{}, err
	}

	return convertContainer(data), nil
}

func convertContainer(data types.ContainerJSON) engine.Container {
	container := engine.Container{
		ID:   data.ID,
		Name: strings.TrimLeft(data.Name, "/"),

		Labels: data.Config.Labels,

		Image:    data.Config.Image,
		Platform: data.Platform,

		Privileged: data.HostConfig.Privileged,

		Env: map[string]string{},
		Dir: data.Config.WorkingDir,

		Command: data.Config.Entrypoint,
		Args:    data.Config.Cmd,

		Hostname:  data.Config.Hostname,
		IPAddress: net.ParseIP(data.NetworkSettings.IPAddress),

		Ports:  []engine.ContainerPort{},
		Mounts: []engine.ContainerMount{},
	}

	if data.Config.User != "" {
		s := strings.Split(data.Config.User, ":")

		if len(s) > 0 {
			container.RunAsUser = s[0]
		}

		if len(s) > 1 {
			container.RunAsGroup = s[1]
		}
	}

	for _, e := range data.Config.Env {
		s := strings.SplitN(e, "=", 2)

		key := s[0]
		val := s[1]

		container.Env[key] = val
	}

	for p, m := range data.NetworkSettings.Ports {
		for _, b := range m {
			port := engine.ContainerPort{
				Port:  p.Int(),
				Proto: engine.Protocol(p.Proto()),
			}

			if val, err := strconv.Atoi(b.HostPort); err == nil {
				port.HostIP = b.HostIP
				port.HostPort = val
			}

			container.Ports = append(container.Ports, port)
		}
	}

	for _, m := range data.Mounts {
		mount := engine.ContainerMount{
			Path: m.Destination,
		}

		if m.Type == "bind" {
			mount.HostPath = m.Source
		}

		if m.Type == "volume" {
			mount.Volume = m.Name
		}

		container.Mounts = append(container.Mounts, mount)
	}

	for _, l := range data.HostConfig.Ulimits {
		if l.Name == "nofile" {
			container.MaxFiles = l.Hard
		}

		if l.Name == "nproc" {
			container.MaxProcesses = l.Hard
		}
	}

	return container
}
