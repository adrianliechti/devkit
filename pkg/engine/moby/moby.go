package moby

import (
	"context"
	"time"

	"github.com/adrianliechti/devkit/pkg/engine"
	"github.com/adrianliechti/devkit/pkg/system"

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

// syncTerminalSize applies the current terminal size and keeps it in sync until ctx is done.
func syncTerminalSize(ctx context.Context, term any, resize func(width, height int) error) error {
	width, height, err := system.TerminalSize(term)

	if err != nil {
		return err
	}

	if err := resize(width, height); err != nil {
		return err
	}

	go func() {
		for ctx.Err() == nil {
			time.Sleep(200 * time.Millisecond)

			w, h, err := system.TerminalSize(term)

			if err != nil || (w == width && h == height) {
				continue
			}

			width, height = w, h
			resize(width, height)
		}
	}()

	return nil
}
