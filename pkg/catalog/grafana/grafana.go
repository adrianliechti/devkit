package grafana

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
	return "grafana"
}

func (m *Manager) Category() catalog.Category {
	return catalog.PlatformCategory
}

func (m *Manager) DisplayName() string {
	return "Grafana"
}

func (m *Manager) Description() string {
	return "Grafana is a multi-platform open source analytics and interactive visualization web application."
}

const (
	DefaultShell = "/bin/bash"
)

func (m *Manager) New() (engine.Container, error) {
	image := "grafana/grafana"

	username := "admin"
	password := password.MustGenerate(10, 4, 0, false, false)

	return engine.Container{
		Image: image,

		Env: map[string]string{
			"GF_SECURITY_ADMIN_USER":     username,
			"GF_SECURITY_ADMIN_PASSWORD": password,
		},

		Ports: []*engine.ContainerPort{
			{
				Port:  3000,
				Proto: engine.ProtocolTCP,
			},
		},

		Mounts: []*engine.ContainerMount{
			{
				Path: "/var/lib/grafana",
			},
		},
	}, nil
}

func (m *Manager) Info(instance engine.Container) (map[string]string, error) {
	username := instance.Env["GF_SECURITY_ADMIN_USER"]
	password := instance.Env["GF_SECURITY_ADMIN_PASSWORD"]

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
		Port:  3000,
		Proto: engine.ProtocolTCP,
	}, nil
}
