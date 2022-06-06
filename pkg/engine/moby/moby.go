package moby

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"

	"github.com/adrianliechti/devkit/pkg/engine"

	"github.com/cpuguy83/dockercfg"
	"github.com/docker/distribution/reference"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

var (
	_ engine.Engine = &Moby{}
)

type Moby struct {
	client *client.Client
}

func New() (*Moby, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	if err != nil {
		return nil, err
	}

	return &Moby{
		client: cli,
	}, nil
}

func (m *Moby) List(ctx context.Context, options engine.ListOptions) ([]engine.Container, error) {
	opts := types.ContainerListOptions{
		Quiet: true,

		All:     options.All,
		Filters: filters.NewArgs(),
	}

	for k, v := range options.LabelSelector {
		opts.Filters.Add("label", fmt.Sprintf("%s=%s", k, v))
	}

	list, err := m.client.ContainerList(ctx, opts)

	if err != nil {
		return nil, err
	}

	containers := []engine.Container{}

	for _, i := range list {
		container, err := m.Inspect(ctx, i.ID)

		if err != nil {
			return nil, err
		}

		containers = append(containers, container)
	}

	return containers, nil
}

func (m *Moby) Pull(ctx context.Context, image string, options engine.PullOptions) error {
	if options.Stdout == nil {
		options.Stdout = io.Discard
	}

	if options.Stderr == nil {
		options.Stderr = io.Discard
	}

	out, err := m.client.ImagePull(ctx, image, types.ImagePullOptions{
		Platform:     options.Platform,
		RegistryAuth: registryCredentials(image),
	})

	if err != nil {
		return err
	}

	defer out.Close()

	if _, err := io.Copy(options.Stdout, out); err != nil {
		return err
	}

	return nil
}

func (m *Moby) Remove(ctx context.Context, container string, options engine.RemoveOptions) error {
	return m.client.ContainerRemove(ctx, container, types.ContainerRemoveOptions{
		Force: true,

		RemoveVolumes: true,
	})
}

func (m *Moby) Inspect(ctx context.Context, container string) (engine.Container, error) {
	data, err := m.client.ContainerInspect(ctx, container)

	if err != nil {
		return engine.Container{}, err
	}

	return convertContainer(data), nil
}

func (m *Moby) Logs(ctx context.Context, container string, options engine.LogsOptions) error {
	if options.Stdout == nil {
		options.Stdout = io.Discard
	}

	if options.Stderr == nil {
		options.Stderr = io.Discard
	}

	out, err := m.client.ContainerLogs(ctx, container, types.ContainerLogsOptions{
		Follow: options.Follow,

		ShowStdout: true,
		ShowStderr: true,
	})

	if err != nil {
		return err
	}

	defer out.Close()

	if _, err := stdcopy.StdCopy(options.Stdout, options.Stderr, out); err != nil {
		return err
	}

	return nil
}

func registryCredentials(image string) string {
	// echo "https://index.docker.io/v1/"|docker-credential-desktop get
	// {"ServerURL":"https://index.docker.io/v1/","Username":"xxxxx","Secret":"xxxxx"}

	ref, err := reference.ParseNormalizedNamed(image)

	if err != nil {
		return ""
	}

	parts := strings.Split(ref.Name(), "/")

	if len(parts) == 0 {
		return ""
	}

	host := dockercfg.ResolveRegistryHost(parts[0])

	username, password, err := dockercfg.GetRegistryCredentials(host)

	if err != nil {
		return ""
	}

	data, err := json.Marshal(types.AuthConfig{
		Username: username,
		Password: password,
	})

	if err != nil {
		return ""
	}

	return base64.URLEncoding.EncodeToString(data)
}

func convertContainer(data types.ContainerJSON) engine.Container {
	container := engine.Container{
		ID:   data.ID,
		Name: strings.TrimLeft(data.Name, "/"),

		Labels: data.Config.Labels,

		Image: data.Config.Image,

		Privileged: data.HostConfig.Privileged,

		Env: map[string]string{},
		Dir: data.Config.WorkingDir,

		Command: data.Config.Entrypoint,
		Args:    data.Config.Cmd,

		Hostname:  data.Config.Hostname,
		IPAddress: net.ParseIP(data.NetworkSettings.IPAddress),

		Ports:  []*engine.ContainerPort{},
		Mounts: []*engine.ContainerMount{},
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
			port := &engine.ContainerPort{
				Port:  p.Int(),
				Proto: engine.Protocol(p.Proto()),
			}

			if b.HostPort != "" {
				if val, err := strconv.Atoi(b.HostPort); err == nil {
					port.HostIP = b.HostIP
					port.HostPort = &val
				}
			}

			container.Ports = append(container.Ports, port)
		}
	}

	for _, m := range data.Mounts {
		mount := &engine.ContainerMount{
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

	return container
}
