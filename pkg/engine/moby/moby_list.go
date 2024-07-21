package moby

import (
	"context"
	"fmt"

	"github.com/adrianliechti/devkit/pkg/engine"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
)

func (m *Moby) List(ctx context.Context, options engine.ListOptions) ([]engine.Container, error) {
	opts := container.ListOptions{
		All:     options.All,
		Filters: filters.NewArgs(),
	}

	for k, v := range options.LabelSelector {
		opts.Filters.Add("label", fmt.Sprintf("%s=%s", k, v))
	}

	list, err := m.client.ContainerList(ctx, opts)

	if err != nil {
		return nil, err
	}

	containers := []engine.Container{}

	for _, i := range list {
		container, err := m.Inspect(ctx, i.ID)

		if err != nil {
			return nil, err
		}

		containers = append(containers, container)
	}

	return containers, nil
}
