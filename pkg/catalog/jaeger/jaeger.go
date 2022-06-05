package jaeger

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
	return "jaeger"
}

func (m *Manager) Category() catalog.Category {
	return catalog.PlatformCategory
}

func (m *Manager) DisplayName() string {
	return "Jaeger"
}

func (m *Manager) Description() string {
	return "Monitor and troubleshoot transactions in complex distributed systems"
}

const (
	DefaultShell = "/bin/bash"
)

func (m *Manager) New() (container.Container, error) {
	image := "jaegertracing/all-in-one:1.35"

	return container.Container{
		Image: image,

		Env: map[string]string{
			"COLLECTOR_ZIPKIN_HOST_PORT": ":9411",
			"COLLECTOR_OTLP_ENABLED":     "true",
		},

		Ports: []*container.ContainerPort{
			{
				Name:     "jaeger-thrift-compact",
				Port:     6831,
				Protocol: container.ProtocolUDP,
			},
			{
				Name:     "jaeger-thrift-binary",
				Port:     6832,
				Protocol: container.ProtocolUDP,
			},
			{
				Name:     "otlp-grpc",
				Port:     4317,
				Protocol: container.ProtocolTCP,
			},
			{
				Name:     "otlp-http",
				Port:     4318,
				Protocol: container.ProtocolTCP,
			},
			{
				Name:     "zipkin",
				Port:     9411,
				Protocol: container.ProtocolTCP,
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
		Port:     16686,
		Protocol: container.ProtocolTCP,
	}, nil
}
