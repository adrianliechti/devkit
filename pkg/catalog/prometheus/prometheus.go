package prometheus

import (
	"github.com/adrianliechti/devkit/pkg/catalog"
	"github.com/adrianliechti/devkit/pkg/container"
)

var (
	_ catalog.Manager         = &Manager{}
	_ catalog.Decorator       = &Manager{}
	_ catalog.ShellProvider   = &Manager{}
	_ catalog.ConsoleProvider = &Manager{}
)

type Manager struct {
}

func (m *Manager) Name() string {
	return "prometheus"
}

func (m *Manager) Category() catalog.Category {
	return catalog.DatabaseCategory
}

func (m *Manager) DisplayName() string {
	return "Prometheus"
}

func (m *Manager) Description() string {
	return "Prometheus is a free software application used for event monitoring and alerting."
}

const (
	DefaultShell = "/bin/bash"
)

func (m *Manager) New() (container.Container, error) {
	image := "prom/prometheus:v2.36.0"

	return container.Container{
		Image: image,

		Env: map[string]string{},

		Ports: []*container.ContainerPort{
			{
				Port:     9090,
				Protocol: container.ProtocolTCP,
			},
		},

		VolumeMounts: []*container.VolumeMount{
			{
				Path: "/prometheus",
			},
		},
	}, nil
}

func (m *Manager) Info(instance container.Container) (map[string]string, error) {
	return map[string]string{}, nil
}

func (m *Manager) Shell(instance container.Container) (string, error) {
	return DefaultShell, nil
}

func (m *Manager) ConsolePort(instance container.Container) (*container.ContainerPort, error) {
	return &container.ContainerPort{
		Port:     9090,
		Protocol: container.ProtocolTCP,
	}, nil
}
