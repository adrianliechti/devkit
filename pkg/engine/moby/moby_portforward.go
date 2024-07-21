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

	info, err := m.client.ContainerInspect(ctx, containerID)

	if err != nil {
		return err
	}

	var target string

	if info.NetworkSettings != nil {
		target = info.NetworkSettings.IPAddress
	}

	if target == "" {
		return errors.New("invalid container ip")
	}

	for s, t := range ports {
		container := engine.Container{
			Image: "alpine/socat",

			Args: []string{
				fmt.Sprintf("TCP4-LISTEN:%d,fork,reuseaddr", s),
				fmt.Sprintf("TCP4:%s:%d", target, t),
			},

			Ports: []engine.ContainerPort{
				{
					HostPort: s,
					HostIP:   address,

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
