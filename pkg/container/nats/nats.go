package nats

import (
	"github.com/adrianliechti/devkit/pkg/container"
	"github.com/sethvargo/go-password/password"
)

func New() container.Container {
	image := "nats:2-linux"

	username := "admin"
	password := password.MustGenerate(10, 4, 0, false, false)

	return container.Container{
		Image: image,

		Env: map[string]string{
			"USERNAME": username,
			"PASSWORD": password,
		},

		Args: []string{
			"-js",
			"--name", "default",
			"--cluster_name", "default",
			"--user", username,
			"--pass", password,
		},

		Ports: []container.ContainerPort{
			{
				Port:     4222,
				Protocol: container.ProtocolTCP,
			},
		},

		VolumeMounts: []container.VolumeMount{
			{
				Path: "/var/lib/mysql",
			},
		},
	}
}

func Info(container *container.Container) map[string]string {
	username := container.Env["USERNAME"]
	password := container.Env["PASSWORD"]

	return map[string]string{
		"Username": username,
		"Password": password,
	}
}
