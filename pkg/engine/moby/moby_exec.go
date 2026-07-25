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

func (m *Moby) Exec(ctx context.Context, containerID string, command []string, options engine.ExecOptions) error {
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
	// This single decision drives the exec config, the attach options and the
	// raw-mode/demultiplexing below.
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

	id, err := m.client.ContainerExecCreate(ctx, containerID, convertExecOptions(command, options, isTTY))

	if err != nil {
		return err
	}

	resp, err := m.client.ContainerExecAttach(ctx, id.ID, container.ExecAttachOptions{Tty: isTTY})

	if err != nil {
		return err
	}

	defer resp.Close()

	if isTTY {
		syncTerminalSize(ctx, options.Stdout, func(width, height int) error {
			return m.client.ContainerResize(ctx, id.ID, container.ResizeOptions{
				Width:  uint(width),
				Height: uint(height),
			})
		})
	}

	go func() {
		io.Copy(resp.Conn, options.Stdin)
		resp.CloseWrite()
	}()

	done := make(chan error, 1)

	go func() {
		if isTTY {
			_, err := io.Copy(options.Stdout, resp.Reader)
			done <- err
			return
		}

		_, err := stdcopy.StdCopy(options.Stdout, options.Stderr, resp.Reader)
		done <- err
	}()

	select {
	case err := <-done:
		if err != nil {
			// A cancelled context tears the stream down; that is not a failure.
			if ctx.Err() != nil {
				return nil
			}

			return err
		}
	case <-ctx.Done():
		return nil
	}

	// Surface the command's exit code instead of reporting success unconditionally.
	inspect, err := m.client.ContainerExecInspect(context.WithoutCancel(ctx), id.ID)

	if err != nil {
		return err
	}

	if inspect.ExitCode != 0 {
		return fmt.Errorf("command exited with code %d", inspect.ExitCode)
	}

	return nil
}

func convertExecOptions(command []string, options engine.ExecOptions, isTTY bool) container.ExecOptions {
	result := container.ExecOptions{
		Cmd: command,

		Privileged: options.Privileged,

		Tty: isTTY,

		AttachStdin:  options.Stdin != nil,
		AttachStdout: options.Stdout != nil,
		AttachStderr: options.Stderr != nil,

		User:       options.User,
		WorkingDir: options.Dir,
	}

	for k, v := range options.Env {
		result.Env = append(result.Env, fmt.Sprintf("%s=%s", k, v))
	}

	if w, h, err := system.TerminalSize(options.Stdout); err == nil {
		size := [2]uint{uint(h), uint(w)}
		result.ConsoleSize = &size
	}

	return result
}
