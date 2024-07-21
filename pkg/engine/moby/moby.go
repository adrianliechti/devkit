package moby

import (
	"github.com/adrianliechti/devkit/pkg/engine"

	"github.com/docker/docker/client"
)

var (
	_ engine.Client = &Moby{}
)

type Moby struct {
	client *client.Client
}

func New() (*Moby, error) {
	cli, err := client.NewClientWithOpts(
		client.FromEnv,
		client.WithAPIVersionNegotiation(),
	)

	if err != nil {
		return nil, err
	}

	return &Moby{
		client: cli,
	}, nil
}
