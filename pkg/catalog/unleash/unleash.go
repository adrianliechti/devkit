package unleash

import (
	"github.com/adrianliechti/devkit/pkg/catalog"
	"github.com/adrianliechti/devkit/pkg/engine"
)

var (
	_ catalog.Manager         = &Manager{}
	_ catalog.Decorator       = &Manager{}
	_ catalog.ConsoleProvider = &Manager{}
)

type Manager struct {
}

func (m *Manager) Name() string {
	return "unleash"
}

func (m *Manager) Category() catalog.Category {
	return catalog.PlatformCategory
}

func (m *Manager) DisplayName() string {
	return "Unleash"
}

func (m *Manager) Description() string {
	return "Feature management lets you turn new features on/off in production with no need for redeployment. A software development best practice for releasing and validating new features."
}

func (m *Manager) New() (engine.Container, error) {
	image := "adrianliechti/loop-unleash"

	return engine.Container{
		Image: image,

		Env: map[string]string{},

		Ports: []*engine.ContainerPort{
			{
				Port:  3000,
				Proto: engine.ProtocolTCP,
			},
		},

		Mounts: []*engine.ContainerMount{
			{},
		},
	}, nil
}

func (m *Manager) Info(instance engine.Container) (map[string]string, error) {
	username := "admin"
	password := "unleash4all"
	token := "default:development.unleash-insecure-api-token"

	return map[string]string{
		"Username":  username,
		"Password":  password,
		"API Token": token,
	}, nil
}

func (m *Manager) ConsolePort(instance engine.Container) (*engine.ContainerPort, error) {
	return &engine.ContainerPort{
		Port:  4242,
		Proto: engine.ProtocolTCP,
	}, nil
}
