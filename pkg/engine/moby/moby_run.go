package moby

import (
	"context"
	"io"
	"os"

	"github.com/adrianliechti/devkit/pkg/engine"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/pkg/stdcopy"
)

func (m *Moby) Run(ctx context.Context, spec engine.Container, options engine.RunOptions) error {
	if options.Stdin == nil {
		options.Stdin = os.Stdin
	}

	if options.Stdout == nil {
		options.Stdout = os.Stdout
	}

	if options.Stderr == nil {
		options.Stderr = os.Stderr
	}

	containerConfig, err := convertContainerConfig(spec)

	if err != nil {
		return err
	}

	containerConfig.OpenStdin = true
	containerConfig.StdinOnce = true

	containerConfig.AttachStdin = options.Stdin != nil
	containerConfig.AttachStdout = options.Stdout != nil
	containerConfig.AttachStderr = options.Stderr != nil

	hostConfig, err := convertHostConfig(spec)

	if err != nil {
		return err
	}

	if err := m.Pull(ctx, spec.Image, spec.Platform, engine.PullOptions{}); err != nil {
		return err
	}

	created, err := m.client.ContainerCreate(ctx, containerConfig, hostConfig, nil, nil, spec.Name)

	if err != nil {
		return err
	}

	defer m.client.ContainerRemove(context.Background(), created.ID, container.RemoveOptions{
		Force: true,

		RemoveVolumes: true,
	})

	attached, err := m.client.ContainerAttach(ctx, created.ID, convertAttachOptions(options))

	if err != nil {
		return err
	}

	defer attached.Close()

	if err := m.client.ContainerStart(ctx, created.ID, container.StartOptions{}); err != nil {
		return err
	}

	result := make(chan error)

	go func() {
		_, err := io.Copy(attached.Conn, os.Stdin)
		result <- err
	}()

	go func() {
		_, err := stdcopy.StdCopy(options.Stdout, options.Stderr, attached.Reader)
		result <- err
	}()

	return <-result
}

func convertAttachOptions(options engine.RunOptions) container.AttachOptions {
	result := container.AttachOptions{
		Stream: true,

		Stdin:  options.Stdin != nil,
		Stdout: options.Stdout != nil,
		Stderr: options.Stderr != nil,
	}

	return result
}
