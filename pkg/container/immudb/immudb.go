package immudb

import (
	"github.com/adrianliechti/devkit/pkg/container"
	"github.com/sethvargo/go-password/password"
)

const (
	DefaultShell = "/bin/bash"
)

func New() container.Container {
	image := "codenotary/immudb"

	password := password.MustGenerate(10, 4, 0, false, false)

	return container.Container{
		Image: image,

		Env: map[string]string{
			"IMMUDB_ADDRESS":        "0.0.0.0",
			"IMMUDB_ADMIN_PASSWORD": password,
		},

		Ports: []*container.ContainerPort{
			{
				Port:     3322,
				Protocol: container.ProtocolTCP,
			},
		},

		VolumeMounts: []*container.VolumeMount{
			{
				Path: "/var/lib/immudb",
			},
		},
	}
}

func Info(container *container.Container) map[string]string {
	username := "immudb"
	password := container.Env["IMMUDB_ADMIN_PASSWORD"]

	return map[string]string{
		"Username": username,
		"Password": password,
	}
}

func ConsolePort() container.ContainerPort {
	return container.ContainerPort{
		Port:     8080,
		Protocol: container.ProtocolTCP,
	}
}
