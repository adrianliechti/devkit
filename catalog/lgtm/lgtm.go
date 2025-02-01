package lgtm

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
	return "lgtm"
}

func (m *Manager) Category() catalog.Category {
	return catalog.PlatformCategory
}

func (m *Manager) DisplayName() string {
	return "LGTM Stack"
}

func (m *Manager) Description() string {
	return "A comprehensive set of open-source tools designed for monitoring, observability, and visualization."
}

const (
	DefaultShell = "/bin/bash"
)

func (m *Manager) New() (engine.Container, error) {
	image := "grafana/otel-lgtm"

	return engine.Container{
		Image: image,

		Env: map[string]string{},

		Ports: []engine.ContainerPort{
			{
				Port:  3000,
				Proto: engine.ProtocolTCP,
			},
			{
				Name: "otlp-grpc",

				Port:  4317,
				Proto: engine.ProtocolTCP,
			},
			{
				Name: "otlp-http",

				Port:  4318,
				Proto: engine.ProtocolTCP,
			},
		},

		Mounts: []engine.ContainerMount{
			{
				Path: "/data",
			},
		},
	}, nil
}

func (m *Manager) Info(instance engine.Container) (map[string]string, error) {
	username := "admin"
	password := "admin"

	return map[string]string{
		"Username": username,
		"Password": password,
	}, nil
}

func (m *Manager) Shell(instance engine.Container) (string, error) {
	return DefaultShell, nil
}
