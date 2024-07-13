package azurite

import (
	"github.com/adrianliechti/devkit/pkg/catalog"
	"github.com/adrianliechti/devkit/pkg/engine"
)

var (
	_ catalog.Manager       = &Manager{}
	_ catalog.Decorator     = &Manager{}
	_ catalog.ShellProvider = &Manager{}
)

type Manager struct {
}

func (m *Manager) Name() string {
	return "azurite"
}

func (m *Manager) Category() catalog.Category {
	return catalog.StorageCategory
}

func (m *Manager) DisplayName() string {
	return "Azure Storage Azurite"
}

func (m *Manager) Description() string {
	return "Azurite is an open source Azure Storage API compatible server (emulator)."
}

const (
	DefaultShell = "/bin/ash"
)

func (m *Manager) New() (engine.Container, error) {
	image := "mcr.microsoft.com/azure-storage/azurite"

	return engine.Container{
		Image: image,

		Env: map[string]string{
			// AZURITE_ACCOUNTS="account1:key1:key2;account2:key1:key2"
		},

		Ports: []*engine.ContainerPort{
			{
				Name:  "blob",
				Port:  10000,
				Proto: engine.ProtocolTCP,
			},
			{
				Name:  "queue",
				Port:  10001,
				Proto: engine.ProtocolTCP,
			},
			{
				Name:  "table",
				Port:  10002,
				Proto: engine.ProtocolTCP,
			},
		},

		Mounts: []*engine.ContainerMount{
			{
				Path: "/data",
			},
		},
	}, nil
}

func (m *Manager) Info(instance engine.Container) (map[string]string, error) {
	return map[string]string{
		"Account":    "devstoreaccount1",
		"AccountKey": "Eby8vdM02xNOcqFlqUwJPLlmEtlCDXJ1OUzFT50uSRZ6IFsuFq2UVErCz4I6tq/K1SZFPTOtr/KBHBeksoGMGw==",
	}, nil
}

func (m *Manager) Shell(instance engine.Container) (string, error) {
	return DefaultShell, nil
}
