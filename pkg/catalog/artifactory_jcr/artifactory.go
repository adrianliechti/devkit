package artifactory_jcr

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
	return "artifactory-jcr"
}

func (m *Manager) Category() catalog.Category {
	return catalog.StorageCategory
}

func (m *Manager) DisplayName() string {
	return "JFrog Container Registry"
}

func (m *Manager) Description() string {
	return "Artifactory is the single solution for managing all the artifacts, binaries, files and containers throughout your software supply chain."
}

const (
	DefaultShell = "/bin/bash"
)

func (m *Manager) New() (engine.Container, error) {
	image := "releases-docker.jfrog.io/jfrog/artifactory-jcr:latest"

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
