package moby

import (
	"context"
	"io"
	"os"
	"time"

	"github.com/adrianliechti/devkit/pkg/engine"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/pkg/stdcopy"
)

func (m *Moby) Run(ctx context.Context, spec engine.Container, options engine.RunOptions) error {
	if options.Stdout == nil {
		options.Stdout = io.Discard
	}

	if options.Stderr == nil {
		options.Stderr = io.Discard
	}

	containerConfig, err := convertContainerConfig(spec)

	if err != nil {
		return err
	}

	containerConfig.Tty = options.TTY

	containerConfig.AttachStdin = options.Stdin != nil
	containerConfig.AttachStdout = options.Stdout != nil
	containerConfig.AttachStderr = options.Stderr != nil

	if options.Interactive {
		containerConfig.OpenStdin = true
		containerConfig.StdinOnce = true
	}

	hostConfig, err := convertHostConfig(spec)

	if err != nil {
		return err
	}

	if err := m.Pull(ctx, spec.Image, spec.Platform, engine.PullOptions{}); err != nil {
		return err
	}

	create, err := m.client.ContainerCreate(ctx, containerConfig, hostConfig, nil, nil, spec.Name)

	if err != nil {
		return err
	}

	if err := m.client.ContainerStart(ctx, create.ID, container.StartOptions{}); err != nil {
		return err
	}

	containerID := create.ID

	// containerID, err := m.Create(ctx, spec, engine.CreateOptions{})

	// if err != nil {
	// 	return err
	// }

	// statusCh, errCh := m.client.ContainerWait(ctx, containerID, container.WaitConditionNotRunning)

	// select {
	// case err := <-errCh:
	// 	if err != nil {
	// 		return err
	// 	}
	// case <-statusCh:
	// }

	time.Sleep(5 * time.Second)

	resp, err := m.client.ContainerAttach(ctx, containerID, convertAttachOptions(options))

	if err != nil {
		return err
	}

	defer resp.Close()

	result := make(chan error)

	go func() {
		_, err := io.Copy(resp.Conn, os.Stdin)
		result <- err
	}()

	go func() {
		_, err := stdcopy.StdCopy(options.Stdout, options.Stderr, resp.Reader)
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
