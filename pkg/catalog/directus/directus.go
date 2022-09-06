package directus

import (
	"github.com/adrianliechti/devkit/pkg/catalog"
	"github.com/adrianliechti/devkit/pkg/engine"

	"github.com/google/uuid"
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
	return "directus"
}

func (m *Manager) Category() catalog.Category {
	return catalog.PlatformCategory
}

func (m *Manager) DisplayName() string {
	return "Directus"
}

func (m *Manager) Description() string {
	return "Directus is an Open Data Platform built to democratize the database."
}

const (
	DefaultShell = "/bin/ash"
)

func (m *Manager) New() (engine.Container, error) {
	image := "directus/directus:9"

	username := "admin@example.com"
	password := password.MustGenerate(10, 4, 0, false, false)

	return engine.Container{
		Image: image,

		Env: map[string]string{
			"KEY":    uuid.New().String(),
			"SECRET": uuid.New().String(),

			"ADMIN_EMAIL":    username,
			"ADMIN_PASSWORD": password,
		},

		Ports: []*engine.ContainerPort{
			{
				Port:  8055,
				Proto: engine.ProtocolTCP,
			},
		},

		Mounts: []*engine.ContainerMount{
			{
				Path: "/directus/uploads",
			},
			{
				Path: "/directus/database",
			},
			{
				Path: "/directus/extensions",
			},
		},
	}, nil
}

func (m *Manager) Info(instance engine.Container) (map[string]string, error) {
	key := instance.Env["KEY"]
	secret := instance.Env["SECRET"]

	username := instance.Env["ADMIN_EMAIL"]
	password := instance.Env["ADMIN_PASSWORD"]

	return map[string]string{
		"Key":    key,
		"Secret": secret,

		"Username": username,
		"Password": password,
	}, nil
}

func (m *Manager) Shell(instance engine.Container) (string, error) {
	return DefaultShell, nil
}

func (m *Manager) ConsolePort(instance engine.Container) (*engine.ContainerPort, error) {
	return &engine.ContainerPort{
		Port:  8055,
		Proto: engine.ProtocolTCP,
	}, nil
}
