package moby

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

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

	isTTY := system.IsTerminal(options.Stdin)

	if isTTY {
		restore, err := system.MakeRawTerminal(options.Stdout)

		if err != nil {
			return err
		}

		defer restore()
	}

	id, err := m.client.ContainerExecCreate(ctx, containerID, convertExecOptions(command, options))

	if err != nil {
		return nil
	}

	resp, err := m.client.ContainerExecAttach(ctx, id.ID, convertExecAttachOptions(options))

	if err != nil {
		return nil
	}

	defer resp.Close()

	if isTTY {
		width, height, err := system.TerminalSize(options.Stdout)

		if err != nil {
			return err
		}

		m.client.ContainerResize(ctx, id.ID, container.ResizeOptions{
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

				m.client.ContainerResize(ctx, id.ID, container.ResizeOptions{
					Width:  uint(width),
					Height: uint(height),
				})
			}
		}()
	}

	result := make(chan error)

	go func() {
		_, err := io.Copy(resp.Conn, options.Stdin)
		result <- err
	}()

	go func() {
		if isTTY {
			_, err := io.Copy(options.Stdout, resp.Reader)
			result <- err
			return
		}

		_, err := stdcopy.StdCopy(options.Stdout, options.Stderr, resp.Reader)
		result <- err
	}()

	select {
	case err := <-result:
		return err
	case <-ctx.Done():
		return nil
	}
}

func convertExecOptions(command []string, options engine.ExecOptions) container.ExecOptions {
	tty := system.IsTerminal(options.Stdout)

	result := container.ExecOptions{
		Cmd: command,

		Privileged: options.Privileged,

		Tty: tty,

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

func convertExecAttachOptions(options engine.ExecOptions) container.ExecAttachOptions {
	tty := system.IsTerminal(options.Stdout)

	result := container.ExecAttachOptions{
		Tty: tty,
	}

	if w, h, err := system.TerminalSize(options.Stdout); err == nil {
		size := [2]uint{uint(h), uint(w)}
		result.ConsoleSize = &size
	}

	return result
}
