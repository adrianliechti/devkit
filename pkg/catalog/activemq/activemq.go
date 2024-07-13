package activemq

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

func (m *Manager) New() (engine.Container, error) {
	image := "apache/activemq-classic"

	return engine.Container{
		Image: image,

		Env: map[string]string{},

		Ports: []*engine.ContainerPort{
			{
				Port:  61616,
				Proto: engine.ProtocolTCP,
			},
		},

		Mounts: []*engine.ContainerMount{
			{
				Path: "/opt/activemq/conf",
			},
			{
				Path: "/opt/activemq/data",
			},
		},
	}, nil
}

func (m *Manager) Info(instance engine.Container) (map[string]string, error) {
	return map[string]string{
		"Username": "admin",
		"Password": "admin",
	}, nil
}

func (m *Manager) Shell(instance engine.Container) (string, error) {
	return DefaultShell, nil
}

func (m *Manager) ConsolePort(instance engine.Container) (*engine.ContainerPort, error) {
	return &engine.ContainerPort{
		Port:  8161,
		Proto: engine.ProtocolTCP,
	}, nil
}
