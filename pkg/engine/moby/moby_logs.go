package moby

import (
	"context"
	"io"

	"github.com/adrianliechti/devkit/pkg/engine"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/pkg/stdcopy"
)

func (m *Moby) Logs(ctx context.Context, containerID string, options engine.LogsOptions) error {
	if options.Stdout == nil {
		options.Stdout = io.Discard
	}

	if options.Stderr == nil {
		options.Stderr = io.Discard
	}

	out, err := m.client.ContainerLogs(ctx, containerID, container.LogsOptions{
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
