package mailpit

import (
	"strings"

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
	return "mailpit"
}

func (m *Manager) Category() catalog.Category {
	return catalog.PlatformCategory
}

func (m *Manager) DisplayName() string {
	return "Mailpit"
}

func (m *Manager) Description() string {
	return "Mailpit is a small, fast, low memory, zero-dependency, multi-platform email testing tool & API for developers."
}

const (
	DefaultShell = "/bin/ash"
)

func (m *Manager) New() (engine.Container, error) {
	image := "axllent/mailpit"

	username := "mailpit"
	password := password.MustGenerate(10, 4, 0, false, false)

	return engine.Container{
		Image: image,

		Privileged: true,

		Env: map[string]string{
			"MP_UI_AUTH":   username + ":" + password,
			"MP_SMTP_AUTH": username + ":" + password,

			"MP_SMTP_AUTH_ALLOW_INSECURE": "true",
		},

		Ports: []engine.ContainerPort{
			{
				Name:  "smtp",
				Port:  1025,
				Proto: engine.ProtocolTCP,
			},
			{
				Name:  "http",
				Port:  8025,
				Proto: engine.ProtocolTCP,
			},
		},

		Mounts: []engine.ContainerMount{
			{
				Path: "/data",
			},
		},
	}, nil
}

func (m *Manager) Info(instance engine.Container) (map[string]string, error) {
	auth := instance.Env["MP_UI_AUTH"]
	parts := strings.Split(auth, ":")

	return map[string]string{
		"Username": parts[0],
		"Password": parts[1],
	}, nil
}

func (m *Manager) Shell(instance engine.Container) (string, error) {
	return DefaultShell, nil
}

func (m *Manager) ConsolePort(instance engine.Container) (*engine.ContainerPort, error) {
	return &engine.ContainerPort{
		Port:  8025,
		Proto: engine.ProtocolTCP,
	}, nil
}
