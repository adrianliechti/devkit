package rabbitmq

import (
	"github.com/adrianliechti/devkit/pkg/container"
	"github.com/sethvargo/go-password/password"
)

const (
	DefaultShell = "/bin/bash"
)

func New() container.Container {
	image := "rabbitmq:3-management"

	username := "admin"
	password := password.MustGenerate(10, 4, 0, false, false)

	return container.Container{
		Image: image,

		Env: map[string]string{
			"RABBITMQ_DEFAULT_USER": username,
			"RABBITMQ_DEFAULT_PASS": password,
		},

		Ports: []*container.ContainerPort{
			{
				Port:     5672,
				Protocol: container.ProtocolTCP,
			},
		},

		VolumeMounts: []*container.VolumeMount{
			{
				Path: "/var/lib/rabbitmq",
			},
		},
	}
}

func Info(container *container.Container) map[string]string {
	username := container.Env["RABBITMQ_DEFAULT_USER"]
	password := container.Env["RABBITMQ_DEFAULT_PASS"]

	return map[string]string{
		"Username": username,
		"Password": password,
	}
}

func ConsolePort() container.ContainerPort {
	return container.ContainerPort{
		Port:     15672,
		Protocol: container.ProtocolTCP,
	}
}
