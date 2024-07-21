package nexus

import (
	"runtime"

	"github.com/adrianliechti/devkit/catalog"
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
	return "nexus"
}

func (m *Manager) Category() catalog.Category {
	return catalog.StorageCategory
}

func (m *Manager) DisplayName() string {
	return "Nexus Repository Manager"
}

func (m *Manager) Description() string {
	return "Nexus Repository OSS is an open source repository that supports many artifact formats, including Docker, Java™, and npm."
}

const (
	DefaultShell = "/bin/bash"
)

func (m *Manager) New() (engine.Container, error) {
	image := "sonatype/nexus3:latest"

	if runtime.GOARCH == "arm64" {
		image = "klo2k/nexus3:latest"
	}

	return engine.Container{
		Image: image,

		Env: map[string]string{},

		Ports: []engine.ContainerPort{
			{
				Port:  8081,
				Proto: engine.ProtocolTCP,
			},
		},

		Mounts: []engine.ContainerMount{
			{
				Path: "/nexus-data",
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
		Port:  8081,
		Proto: engine.ProtocolTCP,
	}, nil
}
