package kafka

import (
	"github.com/adrianliechti/devkit/catalog"
	"github.com/adrianliechti/devkit/pkg/engine"
)

var (
	_ catalog.Manager       = &Manager{}
	_ catalog.Decorator     = &Manager{}
	_ catalog.ShellProvider = &Manager{}
)

type Manager struct {
}

func (m *Manager) Name() string {
	return "kafka"
}

func (m *Manager) Category() catalog.Category {
	return catalog.MessagingCategory
}

func (m *Manager) DisplayName() string {
	return "Kafka"
}

func (m *Manager) Description() string {
	return "Apache Kafka is a distributed event store and stream-processing platform."
}

const (
	DefaultShell = "/bin/bash"
)

func (m *Manager) New() (engine.Container, error) {
	image := "apache/kafka:3.7.1"

	return engine.Container{
		Image: image,

		Ports: []engine.ContainerPort{
			{
				Port:  9092,
				Proto: engine.ProtocolTCP,
			},
		},

		Mounts: []engine.ContainerMount{
			{
				Path: "/tmp/kafka-logs",
			},
			{
				Path: "/tmp/kraft-combined-logs",
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
