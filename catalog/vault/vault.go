package vault

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

const (
	DefaultShell = "/bin/ash"
)

func (m *Manager) Name() string {
	return "vault"
}

func (m *Manager) Category() catalog.Category {
	return catalog.StorageCategory
}

func (m *Manager) DisplayName() string {
	return "HashiCorp Vault"
}

func (m *Manager) Description() string {
	return "Vault secures, stores, and tightly controls access to tokens, passwords, certificates, API keys, and other secrets in modern computing."
}

func (m *Manager) New() (engine.Container, error) {
	image := "vault:latest"

	token := password.MustGenerate(10, 4, 0, false, false)

	return engine.Container{
		Image: image,

		Env: map[string]string{
			"VAULT_DEV_ROOT_TOKEN_ID": token,
		},

		Ports: []*engine.ContainerPort{
			{
				Port:  8200,
				Proto: engine.ProtocolTCP,
			},
		},
	}, nil
}

func (m *Manager) Info(instance engine.Container) (map[string]string, error) {
	token := instance.Env["VAULT_DEV_ROOT_TOKEN_ID"]

	return map[string]string{
		"Token": token,
	}, nil
}

func (m *Manager) Shell(instance engine.Container) (string, error) {
	return DefaultShell, nil
}

func (m *Manager) ConsolePort(instance engine.Container) (*engine.ContainerPort, error) {
	return &engine.ContainerPort{
		Port:  8200,
		Proto: engine.ProtocolTCP,
	}, nil
}
