package cratedb

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
	return "cratedb"
}

func (m *Manager) Category() catalog.Category {
	return catalog.DatabaseCategory
}

func (m *Manager) DisplayName() string {
	return "CrateDB"
}

func (m *Manager) Description() string {
	return "Enabling Data Insights at Scale."
}

const (
	DefaultShell = "/bin/bash"
)

func (m *Manager) New() (engine.Container, error) {
	image := "crate/crate:latest"

	return engine.Container{
		Image: image,

		Env: map[string]string{},

		Ports: []*engine.ContainerPort{
			{
				Port:  4200,
				Proto: engine.ProtocolTCP,
			},
		},

		Args: []string{
			"-Cdiscovery.type=single-node",
		},

		Mounts: []*engine.ContainerMount{
			{
				Path: "/cratedb/data",
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
		Port:  4200,
		Proto: engine.ProtocolTCP,
	}, nil
}
