package activemq

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
	return "activemq"
}

func (m *Manager) Category() catalog.Category {
	return catalog.MessagingCategory
}

func (m *Manager) DisplayName() string {
	return "ActiveMQ"
}

func (m *Manager) Description() string {
	return "ActiveMQ is an open source message broker written in Java together with a full Java Message Service client."
}

const (
	DefaultShell = "/bin/bash"
)

func (m *Manager) New() (container.Container, error) {
	image := "rmohr/activemq:5.15.9"

	return container.Container{
		Image: image,

		Env: map[string]string{},

		Ports: []*container.ContainerPort{
			{
				Port:     61616,
				Protocol: container.ProtocolTCP,
			},
		},

		VolumeMounts: []*container.VolumeMount{
			{
				Path: "/opt/activemq/conf",
			},
			{
				Path: "/opt/activemq/data",
			},
		},
	}, nil
}

func (m *Manager) Info(instance container.Container) (map[string]string, error) {
	return map[string]string{
		"Username": "admin",
		"Password": "admin",
	}, nil
}

func (m *Manager) Shell(instance container.Container) (string, error) {
	return DefaultShell, nil
}

func (m *Manager) ConsolePort(instance container.Container) (*container.ContainerPort, error) {
	return &container.ContainerPort{
		Port:     8161,
		Protocol: container.ProtocolTCP,
	}, nil
}
