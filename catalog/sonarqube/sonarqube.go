package sonarqube

import (
	"github.com/adrianliechti/devkit/catalog"
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
	return "sonarqube"
}

func (m *Manager) Category() catalog.Category {
	return catalog.PlatformCategory
}

func (m *Manager) DisplayName() string {
	return "SonarQube"
}

func (m *Manager) Description() string {
	return "SonarQube empowers all developers to write cleaner and safer code."
}

const (
	DefaultShell = "/bin/bash"
)

func (m *Manager) New() (engine.Container, error) {
	image := "sonarqube:10-community"

	return engine.Container{
		Image: image,

		MaxFiles:     131072,
		MaxProcesses: 8192,

		Env: map[string]string{
			"SONAR_ES_BOOTSTRAP_CHECKS_DISABLE": "true",
		},

		Ports: []*engine.ContainerPort{
			{
				Port:  9000,
				Proto: engine.ProtocolTCP,
			},
		},

		Mounts: []*engine.ContainerMount{
			{
				Path: "/opt/sonarqube/data",
			},
			{
				Path: "/opt/sonarqube/logs",
			},
			{
				Path: "/opt/sonarqube/extensions",
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

func (m *Manager) ConsolePort(instance engine.Container) (*engine.ContainerPort, error) {
	return &engine.ContainerPort{
		Port:  9000,
		Proto: engine.ProtocolTCP,
	}, nil
}
