package mailtrap

import (
	"github.com/adrianliechti/devkit/pkg/catalog"
	"github.com/adrianliechti/devkit/pkg/container"
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
	return "mailtrap"
}

func (m *Manager) Category() catalog.Category {
	return catalog.PlatformCategory
}

func (m *Manager) DisplayName() string {
	return "MailTrap"
}

func (m *Manager) Description() string {
	return ""
}

const (
	DefaultShell = "/bin/bash"
)

func (m *Manager) New() (container.Container, error) {
	image := "adrianliechti/loop-mailtrap"

	username := "admin"
	password := password.MustGenerate(10, 4, 0, false, false)

	return container.Container{
		Image: image,

		PlatformContext: &container.PlatformContext{
			Platform: "linux/amd64",
		},

		Env: map[string]string{
			"USERNAME": username,
			"PASSWORD": password,
		},

		Ports: []*container.ContainerPort{
			{
				Port:     25,
				Protocol: container.ProtocolTCP,
			},
		},
	}, nil
}

func (m *Manager) Info(instance container.Container) (map[string]string, error) {
	username := instance.Env["USERNAME"]
	password := instance.Env["PASSWORD"]

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
		Port:     80,
		Protocol: container.ProtocolTCP,
	}, nil
}
