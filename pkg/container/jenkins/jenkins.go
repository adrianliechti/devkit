package jenkins

import (
	"github.com/adrianliechti/devkit/pkg/container"
	"github.com/adrianliechti/devkit/pkg/to"
	"github.com/sethvargo/go-password/password"
)

const (
	DefaultShell = "/bin/bash"
)

func New() container.Container {
	image := "adrianliechti/loop-jenkins:dind"

	username := "admin"
	password := password.MustGenerate(10, 4, 0, false, false)

	return container.Container{
		Image: image,

		PlatformContext: &container.PlatformContext{
			Platform: "linux/amd64",
		},

		SecurityContext: &container.SecurityContext{
			Privileged: to.BoolPtr(true),
		},

		Env: map[string]string{
			"BASE_URL": "http://localhost:8080",

			"ADMIN_USERNAME": username,
			"ADMIN_PASSWORD": password,
		},

		Ports: []*container.ContainerPort{
			{
				Port:     8080,
				Protocol: container.ProtocolTCP,
			},
		},

		VolumeMounts: []*container.VolumeMount{
			{
				Path: "/var/jenkins_home",
			},
		},
	}
}

func Info(container *container.Container) map[string]string {
	username := container.Env["ADMIN_USERNAME"]
	password := container.Env["ADMIN_PASSWORD"]

	return map[string]string{
		"Username": username,
		"Password": password,
	}
}
