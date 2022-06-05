package sonarqube

import (
	"runtime"

	"github.com/adrianliechti/devkit/pkg/catalog"
	"github.com/adrianliechti/devkit/pkg/container"
	"github.com/adrianliechti/devkit/pkg/to"
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

func (m *Manager) New() (container.Container, error) {
	image := "sonarqube:9-community"

	if runtime.GOARCH == "arm64" {
		image = "mwizner/sonarqube:9.4.0-community"
	}

	// MaxNoProcs: 8192,
	// MaxNoFiles: 131072,

	return container.Container{
		Image: image,

		Env: map[string]string{
			"SONAR_ES_BOOTSTRAP_CHECKS_DISABLE": "true",
			"SONAR_SEARCH_JAVAADDITIONALOPTS":   "-Dbootstrap.system_call_filter=false",
		},

		Ports: []*container.ContainerPort{
			{
				Port:     9000,
				Protocol: container.ProtocolTCP,
			},
		},

		VolumeMounts: []*container.VolumeMount{
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

		PlatformContext: &container.PlatformContext{
			MaxNoProcs: to.IntPtr(8192),
			MaxNoFiles: to.IntPtr(131072),
		},
	}, nil
}

func (m *Manager) Info(instance container.Container) (map[string]string, error) {
	username := "admin"
	password := "admin"

	return map[string]string{
		"Username": username,
		"Password": password,
	}, nil
}

func (m *Manager) Shell(instance container.Container) (string, error) {
	return DefaultShell, nil
}

func (m *Manager) ConsolePort(instance container.Container) (*container.ContainerPort, error) {
	return &container.ContainerPort{
		Port:     9000,
		Protocol: container.ProtocolTCP,
	}, nil
}
