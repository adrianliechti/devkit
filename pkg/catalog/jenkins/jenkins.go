package jenkins

import (
	"github.com/adrianliechti/devkit/pkg/catalog"
	"github.com/adrianliechti/devkit/pkg/engine"
	"github.com/sethvargo/go-password/password"
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
	return "jenkins"
}

func (m *Manager) Category() catalog.Category {
	return catalog.PlatformCategory
}

func (m *Manager) DisplayName() string {
	return "Jenkins"
}

func (m *Manager) Description() string {
	return "Jenkins is an open source automation server. It helps automate the parts of software development related to building, testing, and deploying, facilitating continuous integration and continuous delivery."
}

const (
	DefaultShell = "/bin/bash"
)

func (m *Manager) New() (engine.Container, error) {
	image := "ghcr.io/adrianliechti/loop-jenkins:dind"

	username := "admin"
	password := password.MustGenerate(10, 4, 0, false, false)

	return engine.Container{
		Image: image,

		Privileged: true,

		Env: map[string]string{
			"BASE_URL": "http://localhost:8080",

			"ADMIN_USERNAME": username,
			"ADMIN_PASSWORD": password,
		},

		Ports: []*engine.ContainerPort{
			{
				Port:  8080,
				Proto: engine.ProtocolTCP,
			},
		},

		Mounts: []*engine.ContainerMount{
			{
				Path: "/var/jenkins_home",
			},
		},
	}, nil
}

func (m *Manager) Info(instance engine.Container) (map[string]string, error) {
	username := instance.Env["ADMIN_USERNAME"]
	password := instance.Env["ADMIN_PASSWORD"]

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
		Port:  8080,
		Proto: engine.ProtocolTCP,
	}, nil
}
