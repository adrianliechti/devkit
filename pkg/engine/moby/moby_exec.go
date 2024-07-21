package moby

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/adrianliechti/devkit/pkg/engine"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/pkg/stdcopy"
)

func (m *Moby) Exec(ctx context.Context, containerID string, command []string, options engine.ExecOptions) error {
	if options.Stdout == nil {
		options.Stdout = io.Discard
	}

	if options.Stderr == nil {
		options.Stderr = io.Discard
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

func convertExecOptions(command []string, options engine.ExecOptions) container.ExecOptions {
	result := container.ExecOptions{
		Cmd: command,

		Privileged: options.Privileged,

		Tty:    options.TTY,
		Detach: !options.Interactive,

		AttachStdin:  options.Stdin != nil,
		AttachStdout: options.Stdout != nil,
		AttachStderr: options.Stderr != nil,

		User:       options.User,
		WorkingDir: options.Dir,
	}

	for k, v := range options.Env {
		result.Env = append(result.Env, fmt.Sprintf("%s=%s", k, v))
	}

	return result
}

func convertExecAttachOptions(options engine.ExecOptions) container.ExecAttachOptions {
	result := container.ExecAttachOptions{}

	return result
}
