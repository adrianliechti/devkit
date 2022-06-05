package mariadb

import (
	"github.com/adrianliechti/devkit/pkg/container"
	"github.com/sethvargo/go-password/password"
)

const (
	DefaultShell = "/bin/bash"
)

func New() container.Container {
	image := "mariadb:10-focal"

	database := "db"
	password := password.MustGenerate(10, 4, 0, false, false)

	return container.Container{
		Image: image,

		Env: map[string]string{
			"MARIADB_DATABASE":      database,
			"MARIADB_ROOT_PASSWORD": password,
		},

		Ports: []container.ContainerPort{
			{
				Port:     3306,
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
	database := container.Env["MARIADB_DATABASE"]
	username := "root"
	password := container.Env["MARIADB_ROOT_PASSWORD"]

	return map[string]string{
		"Database": database,
		"Username": username,
		"Password": password,
	}
}

func ClientCmd() (string, []string) {
	return DefaultShell, []string{
		"-c",
		"mysql --user=root --password=${MARIADB_ROOT_PASSWORD} ${MARIADB_DATABASE}",
	}
}
