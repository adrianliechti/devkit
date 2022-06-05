package etcd

import (
	"github.com/adrianliechti/devkit/pkg/container"
	"github.com/sethvargo/go-password/password"
)

const (
	DefaultShell = "/bin/bash"
)

func New() container.Container {
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

		Ports: []container.ContainerPort{
			{
				Port:     2379,
				Protocol: container.ProtocolTCP,
			},
		},

		VolumeMounts: []container.VolumeMount{
			{
				Path: "/etcd-data",
			},
		},
	}
}

func Info(container *container.Container) map[string]string {
	return map[string]string{}
}
