package vault

import (
	"github.com/adrianliechti/devkit/pkg/container"
	"github.com/sethvargo/go-password/password"
)

const (
	DefaultShell = "/bin/ash"
)

func New() container.Container {
	image := "vault:latest"

	token := password.MustGenerate(10, 4, 0, false, false)

	return container.Container{
		Image: image,

		Env: map[string]string{
			"VAULT_DEV_ROOT_TOKEN_ID": token,
		},

		Ports: []*container.ContainerPort{
			{
				Port:     8200,
				Protocol: container.ProtocolTCP,
			},
		},
	}
}

func Info(container *container.Container) map[string]string {
	token := container.Env["VAULT_DEV_ROOT_TOKEN_ID"]

	return map[string]string{
		"Token": token,
	}
}
