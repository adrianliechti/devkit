package pocketbase

import (
	"github.com/adrianliechti/devkit/catalog"
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
	return "pocketbase"
}

func (m *Manager) Category() catalog.Category {
	return catalog.StorageCategory
}

func (m *Manager) DisplayName() string {
	return "PocketBase"
}

func (m *Manager) Description() string {
	return "Open Source backend for your next SaaS and Mobile app in 1 file."
}

const (
	DefaultShell = "/bin/ash"
)

func (m *Manager) New() (engine.Container, error) {
	image := "ghcr.io/adrianliechti/loop-pocketbase"

	username := "admin@pocketbase.local"
	password := password.MustGenerate(10, 4, 0, false, false)

	return engine.Container{
		Image: image,

		Env: map[string]string{
			"APP_URL":        "http://localhost:8090",
			"APP_NAME":       "PocketBase",
			"ADMIN_USERNAME": username,
			"ADMIN_PASSWORD": password,
		},

		Ports: []*engine.ContainerPort{
			{
				Port:  8090,
				Proto: engine.ProtocolTCP,
			},
		},

		Mounts: []*engine.ContainerMount{
			{
				Path: "/pb_data",
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
		Port:  8090,
		Proto: engine.ProtocolTCP,
	}, nil
}
