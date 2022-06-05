package azurite

import (
	"github.com/adrianliechti/devkit/pkg/catalog"
	"github.com/adrianliechti/devkit/pkg/container"
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

func (m *Manager) New() (container.Container, error) {
	image := "mcr.microsoft.com/azure-storage/azurite:3.17.1"

	return container.Container{
		Image: image,

		Env: map[string]string{
			// AZURITE_ACCOUNTS="account1:key1:key2;account2:key1:key2"
		},

		Ports: []*container.ContainerPort{
			{
				Name:     "blob",
				Port:     10000,
				Protocol: container.ProtocolTCP,
			},
			{
				Name:     "queue",
				Port:     10001,
				Protocol: container.ProtocolTCP,
			},
			{
				Name:     "table",
				Port:     10002,
				Protocol: container.ProtocolTCP,
			},
		},

		VolumeMounts: []*container.VolumeMount{
			{
				Path: "/data",
			},
		},
	}, nil
}

func (m *Manager) Info(instance container.Container) (map[string]string, error) {
	return map[string]string{
		"Account":    "devstoreaccount1",
		"AccountKey": "Eby8vdM02xNOcqFlqUwJPLlmEtlCDXJ1OUzFT50uSRZ6IFsuFq2UVErCz4I6tq/K1SZFPTOtr/KBHBeksoGMGw==",
	}, nil
}

func (m *Manager) Shell(instance container.Container) (string, error) {
	return DefaultShell, nil
}
