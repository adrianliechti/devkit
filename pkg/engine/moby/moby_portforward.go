package moby

import (
	"context"
	"errors"
	"fmt"

	"github.com/adrianliechti/devkit/pkg/engine"
)

func (m *Moby) PortForward(ctx context.Context, containerID, address string, ports map[int]int, readyChan chan struct{}) error {
	if address == "" {
		address = "127.0.0.1"
	}

	info, err := m.Inspect(ctx, containerID)

	if err != nil {
		return err
	}

	if info.IPAddress == nil {
		return errors.New("invalid container ip")
	}

	target := info.IPAddress.String()

	for s, t := range ports {
		container := engine.Container{
			Image: "alpine/socat",

			Args: []string{
				fmt.Sprintf("TCP4-LISTEN:%d,fork,reuseaddr", t),
				fmt.Sprintf("TCP4:%s:%d", target, t),
			},

			Ports: []engine.ContainerPort{
				{
					HostPort: s,
					HostIP:   address,

					// socat listens on the target port inside the sidecar, so the
					// published container port must match it, not the host port.
					Port: t,
				},
			},
		}

		id, err := m.Create(ctx, container, engine.CreateOptions{})

		if err != nil {
			return err
		}

		defer m.Delete(context.Background(), id, engine.DeleteOptions{})
	}

	if readyChan != nil {
		close(readyChan)
	}

	<-ctx.Done()

	return nil
}
