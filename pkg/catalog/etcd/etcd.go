package etcd

import (
	"github.com/adrianliechti/devkit/pkg/catalog"
	"github.com/adrianliechti/devkit/pkg/container"
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
	return "etcd"
}

func (m *Manager) Category() catalog.Category {
	return catalog.DatabaseCategory
}

func (m *Manager) DisplayName() string {
	return "etcd"
}

func (m *Manager) Description() string {
	return "etcd is a distributed reliable key-value store for the most critical data of a distributed system."
}

const (
	DefaultShell = "/bin/bash"
)

func (m *Manager) New() (container.Container, error) {
	image := "gcr.io/etcd-development/etcd:v3.3.8"

	token := password.MustGenerate(10, 4, 0, false, false)

	return container.Container{
		Image: image,

		Env: map[string]string{
			"ETCD_NAME": "etcd0",

			"ETCD_DATA_DIR": "/etcd-data",

			"ETCD_LISTEN_PEER_URLS":   "http://0.0.0.0:2380",
			"ETCD_LISTEN_CLIENT_URLS": "http://0.0.0.0:2379",

			"ETCD_INITIAL_CLUSTER":             "etcd0=http://127.0.0.1:2380",
			"ETCD_INITIAL_CLUSTER_STATE":       "new",
			"ETCD_INITIAL_CLUSTER_TOKEN":       token,
			"ETCD_INITIAL_ADVERTISE_PEER_URLS": "http://127.0.0.1:2380",

			"ETCD_ADVERTISE_CLIENT_URLS": "http://127.0.0.1:2379",
		},

		Ports: []*container.ContainerPort{
			{
				Port:     2379,
				Protocol: container.ProtocolTCP,
			},
		},

		VolumeMounts: []*container.VolumeMount{
			{
				Path: "/etcd-data",
			},
		},
	}, nil
}

func (m *Manager) Info(instance container.Container) (map[string]string, error) {
	return map[string]string{}, nil
}

func (m *Manager) Shell(instance container.Container) (string, error) {
	return DefaultShell, nil
}
