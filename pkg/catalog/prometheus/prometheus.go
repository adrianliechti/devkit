package prometheus

import (
	"github.com/adrianliechti/devkit/pkg/catalog"
	"github.com/adrianliechti/devkit/pkg/engine"
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

func (m *Manager) New() (engine.Container, error) {
	image := "prom/prometheus:v2.36.0"

	return engine.Container{
		Image: image,

		Env: map[string]string{},

		Ports: []*engine.ContainerPort{
			{
				Port:  9090,
				Proto: engine.ProtocolTCP,
			},
		},

		Mounts: []*engine.ContainerMount{
			{
				Path: "/prometheus",
			},
		},
	}, nil
}

func (m *Manager) Info(instance engine.Container) (map[string]string, error) {
	return map[string]string{}, nil
}

func (m *Manager) Shell(instance engine.Container) (string, error) {
	return DefaultShell, nil
}

func (m *Manager) ConsolePort(instance engine.Container) (*engine.ContainerPort, error) {
	return &engine.ContainerPort{
		Port:  9090,
		Proto: engine.ProtocolTCP,
	}, nil
}
