package minio

import (
	"github.com/adrianliechti/devkit/pkg/container"
	"github.com/sethvargo/go-password/password"
)

const (
	DefaultShell = "/bin/bash"
)

func New() container.Container {
	image := "minio/minio"

	username := "root"
	password := password.MustGenerate(10, 4, 0, false, false)

	return container.Container{
		Image: image,

		Env: map[string]string{
			"MINIO_ROOT_USER":     username,
			"MINIO_ROOT_PASSWORD": password,
		},

		Args: []string{
			"server", "/data", "--console-address", ":9001",
		},

		Ports: []*container.ContainerPort{
			{
				Port:     9000,
				Protocol: container.ProtocolTCP,
			},
		},

		VolumeMounts: []*container.VolumeMount{
			{
				Path: "/data",
			},
		},
	}
}

func Info(container *container.Container) map[string]string {
	username := container.Env["MINIO_ROOT_USER"]
	password := container.Env["MINIO_ROOT_PASSWORD"]

	return map[string]string{
		"Username": username,
		"Password": password,
	}
}

func ConsolePort() container.ContainerPort {
	return container.ContainerPort{
		Port:     9001,
		Protocol: container.ProtocolTCP,
	}
}
