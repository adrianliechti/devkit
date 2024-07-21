package moby

import (
	"context"
	"fmt"
	"strconv"

	"github.com/adrianliechti/devkit/pkg/engine"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/go-connections/nat"
	"github.com/docker/go-units"
)

func (m *Moby) Create(ctx context.Context, spec engine.Container, options engine.CreateOptions) (string, error) {
	containerConfig, err := convertContainerConfig(spec)

	if err != nil {
		return "", err
	}

	hostConfig, err := convertHostConfig(spec)

	if err != nil {
		return "", err
	}

	if err := m.Pull(ctx, spec.Image, spec.Platform, engine.PullOptions{}); err != nil {
		return "", err
	}

	resp, err := m.client.ContainerCreate(ctx, containerConfig, hostConfig, nil, nil, spec.Name)

	if err != nil {
		return "", err
	}

	if err := m.client.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		return "", err
	}

	return resp.ID, nil
}

func convertContainerConfig(spec engine.Container) (*container.Config, error) {
	config := &container.Config{
		Labels: spec.Labels,

		Image: spec.Image,

		Env:        []string{},
		WorkingDir: spec.Dir,

		Entrypoint: spec.Command,
		Cmd:        spec.Args,

		Hostname: spec.Hostname,

		ExposedPorts: nat.PortSet{},
	}

	if spec.RunAsUser != "" {
		user := spec.RunAsUser

		if spec.RunAsGroup != "" {
			user += ":" + spec.RunAsGroup
		}

		config.User = user
	}

	for k, v := range spec.Env {
		config.Env = append(config.Env, fmt.Sprintf("%s=%s", k, v))
	}

	for _, p := range spec.Ports {
		proto := string(p.Proto)

		if proto == "" {
			proto = "tcp"
		}

		port, err := nat.NewPort(proto, strconv.Itoa(p.Port))

		if err != nil {
			return nil, err
		}

		config.ExposedPorts[port] = struct{}{}
	}

	return config, nil
}

func convertHostConfig(spec engine.Container) (*container.HostConfig, error) {
	config := &container.HostConfig{
		Privileged: spec.Privileged,

		PortBindings: nat.PortMap{},
		Mounts:       []mount.Mount{},
	}

	for _, p := range spec.Ports {
		proto := string(p.Proto)

		if proto == "" {
			proto = "tcp"
		}

		port, err := nat.NewPort(proto, strconv.Itoa(p.Port))

		if err != nil {
			return nil, err
		}

		binding := nat.PortBinding{
			HostIP: p.HostIP,
		}

		if binding.HostIP == "" {
			binding.HostIP = "127.0.0.1"
		}

		if p.HostPort != 0 {
			binding.HostPort = strconv.Itoa(p.HostPort)
		}

		config.PortBindings[port] = []nat.PortBinding{
			binding,
		}
	}

	for _, m := range spec.Mounts {
		if m.Volume != "" {
			config.Mounts = append(config.Mounts, mount.Mount{
				Type:   mount.TypeVolume,
				Target: m.Path,
				Source: m.Volume,
			})
		}

		if m.HostPath != "" {
			config.Mounts = append(config.Mounts, mount.Mount{
				Type:   mount.TypeBind,
				Target: m.Path,
				Source: m.HostPath,
			})
		}
	}

	ulimits := []*units.Ulimit{}

	if spec.MaxFiles != 0 {
		ulimits = append(ulimits, &units.Ulimit{
			Name: "nofile",
			Soft: spec.MaxFiles,
			Hard: spec.MaxFiles,
		})
	}

	if spec.MaxProcesses != 0 {
		ulimits = append(ulimits, &units.Ulimit{
			Name: "nproc",
			Soft: spec.MaxProcesses,
			Hard: spec.MaxProcesses,
		})
	}

	config.Ulimits = ulimits

	return config, nil
}
