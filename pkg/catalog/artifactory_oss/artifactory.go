package artifactory_oss

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
	return "artifactor-oss"
}

func (m *Manager) Category() catalog.Category {
	return catalog.StorageCategory
}

func (m *Manager) DisplayName() string {
	return "JFrog Artifactory OSS"
}

func (m *Manager) Description() string {
	return "The world's most advanced, powerful, Docker and Helm registry. Power your world of Docker without limits."
}

const (
	DefaultShell = "/bin/bash"
)

func (m *Manager) New() (engine.Container, error) {
	image := "releases-docker.jfrog.io/jfrog/artifactory-oss:latest"

	return engine.Container{
		Image: image,

		Env: map[string]string{},

		Ports: []*engine.ContainerPort{
			{
				Name:  "api",
				Port:  8081,
				Proto: engine.ProtocolTCP,
			},
			{
				Port:  8082,
				Proto: engine.ProtocolTCP,
			},
		},

		Mounts: []*engine.ContainerMount{
			{
				Path: "/var/opt/jfrog/artifactory",
			},
		},
	}, nil
}

func (m *Manager) Info(instance engine.Container) (map[string]string, error) {
	username := "admin"
	password := "password"

	return map[string]string{
		"Username": username,
		"Password": password,
	}, nil
}

func (m *Manager) Shell(instance engine.Container) (string, error) {
	return DefaultShell, nil
}

func (m *Manager) ConsolePort(instance engine.Container) (*engine.ContainerPort, error) {
	return &engine.ContainerPort{
		Port:  8082,
		Proto: engine.ProtocolTCP,
	}, nil
}
