package cassandra

import (
	"github.com/adrianliechti/devkit/pkg/container"
)

const (
	DefaultShell = "/bin/bash"
)

func New() container.Container {
	image := "cassandra:4"

	return container.Container{
		Image: image,

		Env: map[string]string{},

		Ports: []*container.ContainerPort{
			{
				Port:     9042,
				Protocol: container.ProtocolTCP,
			},
		},

		VolumeMounts: []*container.VolumeMount{
			{
				Path: "/var/lib/cassandra",
			},
		},
	}
}

func Info(container *container.Container) map[string]string {
	return map[string]string{}
}

func ClientCmd() (string, []string) {
	return DefaultShell, []string{
		"-c",
		"/opt/cassandra/bin/cqlsh",
	}
}
