package ghost

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
	return "ghost"
}

func (m *Manager) Category() catalog.Category {
	return catalog.PlatformCategory
}

func (m *Manager) DisplayName() string {
	return "Ghost"
}

func (m *Manager) Description() string {
	return "Ghost is a powerful app for new-media creators to publish, share, and grow a business around their content."
}

const (
	DefaultShell = "/bin/bash"
)

func (m *Manager) New() (container.Container, error) {
	image := "ghost:5"

	return container.Container{
		Image: image,

		Env: map[string]string{
			"url": "http://localhost:2368",
		},

		Ports: []*container.ContainerPort{
			{
				Port:     2368,
				Protocol: container.ProtocolTCP,
			},
		},

		VolumeMounts: []*container.VolumeMount{
			{
				Path: "/var/lib/ghost/content",
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
		Port:     2368,
		Protocol: container.ProtocolTCP,
	}, nil
}
