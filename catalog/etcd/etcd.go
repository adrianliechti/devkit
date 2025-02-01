package etcd

import (
	"fmt"

	"github.com/adrianliechti/devkit/catalog"
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

func (m *Manager) New() (engine.Container, error) {
	image := "gcr.io/etcd-development/etcd:v3.5.18"

	token := password.MustGenerate(10, 4, 0, false, false)

	return engine.Container{
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

		Ports: []engine.ContainerPort{
			{
				Port:  2379,
				Proto: engine.ProtocolTCP,
			},
		},

		Mounts: []engine.ContainerMount{
			{
				Path: "/etcd-data",
			},
		},
	}, nil
}

func (m *Manager) Info(instance engine.Container) (map[string]string, error) {
	var uri string

	for _, p := range instance.Ports {
		if p.HostPort == 0 || p.Port != 2379 {
			continue
		}

		uri = fmt.Sprintf("http://localhost:%d", p.HostPort)
	}

	return map[string]string{
		"URI": uri,
	}, nil
}

func (m *Manager) Shell(instance engine.Container) (string, error) {
	return DefaultShell, nil
}
