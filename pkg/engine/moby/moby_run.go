package moby

import (
	"context"
	"fmt"
	"io"
	"os"

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

	// A TTY needs both ends to be a terminal: stdin because it gets put into raw
	// mode, stdout because a TTY stream is not multiplexed and carries escapes.
	// This single decision drives the container config, the attach options and
	// the raw-mode/demultiplexing below.
	isTTY := system.IsTerminal(options.Stdin) && system.IsTerminal(options.Stdout)

	if isTTY {
		restore, err := system.MakeRawTerminal(options.Stdin)

		if err != nil {
			// Some consoles report a terminal but refuse raw mode; run without a TTY.
			isTTY = false
		} else {
			defer restore()
		}
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

	defer m.client.ContainerRemove(context.WithoutCancel(ctx), created.ID, container.RemoveOptions{
		Force: true,

		RemoveVolumes: true,
	})

	attached, err := m.client.ContainerAttach(ctx, created.ID, container.AttachOptions{
		Stream: true,

		Stdin:  options.Stdin != nil,
		Stdout: options.Stdout != nil,
		Stderr: options.Stderr != nil,
	})

	if err != nil {
		return err
	}

	defer attached.Close()

	// Start waiting before starting the container, so a short-lived container
	// cannot exit before the wait is established.
	waitCh, waitErrCh := m.client.ContainerWait(ctx, created.ID, container.WaitConditionNextExit)

	if err := m.client.ContainerStart(ctx, created.ID, container.StartOptions{}); err != nil {
		return err
	}

	if isTTY {
		syncTerminalSize(ctx, options.Stdout, func(width, height int) error {
			return m.client.ContainerResize(ctx, created.ID, container.ResizeOptions{
				Width:  uint(width),
				Height: uint(height),
			})
		})
	}

	go func() {
		io.Copy(attached.Conn, options.Stdin)
		attached.CloseWrite()
	}()

	copied := make(chan error, 1)

	go func() {
		if isTTY {
			_, err := io.Copy(options.Stdout, attached.Reader)
			copied <- err
			return
		}

		_, err := stdcopy.StdCopy(options.Stdout, options.Stderr, attached.Reader)
		copied <- err
	}()

	if err := <-copied; err != nil {
		// A cancelled context tears the stream down; that is not a failure.
		if ctx.Err() != nil {
			return nil
		}

		return err
	}

	// Surface the container's exit code instead of reporting success unconditionally.
	select {
	case result := <-waitCh:
		if result.Error != nil && result.Error.Message != "" {
			return fmt.Errorf("%s", result.Error.Message)
		}

		if result.StatusCode != 0 {
			return fmt.Errorf("container exited with code %d", result.StatusCode)
		}

		return nil

	case err := <-waitErrCh:
		if ctx.Err() != nil {
			return nil
		}

		return err

	case <-ctx.Done():
		return nil
	}
}
