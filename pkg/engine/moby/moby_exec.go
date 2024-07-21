package moby

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"

	"github.com/adrianliechti/devkit/pkg/engine"
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

	isTTY := IsTerminal(options.Stdin)

	if isTTY {
		restore, err := MakeRawTerminal(options.Stdout)

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
		if w, h, err := TerminalSize(options.Stdout); err == nil {
			m.client.ContainerExecResize(ctx, id.ID, container.ResizeOptions{
				Width:  uint(w),
				Height: uint(h),
			})
		}

		go func() {
			sigs := make(chan os.Signal, 1)
			signal.Notify(sigs, syscall.SIGWINCH)

			for range sigs {
				if w, h, err := TerminalSize(options.Stdout); err == nil {
					m.client.ContainerExecResize(ctx, id.ID, container.ResizeOptions{
						Width:  uint(w),
						Height: uint(h),
					})
				}
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
	tty := IsTerminal(options.Stdout)

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

	if w, h, err := TerminalSize(options.Stdout); err == nil {
		size := [2]uint{uint(h), uint(w)}
		result.ConsoleSize = &size
	}

	return result
}

func convertExecAttachOptions(options engine.ExecOptions) container.ExecAttachOptions {
	tty := IsTerminal(options.Stdout)

	result := container.ExecAttachOptions{
		Tty: tty,
	}

	if w, h, err := TerminalSize(options.Stdout); err == nil {
		size := [2]uint{uint(h), uint(w)}
		result.ConsoleSize = &size
	}

	return result
}
