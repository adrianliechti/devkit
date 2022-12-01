package keycloak

import (
	"github.com/adrianliechti/devkit/pkg/catalog"
	"github.com/adrianliechti/devkit/pkg/engine"
	"github.com/sethvargo/go-password/password"
)

var (
	_ catalog.Manager       = &Manager{}
	_ catalog.Decorator     = &Manager{}
	_ catalog.ShellProvider = &Manager{}
)

type Manager struct {
}

func (m *Manager) Name() string {
	return "keycloak"
}

func (m *Manager) Category() catalog.Category {
	return catalog.PlatformCategory
}

func (m *Manager) DisplayName() string {
	return "Keycloak"
}

func (m *Manager) Description() string {
	return "Keycloak is an Open Source Identity and Access Management."
}

const (
	DefaultShell = "/bin/bash"
)

func (m *Manager) New() (engine.Container, error) {
	image := "quay.io/keycloak/keycloak:20.0.1"

	user := "admin"
	password := password.MustGenerate(10, 4, 0, false, false)

	return engine.Container{
		Image: image,

		Env: map[string]string{
			"KEYCLOAK_ADMIN":          user,
			"KEYCLOAK_ADMIN_PASSWORD": password,
		},

		Args: []string{
			"start-dev",
			"--http-port", "8443",
		},

		Ports: []*engine.ContainerPort{
			{
				Port:  8443,
				Proto: engine.ProtocolTCP,
			},
		},
	}, nil
}

func (m *Manager) Info(instance engine.Container) (map[string]string, error) {
	user := instance.Env["KEYCLOAK_ADMIN"]
	password := instance.Env["KEYCLOAK_ADMIN_PASSWORD"]

	return map[string]string{
		"Username": user,
		"Password": password,
	}, nil
}

func (m *Manager) Shell(instance engine.Container) (string, error) {
	return DefaultShell, nil
}
