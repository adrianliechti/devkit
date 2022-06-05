package elasticsearch

import (
	"github.com/adrianliechti/devkit/pkg/container"
	"github.com/sethvargo/go-password/password"
)

const (
	DefaultShell = "/bin/bash"
)

func New() container.Container {
	image := "docker.elastic.co/elasticsearch/elasticsearch:8.2.0"

	password := password.MustGenerate(10, 4, 0, false, false)

	return container.Container{
		Image: image,

		Env: map[string]string{
			"node.name": "es",

			"cluster.name":   "default",
			"discovery.type": "single-node",

			"xpack.security.enabled": "true",

			"ELASTIC_PASSWORD": password,
		},

		Ports: []*container.ContainerPort{
			{
				Port:     9200,
				Protocol: container.ProtocolTCP,
			},
		},

		VolumeMounts: []*container.VolumeMount{
			{
				Path: "/usr/share/elasticsearch/data",
			},
		},
	}
}

func Info(container *container.Container) map[string]string {
	username := "elastic"
	password := container.Env["ELASTIC_PASSWORD"]

	return map[string]string{
		"Username": username,
		"Password": password,
	}
}
