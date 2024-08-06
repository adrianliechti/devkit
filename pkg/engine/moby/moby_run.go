package moby

import (
	"context"
	"io"
	"os"
	"time"

	"github.com/adrianliechti/devkit/pkg/engine"
	"github.com/adrianliechti/devkit/pkg/system"
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

	isTTY := system.IsTerminal(options.Stdout)

	if isTTY {
		restore, err := system.MakeRawTerminal(options.Stdout)

		if err != nil {
			return err
		}

		defer restore()
	}

	containerConfig, err := convertContainerConfig(spec)

	if err != nil {
		return err
	}

	containerConfig.Tty = isTTY

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

	if isTTY {
		width, height, err := system.TerminalSize(options.Stdout)

		if err != nil {
			return err
		}

		m.client.ContainerResize(ctx, created.ID, container.ResizeOptions{
			Width:  uint(width),
			Height: uint(height),
		})

		go func() {
			for ctx.Err() == nil {
				time.Sleep(200 * time.Millisecond)

				w, h, err := system.TerminalSize(options.Stdout)

				if err != nil {
					continue
				}

				if w == width && h == height {
					continue
				}

				width = w
				height = h

				m.client.ContainerResize(ctx, created.ID, container.ResizeOptions{
					Width:  uint(width),
					Height: uint(height),
				})
			}
		}()
	}

	result := make(chan error)

	go func() {
		_, err := io.Copy(attached.Conn, options.Stdin)
		result <- err
	}()

	go func() {
		if isTTY {
			_, err := io.Copy(options.Stdout, attached.Reader)
			result <- err
			return
		}

		_, err := stdcopy.StdCopy(options.Stdout, options.Stderr, attached.Reader)
		result <- err
	}()

	select {
	case err := <-result:
		return err
	case <-ctx.Done():
		return nil
	}
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
