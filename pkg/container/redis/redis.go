package redis

import (
	"github.com/adrianliechti/devkit/pkg/container"
	"github.com/sethvargo/go-password/password"
)

const (
	DefaultShell = "/bin/bash"
)

func New() container.Container {
	image := "redis:6-bullseye"

	password := password.MustGenerate(10, 4, 0, false, false)

	return container.Container{
		Image: image,

		Env: map[string]string{
			"REDIS_PASSWORD": password,
		},

		Ports: []*container.ContainerPort{
			{
				Port:     6379,
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
	password := container.Env["REDIS_PASSWORD"]

	return map[string]string{
		"Password": password,
	}
}

func ClientCmd() (string, []string) {
	return DefaultShell, []string{
		"-c",
		"redis-cli",
	}
}
