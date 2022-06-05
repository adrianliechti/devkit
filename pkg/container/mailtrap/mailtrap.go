package mailtrap

import (
	"github.com/adrianliechti/devkit/pkg/container"
	"github.com/sethvargo/go-password/password"
)

const (
	DefaultShell = "/bin/bash"
)

func New() container.Container {
	image := "adrianliechti/loop-mailtrap"

	username := "admin"
	password := password.MustGenerate(10, 4, 0, false, false)

	return container.Container{
		Image: image,

		PlatformContext: &container.PlatformContext{
			Platform: "linux/amd64",
		},

		Env: map[string]string{
			"USERNAME": username,
			"PASSWORD": password,
		},

		Ports: []container.ContainerPort{
			{
				Port:     25,
				Protocol: container.ProtocolTCP,
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

func ConsolePort() container.ContainerPort {
	return container.ContainerPort{
		Port:     80,
		Protocol: container.ProtocolTCP,

		HostPort: 2580,
	}
}
