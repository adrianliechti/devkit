package moby

import (
	"context"
	"io"
	"os"
	"os/signal"
	"syscall"

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

	isTTY := IsTerminal(options.Stdin)

	if isTTY {
		state, err := MakeRawTerminal(options.Stdout)

		if err != nil {
			return err
		}

		defer RestoreTerminal(options.Stdout, state)
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
		if w, h, err := TerminalSize(options.Stdout); err == nil {
			m.client.ContainerResize(ctx, created.ID, container.ResizeOptions{
				Width:  uint(w),
				Height: uint(h),
			})
		}

		go func() {
			sigs := make(chan os.Signal, 1)
			signal.Notify(sigs, syscall.SIGWINCH)

			for range sigs {
				if w, h, err := TerminalSize(options.Stdout); err == nil {
					m.client.ContainerResize(ctx, created.ID, container.ResizeOptions{
						Width:  uint(w),
						Height: uint(h),
					})
				}
			}
		}()
	}

	// func() {
	// 	for ctx.Err() == nil {
	// 		if !tty {
	// 			break
	// 		}

	// 		time.Sleep(1 * time.Second)

	// 		var width int
	// 		var height int

	// 		if w, h, err := TerminalSize(options.Stdin); err == nil {
	// 			if w == width && h == height {
	// 				continue
	// 			}

	// 			m.client.ContainerResize(ctx, created.ID, container.ResizeOptions{
	// 				Height: uint(h),
	// 				Width:  uint(w),
	// 			})
	// 		}
	// 	}
	// }()

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
