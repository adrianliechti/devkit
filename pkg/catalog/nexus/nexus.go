package nexus

import (
	"runtime"

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

func (m *Manager) New() (container.Container, error) {
	image := "sonatype/nexus3"

	if runtime.GOARCH == "arm64" {
		image = "klo2k/nexus3"
	}

	return container.Container{
		Image: image,

		Env: map[string]string{},

		Ports: []*container.ContainerPort{
			{
				Port:     8081,
				Protocol: container.ProtocolTCP,
			},
		},

		VolumeMounts: []*container.VolumeMount{
			{
				Path: "/nexus-data",
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
		Port:     8081,
		Protocol: container.ProtocolTCP,
	}, nil
}
