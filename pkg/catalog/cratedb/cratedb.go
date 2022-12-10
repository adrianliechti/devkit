package cratedb

import (
	"runtime"

	"github.com/adrianliechti/devkit/pkg/catalog"
	"github.com/adrianliechti/devkit/pkg/engine"
)

var (
	_ catalog.Manager         = &Manager{}
	_ catalog.Decorator       = &Manager{}
	_ catalog.ShellProvider   = &Manager{}
	_ catalog.ClientProvider  = &Manager{}
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

	if runtime.GOARCH == "arm64" {
		image = "arm64v8/crate:latest"
	}

	return engine.Container{
		Image: image,

		Env: map[string]string{},

		Ports: []*engine.ContainerPort{
			{
				Name:  "http",
				Port:  4200,
				Proto: engine.ProtocolTCP,
			},
			{
				Name:  "psql",
				Port:  5432,
				Proto: engine.ProtocolTCP,
			},
		},

		Args: []string{
			"-Cdiscovery.type=single-node",
		},

		Mounts: []*engine.ContainerMount{
			{
				Path: "/data",
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

func (m *Manager) Client(instance engine.Container) (string, []string, error) {
	return "", []string{
		DefaultShell,
		"-c",
		"su - crate -c 'crash --hosts localhost'",
	}, nil
}

func (m *Manager) ConsolePort(instance engine.Container) (*engine.ContainerPort, error) {
	return &engine.ContainerPort{
		Port:  4200,
		Proto: engine.ProtocolTCP,
	}, nil
}
