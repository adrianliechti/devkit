package moby

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"strings"

	"github.com/adrianliechti/devkit/pkg/container/engine"

	"github.com/cpuguy83/dockercfg"
	"github.com/docker/distribution/reference"
	"github.com/docker/docker/api/types"
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
