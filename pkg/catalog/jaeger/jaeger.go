package jaeger

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

func (m *Manager) New() (engine.Container, error) {
	image := "jaegertracing/all-in-one:1.35"

	return engine.Container{
		Image: image,

		Env: map[string]string{
			"COLLECTOR_ZIPKIN_HOST_PORT": ":9411",
			"COLLECTOR_OTLP_ENABLED":     "true",
		},

		Ports: []*engine.ContainerPort{
			{
				Name:  "jaeger-thrift-compact",
				Port:  6831,
				Proto: engine.ProtocolUDP,
			},
			{
				Name:  "jaeger-thrift-binary",
				Port:  6832,
				Proto: engine.ProtocolUDP,
			},
			{
				Name:  "otlp-grpc",
				Port:  4317,
				Proto: engine.ProtocolTCP,
			},
			{
				Name:  "otlp-http",
				Port:  4318,
				Proto: engine.ProtocolTCP,
			},
			{
				Name:  "zipkin",
				Port:  9411,
				Proto: engine.ProtocolTCP,
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
		Port:  16686,
		Proto: engine.ProtocolTCP,
	}, nil
}
