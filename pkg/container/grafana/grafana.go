package grafana

import (
	"github.com/adrianliechti/devkit/pkg/container"
	"github.com/sethvargo/go-password/password"
)

const (
	DefaultShell = "/bin/bash"
)

func New() container.Container {
	image := "grafana/grafana"

	username := "admin"
	password := password.MustGenerate(10, 4, 0, false, false)

	return container.Container{
		Image: image,

		Env: map[string]string{
			"GF_SECURITY_ADMIN_USER":     username,
			"GF_SECURITY_ADMIN_PASSWORD": password,
		},

		Ports: []container.ContainerPort{
			{
				Port:     3000,
				Protocol: container.ProtocolTCP,
			},
		},

		VolumeMounts: []container.VolumeMount{
			{
				Path: "/var/lib/grafana",
			},
		},
	}
}

func Info(container *container.Container) map[string]string {
	username := container.Env["GF_SECURITY_ADMIN_USER"]
	password := container.Env["GF_SECURITY_ADMIN_PASSWORD"]

	return map[string]string{
		"Username": username,
		"Password": password,
	}
}
