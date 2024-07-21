package moby

import (
	"context"

	"github.com/adrianliechti/devkit/pkg/engine"
	"github.com/docker/docker/api/types/container"
)

func (m *Moby) Delete(ctx context.Context, containerID string, options engine.DeleteOptions) error {
	m.client.ContainerStop(ctx, containerID, container.StopOptions{})

	return m.client.ContainerRemove(ctx, containerID, container.RemoveOptions{
		Force: true,

		RemoveVolumes: true,
	})
}
